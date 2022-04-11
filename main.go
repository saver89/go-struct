package main

import "fmt"

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

type ContactList struct {
	list []Contact
}

func (cl *ContactList) Create(elem Contact) bool {
	cl.list = append(cl.list, elem)
	return true
}

func (cl ContactList) Update(id int, elem Contact) bool {
	for i, v := range cl.list {
		if v.ID == id {
			cl.list[i] = elem
			return true
		}
	}
	return false
}

func (cl ContactList) Get(id int) (Contact, bool) {
	for i, v := range cl.list {
		if v.ID == id {
			return cl.list[i], true
		}
	}
	return Contact{}, false
}

func (cl ContactList) GetAll() []Contact {
	return cl.list
}

func (cl *ContactList) Delete(id int) bool {
	for i, v := range cl.list {
		if v.ID == id {
			cl.list[i] = cl.list[0]
			cl.list = cl.list[1:]
			return true
		}
	}
	return false
}

type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

type TaskList struct {
	list []Task
}

func (tl *TaskList) Create(elem Task) bool {
	tl.list = append(tl.list, elem)
	return true
}

func (tl TaskList) Update(id int, elem Task) bool {
	for i, v := range tl.list {
		if v.ID == id {
			tl.list[i] = elem
			return true
		}
	}
	return false
}

func (tl TaskList) Get(id int) (Task, bool) {
	for i, v := range tl.list {
		if v.ID == id {
			return tl.list[i], true
		}
	}
	return Task{}, false
}

func (tl TaskList) GetAll() []Task {
	return tl.list
}

func (tl *TaskList) Delete(id int) bool {
	for i, v := range tl.list {
		if v.ID == id {
			tl.list[i] = tl.list[0]
			tl.list = tl.list[1:]
			return true
		}
	}
	return false
}

func main() {
	cl := ContactList{}
	cl.Create(Contact{
		ID:        1,
		FirstName: "Vova",
		LastName:  "Nee",
		Phone:     "+998 (97) 445 88 57",
		Email:     "neelkim@mail.ru",
		Position:  "developer",
	})

	cl.Create(Contact{
		ID:        2,
		FirstName: "Misha",
		LastName:  "Kim",
		Phone:     "+998 (97) 123 45 68",
		Email:     "kimisha@mail.ru",
		Position:  "manager",
	})

	fmt.Printf("GetAll result: %v\n", cl.GetAll())
	getContact, _ := cl.Get(2)
	fmt.Printf("Get(2) result: %v\n", getContact)
	cl.Delete(2)
	fmt.Printf("after Delete(2) result: %v\n", cl.GetAll())

	cl.Update(1, Contact{
		ID:        1,
		FirstName: "Vladislav",
		LastName:  "Nee",
		Phone:     "+998 (97) 445 88 57",
		Email:     "neelkim@mail.ru",
		Position:  "developer",
	})

	fmt.Printf("after Update(1) result: %v\n", cl.GetAll())
}
