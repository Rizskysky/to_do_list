package requests

// AddTodo is a structure for any st
type AddTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

// EditTodo is a structure
type EditTodo struct {
	Id uint `json:"id"`
}

//DeleteTodo is a structure
type DeleteTodo struct {
	Id uint `json:"id"`
}
