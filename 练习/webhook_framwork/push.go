package main

import (
	"encoding/json"
	"fmt"
)

// PushEvent denotes a push event.
type PushEvent struct {
	User        string               `json:"user"`
	Before      string               `json:"before"`       // "427f1307c6460b6908dfb814246c8a3cb7e6be44"
	After       string               `json:"after"`        // "54acf9c910a7d6516a6f711a28be5f0e0444138b"
	Ref         string               `json:"ref"`          // "refs/heads/cl/init"
	CheckoutSHA string               `json:"checkout_sha"` // "54acf9c910a7d6516a6f711a28be5f0e0444138b"
	UserName    string               `json:"user_name"`    // "xshi"
	UserID      int                  `json:"user_id"`      // 142910
	UserEmail   string               `json:"user_email"`   // "xshi@tencent.com"
	ProjectID   string               `json:"project_id"`   // 412212
	Repository  *PushEventRepository `json:"repository"`
	Commits     []*Commit            `json:"commits"`
	Operation   PushOperation        `json:"operation_kind,omitempty"`
	Action      PushAction           `json:"action_kind,omitempty"`
}

// PushOperation is the operation of a PushEvent.
type PushOperation int

const (
	PushOperationUnspecified PushOperation = iota
	PushOperationCreate                    // 创建分支、合并普通 MR 的 push
	PushOperationDelete                    // 删除分支的 push
	PushOperationUpdate                    // 文件修改的 push
	PushOperationUpdateForce               // non-fast-forward push
)

// MarshalJSON implements json.Marshaler.
func (op PushOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.String())
}

// String implements fmt.Stringer.
func (op PushOperation) String() string {
	switch op {
	case PushOperationUnspecified:
		return "<unspecified>"
	case PushOperationCreate:
		return "create"
	case PushOperationDelete:
		return "delete"
	case PushOperationUpdate:
		return "update"
	case PushOperationUpdateForce:
		return "update_nonfastword"
	default:
		return fmt.Sprintf("PushOperation(%d)", op)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (op *PushOperation) UnmarshalJSON(b []byte) error {
	var t string
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	switch t {
	case "create":
		*op = PushOperationCreate
	case "delete":
		*op = PushOperationDelete
	case "update":
		*op = PushOperationUpdate
	case "update_nonfastword":
		*op = PushOperationUpdateForce
	default:
		return fmt.Errorf("unknown operation %q", t)
	}
	return nil
}

// PushAction is the action of a PushEvent.
type PushAction int

const (
	PushActionUnspecified        PushAction = iota
	PushActionClientPush                    // 客户端请求（不在工蜂 Web 上操作的都默认是这个）
	PushActionCreateBranch                  // 在工蜂 Web 上创建分支
	PushActionDeleteBranch                  // 在工蜂 Web 上删除分支
	PushActionCreateTag                     // 在工蜂 Web 上创建 Tag
	PushActionDeleteTag                     // 在工蜂 Web 上删除 Tag
	PushActionCreateFile                    // 在工蜂 Web 上新建文件
	PushActionModifyFile                    // 在工蜂 Web 上修改文件
	PushActionDeleteFile                    // 在工蜂 Web 上删除文件
	PushActionReplaceFile                   // 在工蜂 Web 上替换文件
	PushActionCreateAMergeCommit            // 在工蜂 Web 上合并普通 MR
	PushActionSquashAndMerge                // 在工蜂 Web 上合并 squash merge
	PushActionRebaseAndMerge                // 在工蜂 Web 上合并 rebase merge
	PushActionCherryPick                    // 在工蜂 Web 上使用 cherry pick 功能
	PushActionRevert                        // 在工蜂 Web 上使用 revert 功能
)

// MarshalJSON implements json.Marshaler.
func (a PushAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// String implements fmt.Stringer.
func (a PushAction) String() string {
	switch a {
	case PushActionUnspecified:
		return "<unspecified>"
	case PushActionClientPush:
		return "client push"
	case PushActionCreateBranch:
		return "create branch"
	case PushActionDeleteBranch:
		return "delete branch"
	case PushActionCreateTag:
		return "create tag"
	case PushActionDeleteTag:
		return "delete tag"
	case PushActionCreateFile:
		return "create file"
	case PushActionModifyFile:
		return "modify file"
	case PushActionDeleteFile:
		return "delete file"
	case PushActionReplaceFile:
		return "replace file"
	case PushActionCreateAMergeCommit:
		return "create a merge commit"
	case PushActionSquashAndMerge:
		return "squash and merge"
	case PushActionRebaseAndMerge:
		return "rebase and merge"
	case PushActionCherryPick:
		return "cherry-pick"
	case PushActionRevert:
		return "revert"
	default:
		return fmt.Sprintf("PushAction(%d)", a)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *PushAction) UnmarshalJSON(b []byte) error {
	var t string
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	switch t {
	case "client push":
		*a = PushActionClientPush
	case "create branch":
		*a = PushActionCreateBranch
	case "delete branch":
		*a = PushActionDeleteBranch
	case "create tag":
		*a = PushActionCreateTag
	case "delete tag":
		*a = PushActionDeleteTag
	case "create file":
		*a = PushActionCreateFile
	case "modify file":
		*a = PushActionModifyFile
	case "delete file":
		*a = PushActionDeleteFile
	case "replace file":
		*a = PushActionReplaceFile
	case "create a merge commit":
		*a = PushActionCreateAMergeCommit
	case "squash and merge":
		*a = PushActionSquashAndMerge
	case "rebase and merge":
		*a = PushActionRebaseAndMerge
	case "cherry-pick":
		*a = PushActionCherryPick
	case "revert":
		*a = PushActionRevert
	default:
		return fmt.Errorf("unknown action %q", t)
	}
	return nil
}

// PushEventRepository denotes a repository in PushEvent.
//
// NB. It looks just like `MergeRequestEventRepository`, but apparently, tgit
// team is really good at naming the same thing differently (yes, i'm being
// sarcasm).
type PushEventRepository struct {
	Name            string `json:"name"`             // "gerrit"
	GitSSHURL       string `json:"git_ssh_url"`      // "git@git.code.oa.com:xshi/gerrit.git"
	GitHTTPURL      string `json:"git_http_url"`     // "http://git.code.oa.com/xshi/gerrit.git"
	URL             string `json:"url"`              // "http://git.code.oa.com/xshi/gerrit"
	VisibilityLevel string `json:"visibility_level"` // 10
}

// Commit denotes a commit in webhook requests.
type Commit struct {
	SHA       string        `json:"id"`
	Message   string        `json:"message"`
	Timestamp string        `json:"timestamp"`
	URL       string        `json:"url"` // "http://git.code.oa.com/xshi/gerrit/commit/54acf9c910a7d6516a6f711a28be5f0e0444138b"
	Author    *CommitAuthor `json:"author"`
	Added     []string      `json:"added"`
	Modified  []string      `json:"modified"`
	Removed   []string      `json:"removed"`
}

// CommitAuthor denotes the a author in Commit.
type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
