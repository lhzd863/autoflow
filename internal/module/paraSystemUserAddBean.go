package module

type MetaParaSystemUserAddBean struct {
	Id           string `json:"id"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	CreateTime   string `json:"createtime"`
	UpdateTime   string `json:"updatetime"`
	Enable       string `json:"enable"`
}
