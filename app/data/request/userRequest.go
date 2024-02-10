package request

type CreateUser struct {
	Name     string `binding:"required" validate:"required,min=1,max=255" json:"name"`
	Email    string `binding:"required" validate:"required,min=1,max=255" json:"email"`
	Password string `binding:"required" validate:"required,min=1,max=15" json:"password"`
	Active   bool   `json:"active"`
} //@name CreateUser

type CreateProduct struct {
	Description string  `binding:"required" validate:"required,min=1,max=45" json:"description"`
	Price       float32 `binding:"required" validate:"required" json:"price"`
	Active      bool    `json:"active"`
	Type        int     `json:"type"`
} //@name CreateProduct
