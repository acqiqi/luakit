package daoru

type Comment struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	Flag        string `json:"flag"`
	Utype       string `json:"utype"`
	ProjectType string `json:"project_type"`
	UID         string `json:"uid"`
	ProjectID   string `json:"project_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Pics        string `json:"pics"`
	Star        string `json:"star"`
	Tags        string `json:"tags"`
	QuckID      string `json:"quck_id"`
}
