package main

import (
	"encoding/json"
	"fmt"
)

// MergeRequestEvent denotes a merge request event.
type MergeRequestEvent struct {
	User         string        `json:"user"`
	MergeRequest *MergeRequest `json:"object_attributes"`
}

// MergeRequest denotes a merge request in webhook requests.
type MergeRequest struct {
	AssigneeID      int                          `json:"assignee_id"`
	AuthorID        int                          `json:"author_id"`
	CreatedAt       string                       `json:"created_at"`
	Description     string                       `json:"description"`
	ID              int                          `json:"id"`
	IID             int                          `json:"iid"`
	LastCommit      string                       `json:"last_commit"`
	MergeStatus     MergeRequestMergeStatus      `json:"merge_status,omitempty"`
	MilestoneID     int                          `json:"milestone_id"`
	Position        int                          `json:"position"`
	Source          *MergeRequestEventRepository `json:"source"`
	SourceBranch    string                       `json:"source_branch"`
	SourceProjectID string                       `json:"source_project_id"`
	State           string                       `json:"state,omitempty"`
	Target          *MergeRequestEventRepository `json:"target"`
	TargetBranch    string                       `json:"target_branch"`
	TargetProjectID string                       `json:"target_project_id"`
	Title           string                       `json:"title"`
	UpdatedAt       string                       `json:"updated_at"`
	URL             string                       `json:"url"`                        // "http://git.code.oa.com/xshi/gerrit/merge_requests/1"
	MergeType       MergeRequestMergeType        `json:"merge_type,omitempty"`       // This field is set when State is tgit.MergeRequestStateMerged and Action is MergeRequestActionMerge
	MergeCommitSHA  string                       `json:"merge_commit_sha,omitempty"` // This field is set when State is tgit.MergeRequestStateMerged and Action is MergeRequestActionMerge
	Action          MergeRequestAction           `json:"extension_action,omitempty"`
}

type MergeRequestMergeStatus int

const (
	MergeRequestMergeStatusUnspecified MergeRequestMergeStatus = iota
	MergeRequestMergeStatusUnchecked
	MergeRequestMergeStatusCanBeMerged
	MergeRequestMergeStatusCanNotBeMerged
	MergeRequestMergeStatusHookIntercept
	MergeRequestMergeStatusMissBranch
)

