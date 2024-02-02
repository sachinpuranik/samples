// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

type F interface {
	Print()
}

type S1 struct{}

func (s S1) Print() {
	fmt.Println("Thats s1")
}

type S2 struct{}

func (s *S2) Print() {
	fmt.Println("Thats s2")
}

func main2() {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	s1Val.Print()
	s1Ptr.Print()
	s2Val.Print()
	s2Ptr.Print()

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr

	fmt.Println(i)
}

func main3() {
	sVals := map[int]S{1: {"A"}}

	// You can only call Read using a value
	sVals[1].Read()

	// This will not compile:
	//  sVals[1].Write("test")

	sPtrs := map[int]*S{1: {"A"}}

	// You can call both Read and Write using a pointer
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}

type myint int

func (i *myint) MyPrint() {
	fmt.Println(i)
}

func (i myint) Print() {
	fmt.Println(i)
}

func main4() {
	var i myint
	i = 88
	p := &i

	i.Print()
	p.Print()

	i.MyPrint()
	p.MyPrint()
}

func checkHoursElapsed(t time.Time, elapsedTime int) bool {
	pastTime := time.Now().Add(time.Duration(-1*elapsedTime) * time.Hour)
	fmt.Println("Calculated PastTIme :", pastTime)
	timePassed := t.Before(pastTime)

	return timePassed
}

func getDate(myDateString string) time.Time {
	myDate, err := time.Parse("2006-01-02 15:04", myDateString)
	if err != nil {
		panic(err)
	}
	return myDate.Local()
}

func main5() {
	myDateString := "2022-04-06 06:00"
	myDate := getDate(myDateString)
	fmt.Println("Input Time :", myDate)

	elapsed := checkHoursElapsed(myDate, 6)
	if elapsed {
		fmt.Println("time is already elapsed")
	} else {
		fmt.Println("time is not elapsed")
	}

	myDateString = "2022-04-06 06:30"
	myDate = getDate(myDateString)
	fmt.Println("Input Time :", myDate)

	elapsed = checkHoursElapsed(myDate, 6)
	if elapsed {
		fmt.Println("time is already elapsed")
	} else {
		fmt.Println("time is not elapsed")
	}
}

func runesToString(runes []rune) (outString string) {
	// don't need index so _
	for i, v := range runes {
		fmt.Println(i, string(v))
	}
	return
}

func main6() {
	var s string = "hellosachin"
	runesToString([]rune(s))
}

func main7() {
	ii := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110}

	// proof - range retains the snapshot of the given data structure
	for k, v := range ii {
		fmt.Printf("k : %d , v: %d , arr : %v\n", k, v, ii)
		if len(ii) >= 2 {
			ii = ii[2:]
		} else if len(ii) == 1 {
			ii = ii[1:]
		} else {
			ii = ii[0:]
		}
	}
}

type adder func(a int, b int) int

func (a adder) CallMe() {
	fmt.Println("wow its", a(1, 2))
}

func getAdder() adder {
	return func(a int, b int) int {
		return a + b
	}
}

func main() {
	a := getAdder()
	a.CallMe()
}
