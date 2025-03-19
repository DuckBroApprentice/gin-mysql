package model

import (
	"fmt"
	"log"
	db "tryweb/database"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	UserPwd  string `json:"user_pwd"`
}

func (u *User) Add() (int64, error) {
	newUser, err := db.MySq.Exec("INSERT INTO users(username,userpwd) VALUES(?,?)", u.Username, u.UserPwd)
	if err != nil {
		return 0, err
	}
	id, err := newUser.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *User) Get() (User, error) {
	user := User{}
	log.Println(db.B)
	err := db.MySq.Ping()
	if err != nil {
		log.Println("Get PONG")
	}
	row := db.MySq.QueryRow("SELECT id,username,user_pwd FROM users WHERE id = ?", u.Id)
	err = row.Scan(&user.Id, &user.Username, &user.UserPwd)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) GetAll() ([]User, error) {
	userList := []User{}
	rows, err := db.MySq.Query("SELECT * FROM users")
	// defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		oneUser := User{}
		rows.Scan(&oneUser.Id, &oneUser.Username, &oneUser.UserPwd)
		userList = append(userList, oneUser)
	}
	return userList, nil
}

func (u *User) Del() (int64, error) {
	sql := fmt.Sprintf("DELETE FROM users WHERE id = %d", u.Id)
	result, err := db.MySq.Exec(sql)
	if err != nil {
		return 0, err
	}
	num, _ := result.RowsAffected() //回傳n行執行
	return num, nil
}

func (u *User) Update() (int64, error) { //只能修改密碼
	result, err := db.MySq.Exec("UPDATE users SET user_pwd=? WHERE id=?", u.UserPwd, u.Id)
	if err != nil {
		return 0, err
	}
	num, _ := result.RowsAffected()
	return num, nil
}
