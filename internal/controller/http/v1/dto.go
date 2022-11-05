package v1

type messageDTO struct {
	Content string  `json:"content"`
	Sender  userDTO `json:"sender"` //AuthorUUID
	Date    string  `json:"date"`   //createdAt
}

type chatItemDTO struct {
	ChatName    string     `json:"chatName"`
	ChatUUID    string     `json:"chatUUID"`
	LastMessage messageDTO `json:"lastMessage"`
	ProjectUUID string     `json:"projectUUID"`
	ImageURL    string     `json:"imageURL"`
}

type chatHistoryDTO struct {
	Messages []messageDTO `json:"messages"`
}

type userDTO struct {
	UUID       string `json:"UUID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Patronymic string `json:"patronymic"`
}

//type chatDTO struct {
//	UUID        string         `json:"UUID"`
//	ProjectUUID string         `json:"projectUUID"`
//	History     chatHistoryDTO `json:"history"`
//	ChatType    string         `json:"type"`
//	Users       []userDTO      `json:"users"`
//}
