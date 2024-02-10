package request

type Login struct {
	Username string `binding:"required" validate:"required,min=1,max=255" json:"username"`
	Password string `binding:"required" validate:"required,min=1,max=15" json:"password"`
} //@name Login

type Register struct {
	Name     string `binding:"required" validate:"required,min=1,max=255" json:"name"`
	Username string `binding:"required" validate:"required,min=1,max=255" json:"username"`
	Password string `binding:"required" validate:"required,min=1,max=15" json:"password"`
} //@name Register
