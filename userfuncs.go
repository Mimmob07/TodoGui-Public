package main

import (
	"crypto/sha256"
	"fmt"
)

func login(liveusername string, livepassword string) []User {
	var users []User
	PwordSum := sha256.Sum256([]byte(livepassword))

	rows, err := db.Query("SELECT * FROM user WHERE username = ? AND password = ?", liveusername, fmt.Sprintf("%x", PwordSum))
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.Name, &usr.Username, &usr.Email, &usr.Password, &usr.Created_at); err != nil {
			return nil
		}
		users = append(users, usr)
		userId = int64(usr.Id)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return users
}

func createUser(usr User) (int64, error) {
	result, err := db.Exec("INSERT INTO user (name, username, email, password) VALUES (?, ?, ?, ?)", usr.Name, usr.Username, usr.Email, usr.Password)
	if err != nil {
		return 0, fmt.Errorf("createUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("createUser: %v", err)
	}
	userId = id
	return id, nil
}
