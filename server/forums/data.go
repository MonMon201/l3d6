package forums

import (
	"database/sql"
	"fmt"
)

type Forum struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Topic   string `json:"topic"`
	UsersId []int  `json:"usersId"`
}

func NewForum(id int, name, topic string, usersId []int) Forum {
	return Forum{Id: id, Name: name, Topic: topic, UsersId: usersId}
}

type ForumStore struct {
	Db *sql.DB
}

func NewForumStore(db *sql.DB) *ForumStore {
	return &ForumStore{Db: db}
}

func (s *ForumStore) ListChannels() ([]*Forum, error) {
	rows, err := s.Db.Query("SELECT id, name FROM forums LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var f Forum
		if err := rows.Scan(&f.Id, &f.Name); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil
}

func (s *ForumStore) CreateForum(name, topicKeyword string) error {
	if len(name) < 0 {
		return fmt.Errorf("channel name is not provided")
	}
	_, err := s.Db.Exec(`INSERT INTO forums (name, "topicKeyword") VALUES ($1, $2)`, name, topicKeyword)
	return err
}