// MarshalJSON implements json.Marshaler.
func (s MergeRequestMergeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// String implements fmt.Stringer.
func (s MergeRequestMergeStatus) String() string {
	switch s {
	case MergeRequestMergeStatusUnspecified:
		return "<unspecified>"
	case MergeRequestMergeStatusUnchecked:
		return "unchecked"
	case MergeRequestMergeStatusCanBeMerged:
		return "can_be_merged"
	case MergeRequestMergeStatusCanNotBeMerged:
		return "cannot_be_merged"
	case MergeRequestMergeStatusHookIntercept:
		return "hook_intercept"
	case MergeRequestMergeStatusMissBranch:
		return "miss_branch"
	default:
		return fmt.Sprintf("MergeRequestMergeStatus(%d)", s)
	}
}

// ParseMergeStatusFromString parse merge status from a given string.
func ParseMergeStatusFromString(str string) MergeRequestMergeStatus {
	switch str {
	case "unchecked":
		return MergeRequestMergeStatusUnchecked
	case "can_be_merged":
		return MergeRequestMergeStatusCanBeMerged
	case "cannot_be_merged":
		return MergeRequestMergeStatusCanNotBeMerged
	case "hook_intercept":
		return MergeRequestMergeStatusHookIntercept
	case "miss_branch":
		return MergeRequestMergeStatusMissBranch
	default:
		return MergeRequestMergeStatusUnspecified
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *MergeRequestMergeStatus) UnmarshalJSON(b []byte) error {
	var t string
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	switch t {
	case "unchecked":
		*s = MergeRequestMergeStatusUnchecked
	case "can_be_merged":
		*s = MergeRequestMergeStatusCanBeMerged
	case "cannot_be_merged":
		*s = MergeRequestMergeStatusCanNotBeMerged
	case "hook_intercept":
		*s = MergeRequestMergeStatusHookIntercept
	case "miss_branch":
		*s = MergeRequestMergeStatusMissBranch
	default:
		return fmt.Errorf("unknown status %q", t)
	}
	return nil
}

// MergeRequestMergeType is the merged type of merge request
type MergeRequestMergeType int

const (
	MergeRequestMergeTypeUnspecified = iota
	MergeRequestMergeTypeMerge       // 源分支的提交通过一个合并提交点添加到目标分支
	MergeRequestMergeTypeSquashMerge // 源分支所有提交压缩成一个提交添加到目标分支
	MergeRequestMergeTypeRebase      // 源分支所有提交变基添加到目标分支，此时 merge_commit_sha 和 last_commit.id 一致
)

// MarshalJSON implements json.Marshaler.
func (s MergeRequestMergeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// String implements fmt.Stringer.
func (s MergeRequestMergeType) String() string {
	switch s {
	case MergeRequestMergeTypeUnspecified:
		return "<unspecified>"
	case MergeRequestMergeTypeMerge:
		return "Merge"
	case MergeRequestMergeTypeSquashMerge:
		return "SquashAndMerge"
	case MergeRequestMergeTypeRebase:
		return "RebaseAndMerge"
	default:
		return fmt.Sprintf("MergeRequestMergeType(%d)", s)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *MergeRequestMergeType) UnmarshalJSON(b []byte) error {
	var t string
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	switch t {
	case "Merge":
		*s = MergeRequestMergeTypeMerge
	case "SquashAndMerge":
		*s = MergeRequestMergeTypeSquashMerge
	case "RebaseAndMerge":
		*s = MergeRequestMergeTypeRebase
	case "":
		// 当 MR 未合入时该字段工蜂返回 null，unmarshal 需要处理 string 的默认值。
		*s = MergeRequestMergeTypeUnspecified
	default:
		return fmt.Errorf("unknown merge type %q", t)
	}
	return nil
}

// MergeRequestAction denotes a merge request action.
type MergeRequestAction int

const (
	MergeRequestActionUnspecified MergeRequestAction = iota
	MergeRequestActionOpen                           // 新建 MR
	MergeRequestActionClose                          // MR 被关闭
	MergeRequestActionReopen                         // MR 重新被打开
	MergeRequestActionUpdate                         // MR 本身信息更新
	MergeRequestActionPushUpdate                     // 源分支有代码 push
	MergeRequestActionMerge                          // 代码已合并
)

// MarshalJSON implements json.Marshaler.
func (a MergeRequestAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// String implements fmt.Stringer.
func (a MergeRequestAction) String() string {
	switch a {
	case MergeRequestActionUnspecified:
		return "<unspecified>"
	case MergeRequestActionOpen:
		return "open"
	case MergeRequestActionClose:
		return "close"
	case MergeRequestActionReopen:
		return "reopen"
	case MergeRequestActionUpdate:
		return "update"
	case MergeRequestActionPushUpdate:
		return "push-update"
	case MergeRequestActionMerge:
		return "merge"
	default:
		return fmt.Sprintf("MergeRequestAction(%d)", a)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *MergeRequestAction) UnmarshalJSON(b []byte) error {
	var t string
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	switch t {
	case "open":
		*a = MergeRequestActionOpen
	case "close":
		*a = MergeRequestActionClose
	case "reopen":
		*a = MergeRequestActionReopen
	case "update":
		*a = MergeRequestActionUpdate
	case "push-update":
		*a = MergeRequestActionPushUpdate
	case "merge":
		*a = MergeRequestActionMerge
	default:
		return fmt.Errorf("unknown action %q", t)
	}
	return nil
}

// MergeRequestEventRepository denotes a repository in MergeRequestEvent.
type MergeRequestEventRepository struct {
	Name            string `json:"name"`             // "gerrit"
	SSHURL          string `json:"ssh_url"`          // "git@git.code.oa.com:xshi/gerrit.git"
	HTTPURL         string `json:"http_url"`         // "http://git.code.oa.com/xshi/gerrit.git"
	WebURL          string `json:"web_url"`          // "http://git.code.oa.com/xshi/gerrit"
	Namespace       string `json:"namespace"`        // "xshi/gerrit"
	VisibilityLevel string `json:"visibility_level"` // 10
}
