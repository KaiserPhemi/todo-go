package todo

// strucutre of a todo
type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Note      string `json:"note"`
	Completed bool   `json:"completed"`
}

// template fro creating todo
type CreateTodoRequest struct {
	Title string `json:"title"`
}

type UpdateTodoRequest struct {
	Title     *string `json:"title,omitempty"`
	Completed *string `json:"completed,omitempty"`
	Note      *string `json:"note,omitempty"`
}
