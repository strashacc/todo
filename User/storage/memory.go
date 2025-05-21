package storage

import (
	"sync"
	"user/models"
)

var (
	Users = make(map[string]models.User)
	mu    sync.Mutex
)

func AddUser(user models.User) {
	mu.Lock()
	defer mu.Unlock()
	Users[user.ID] = user
}

func GetUser(id string) (models.User, bool) {
	mu.Lock()
	defer mu.Unlock()
	user, exists := Users[id]
	return user, exists
}

func DeleteUser(id string) {
	mu.Lock()
	defer mu.Unlock()
	delete(Users, id)
}

func GetAllUsers() []models.User {
	mu.Lock()
	defer mu.Unlock()
	users := make([]models.User, 0, len(Users))
	for _, user := range Users {
		users = append(users, user)
	}
	return users
}
