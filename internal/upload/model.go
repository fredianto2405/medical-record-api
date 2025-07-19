package upload

import "time"

type Entity struct {
	ID         string     `db:"id"`
	Filename   string     `db:"filename"`
	Filepath   string     `db:"filepath"`
	Size       int64      `db:"size"`
	MimeType   string     `db:"mime_type"`
	UploadedAt *time.Time `db:"uploaded_at"`
}
