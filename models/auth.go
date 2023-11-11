package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique "json:"email"`
}

// Add User
func AddUser(u *User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func TryLogin(username, email, password string) bool {
	// Get users by username or email and retrieve the stored password hash
	var u User
	if username != "" {
		if err := db.Where("username = ?", username).First(&u).Error; err != nil {
			return false
		}
	}
	// Check email
	if email != "" {
		if err := db.Where("email = ?", email).First(&u).Error; err != nil {
			return false
		}
	}

	// Check password
	if !checkPasswordHash(password, u.Password) {
		return false
	}
	// TODO: Increment login count
	return true
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserByEmail(email string) (*User, error) {
	var user *User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
