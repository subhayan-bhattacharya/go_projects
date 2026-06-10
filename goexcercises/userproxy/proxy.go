package userproxy

import "fmt"

type User struct {
	Id int
}
type UserList []User

func (u *UserList) Find(id int) (User, error) {
	for _, user := range *u {
		if user.Id == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("could not find user with this id %d", id)
}

func (u *UserList) AddUser(user User) {
	*u = append(*u, user)
}

type UserFinder interface {
	Find(id int) (User, error)
}

type UserListProxy struct {
	Database  UserList
	Cache     UserList
	Capacity  int
	UsedCache bool
}

func (u *UserListProxy) AddUser(user User) {
	if len(u.Cache) == u.Capacity {
		u.Cache = u.Cache[1:]
		u.Cache.AddUser(user)
	} else {
		u.Cache.AddUser(user)
	}

}

func (u *UserListProxy) Find(id int) (User, error) {
	user, err := u.Cache.Find(id)
	if err != nil {
		user, err = u.Database.Find(id)
		if err != nil {
			return User{}, fmt.Errorf("the user with id %d does not exist", id)
		}
		u.Cache.AddUser(user)
		return user, nil
	}
	u.UsedCache = true
	return user, nil
}
