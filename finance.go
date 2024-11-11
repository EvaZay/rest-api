package finance

type FinanceList struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type FinanceItem struct {
	Id   int  `json:"id"`
	Done bool `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
