package forums

type Forum struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Topic   string `json:"topic"`
	UsersId []int  `json:"usersId"`
}

func NewForum(id int, name, topic string, usersId []int) Forum {
	return Forum{Id: id, Name: name, Topic: topic, UsersId: usersId}
}
