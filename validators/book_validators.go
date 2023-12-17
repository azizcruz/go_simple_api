package validators

type CreateBookValidator struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Genre  string `json:"genre" binding:"required"`
}

type UpdateBookValidator struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Genre  string `json:"genre" binding:"required"`
}
