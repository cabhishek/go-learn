package main

import (
    "fmt"
)

type Person struct {
    name string
}

func (p Person) Name(){
    fmt.Println(p.name)
}

type Manager struct {
    company string
}

func (m Manager) Company(){
    fmt.Println(m.company)
}

// multiple inheritance
// anonymous Person and Manager declaration
type PersonManager struct {
    Person
    Manager
}

func main(){

    p := PersonManager{
            Person {"Slim Shady"},
            Manager{"Google"},
        }

    p.Name()
    p.Company()
}
