package response

type CreateUser struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
