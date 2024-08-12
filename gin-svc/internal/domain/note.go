package domain

type DNote struct {
	ID          int `json:"id"`
	NoteTitle   string
	NoteContent string
	Cover       string
	Address     string
	PublishTime string
	Private     bool
}
