package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name   string
	Age    int
	Saying []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []user

func (l ByName) Len() int           { return len(l) }
func (l ByName) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByName) Less(i, j int) bool { return l[i].Name < l[j].Name }

func main() {
	u1 := user{"logesh", 20, []string{"mney ", "tha ", "ellam"}}
	u2 := user{"kan", 30, []string{"mney ", "matum  ", "pathathu"}}
	u3 := user{"har", 21, []string{"mney ", "vendam ", "ellam paper"}}
	users := []user{u1, u2, u3}

	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByName(users))
	fmt.Println(users)

}
