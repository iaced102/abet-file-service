package outgoing

/*
create by: Hoangnd
create at: 2023-01-01
des: Khai báo và xử lý thông tin nhận từ request
*/

type AddSystemAccountParams struct {
	Id                  string `json:"Id" param:"Id" db:"Id" query:"id" form:"Id" query:"Id"  primary:"true"`
	Description         string `json:"Description" db:"Description" form:"Description" pg:"description"`
	UserRelated         string `json:"UserRelated" db:"UserRelated" form:"UserRelated" pg:"user_related"`
	SystemAccountTypeId string `json:"SystemAccountTypeId" db:"SystemAccountTypeId" form:"SystemAccountTypeId" pg:"system_account_type_id"`
	UserName            string `json:"UserName" db:"UserName" form:"UserName" pg:"user_name"`
	CreatedAt           string `json:"CreatedAt" db:"CreatedAt" form:"CreatedAt" pg:"created_at"`
	UpdatedAt           string `json:"UpdatedAt" db:"UpdatedAt" form:"UpdatedAt" pg:"updated_at"`
	Password            string `json:"Password" db:"Password" form:"Password" pg:"password"`
}
