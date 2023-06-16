package main

// ReviewEvent denotes a review event.
type ReviewEvent struct {
	Event            string      `json:"event"`
	Author           string      `json:"author"`
	Reviewers        []*Reviewer `json:"reviewers"`
	Reviewer         *Reviewer   `json:"reviewer"`
	ID               int         `json:"id"`
	ProjectID        string      `json:"project_id"`
	AuthorID         int         `json:"author_id"`
	ReviewableID     int         `json:"reviewable_id"`
	ReviewableType   string      `json:"reviewable_type"`
	CommitID         string      `json:"commit_id"`
	State            string      `json:"state,omitempty"`
	RestrictType     string      `json:"restrict_type"`
	PushResetEnabled bool        `json:"push_reset_enabled"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
}

// Reviewer denotes a reviewer in ReviewEvent.
type Reviewer struct {
	Reviewer  string `json:"reviewer"`
	ID        int    `json:"id"`
	ReviewID  int    `json:"review_id"`
	UserID    int    `json:"user_id"`
	ProjectID int    `json:"project_id"`
	Type      string `json:"type"`
	State     string `json:"state,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
