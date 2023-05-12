package model

type File struct {
	Name string `db:"name" json:"fileName" form:"fileName"`
}

func (File) GetTableName() string { return "file" }
