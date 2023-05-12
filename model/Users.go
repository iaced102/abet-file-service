package model

type Users struct {
	Id            string   `json:"id" param:"id" query:"id" form:"id" pg:"id,pk"`
	UserType      int      `json:"userType" db:"userType" form:"userType" pg:"user_type,use_zero"`
	UserName      string   `json:"userName" db:"userName" form:"userName" pg:"user_name"`
	CreatedAt     string   `json:"createdAt" db:"createdAt" form:"createdAt" pg:"created_at"`
	UpdatedAt     string   `json:"updatedAt" db:"updatedAt" form:"updatedAt" pg:"updated_at"`
	Password      string   `json:"password" db:"password" form:"password" pg:"password"`
	FirstName     string   `json:"firstName" db:"first_name" form:"firstName" pg:"first_name"`
	LastName      string   `json:"lastName" db:"last_name" form:"lastName" pg:"last_name"`
	Email         string   `json:"email" db:"email" form:"email" pg:"email"`
	CryptPassword string   `json:"cryptPassword" db:"cryptPassword" form:"cryptPassword" pg:"crypt_password"`
	CustomField   string   `json:"customField" form:"customField" pg:"custom_field"`
	tableName     struct{} `pg:"users"`
}
