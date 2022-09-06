package student

type Student struct {
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	EMAIL string `json:"email"`
}

type MESSAGE struct {
	MESSAGE string `json:"message"`
}
