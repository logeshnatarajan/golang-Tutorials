package main

import "fmt"

type salary float64

func (s salary) total() total {
	return total(s)
}

type total float64

func (t total) hra() hra {
	t += t * 0.3 // 30% HRA Addition
	return hra(t)
}
func (t total) salary() salary {
	t -= t * 0.10 // 10% Tax Deduction
	return salary(t)
}
func (s salary) basic() basic {
	return basic(s)
}

type hra float64

func (h hra) basic() basic {
	h += h * 0.3 // 30% HRA Addition
	return basic(h)
}
func (h hra) total() total {
	return total(h)
}

type basic float64

func (b basic) total() total {
	return total(b)
}

func main() {
	fmt.Println("Salary calculation for First Employee:")
	sal1 := basic(9000.00)                           //ipo ithu basic type
	fmt.Println(sal1.total())                        //basic total method kuda attach airuku ipo ithu vanthu type total
	fmt.Println(sal1.total().hra().total())          //ipo type total vanthu hrs la attach airuku anga 30% athigam agi type hra aguthu apro total aguthu
	fmt.Println(sal1.total().hra().total().salary()) //type total vanthu salary method kuda attach airuku anga 10% ++ aguthu

	fmt.Println("\nSalary calculation for Second Employee:")
	sal2 := basic(5000.00)
	fmt.Println(sal2.total())
	fmt.Println(sal2.total().salary())
}
