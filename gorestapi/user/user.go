package userpkg

import (
	"errors"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const (
	dbPath = "users.db"
)

var (
	ErrRecordInvalid = errors.New("user is blank")
)

//返回所有用户信息

func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	users := []User{}

	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

//查询用户
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//删除用户
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}

// 保存用户
func (u *User) Save() error {
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)
}

//用户验证

func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
