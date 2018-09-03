package main

import "fmt"

type Employee interface {
	GetPay() int
	SetPay(p int)
}

type Developer struct {
	salary   int
	language string
}

func (dev *Developer) SetPay(p int) {
	dev.salary = p
}

func (dev *Developer) GetPay() int {
	return dev.salary
}

func NewDeveloper() *Developer {
	return &Developer{language: "Golang"}
}

type Manager struct {
	salary      int
	teamMembers []Employee
}

func (mana *Manager) SetPay(p int) {
	mana.salary = p
}

func (mana *Manager) GetPay() int {
	return mana.salary
}

func NewManager() *Manager {
	return &Manager{}
}

func main() {
	var employees []Employee
	var dev = NewDeveloper()
	var mana = new(Manager)

	mana.teamMembers = append(mana.teamMembers, dev)

	employees = append(employees, dev)
	employees = append(employees, mana)

	employees[0].SetPay(1500)
	employees[1].SetPay(3000)

	for _, employee := range employees {
		switch specialized := employee.(type) {
		case *Developer:
			fmt.Printf("This employee is a Developer !  %s\n", specialized)
		case *Manager:
			fmt.Printf("This employee is a Manager !  %s\n", employee)
			dev := specialized.teamMembers[0]
			fmt.Printf("His teamMembers are  %s", dev)

		}
	}

}
