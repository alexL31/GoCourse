package main

import (
	"fmt"
	"sort"
)

type contact struct {
	Name  string
	Phone string
}

func (с contact) String() string {
	return fmt.Sprintf("%s: %s", с.Name, с.Phone)
}

type contactBook []contact

func (a contactBook) Len() int           { return len(a) }
func (a contactBook) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a contactBook) Less(i, j int) bool { return a[i].Phone < a[j].Phone }

func main() {
	people := []contact{
		{"Bob", "87789890"},
		{"John", "123456"},
		{"Michael", "567899000"},
		{"Jenny", "988765433"},
	}

	fmt.Println(people)

	sort.Sort(contactBook(people))
	fmt.Println(people)

}
