package main

// NoteEvent denotes a note event.
type NoteEvent struct {
	User         string               `json:"user"`
	ProjectID    string               `json:"project_id"`
	Repository   *NoteEventRepository `json:"repository"`
	Note         *Note                `json:"object_attributes"`
	Commit       string               `json:"commit,omitempty"`
	MergeRequest string               `json:"merge_request,omitempty"`
}

// NoteEventRepository denotes a repository in NoteEvent.
type NoteEventRepository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

// Note denotes a note.
type Note struct {
	ID           int    `json:"id"`
	Note         string `json:"note"`
	NoteableType string `json:"noteable_type"`
	AuthorID     int    `json:"author_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	ProjectID    string `json:"project_id"`
	// Attachment   interface{} `json:"attachment"`
	// LineCode     *string     `json:"line_code"`
	CommitID string `json:"commit_id"`
	// NoteableID *int64 `json:"noteable_id"`
	System bool `json:"system"`
	// StDiff       *StDiff     `json:"st_diff"`
	URL string `json:"url"`
}
