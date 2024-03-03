package request


type PVCount struct{
	ID         uint32 `gorm:"primary_key" json:"id"` 
	Visitor    string `json:"visitor"` // 访问者ip地址或者用户名
	VisiteDate string  `json:"visite_date"` // 访问时间
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"` // 访问内容
	Action        string `json:"action"` // 点赞，评论，等等
	Count         uint32 `json:"count"`
}