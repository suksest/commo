package model

import (
	"crypto/md5"
	"fmt"
	"io"
)

// User model
type User struct {
	ID       uint
	Name     string
	Phone    string
	Role     string
	Password string `size:"4"`
}

// TableName will return table name
func (u *User) TableName() string {
	return "user"
}

//hashPassword will generate hashed password
func hashPassword(password string) (hashed string) {
	h := md5.New()
	io.WriteString(h, password)

	hashed = fmt.Sprintf("%x", h.Sum(nil))

	return
}

// FindUserByID will lookup user by ID
func FindUserByID(ID uint) (u *User, err error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Table(u.TableName()).Select("phone, name, role").Where("id = ?", ID).First(u).Error
	if err != nil {
		return nil, err
	}

	return
}

// FindUserByPhone will lookup user by phone number
func FindUserByPhone(Phone string) (u *User, err error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	u = &User{}
	err = db.Table(u.TableName()).Select("phone, name, role, password").Where("phone = ?", Phone).First(u).Error
	if err != nil {
		return nil, err
	}

	return
}

// CreateUser will create single user
func CreateUser(user User) (result *User, err error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user.Password = hashPassword(user.Password)

	err = db.Table(user.TableName()).Create(&user).Error
	if err != nil {
		return result, err
	}

	result, err = FindUserByPhone(user.Phone)
	if err != nil {
		return nil, err
	}

	return
}
