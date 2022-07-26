package model

import "log"

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
}

func (fm *FlowModel) ListUser() []User {
	var ul []User
	fm.db.Find(&ul)
	log.Printf("user list is: %s\n", ul)
	return ul
}
