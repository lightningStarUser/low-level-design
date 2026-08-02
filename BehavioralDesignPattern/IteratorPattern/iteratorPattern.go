package main

import "fmt"

type Users struct {
	name string
	age  int
}

type Collection interface {
	createIterator() Iterator
}

type UserCollection struct {
	users []*Users
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

type Iterator interface {
	hasNext() bool
	getNext() *Users
}

type UserIterator struct {
	index int
	users []*Users
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() *Users {
	user := u.users[u.index]
	u.index++
	return user
}

func main() {
	userA := &Users{
		name: "Tejaswini",
		age:  25,
	}
	userB := &Users{
		name: "Soundarya",
		age:  30,
	}
	userC := &Users{
		name: "Rathna",
		age:  56,
	}
	uc := &UserCollection{
		users: []*Users{userA, userB, userC},
	}
	itr := uc.createIterator()
	for itr.hasNext() {
		user := itr.getNext()
		fmt.Printf("User: %v\n", *user)
	}
}