// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type F interface {
	Print()
}
type S struct {
	S1
	data string
}

func (s S) Read() string {
	return s.data
}

// func (s S) Print() {
// 	fmt.Println("Thats S")
// }

func (s *S) Write(str string) {
	s.data = str
}

type S1 struct{}

func (s S1) Print() {
	fmt.Println("Thats S1")
}

type S2 struct{}

func (s *S2) Print() {
	fmt.Println("Thats S2")
}

func main2() {
	sVal := S{S1{}, "A"}
	// sPtr := &S{}
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	sVal.Print()
	// sPtr.Print()

	s1Val.Print()
	s1Ptr.Print()

	s2Val.Print()
	s2Ptr.Print()

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr

	fmt.Println(i)
	i.Print()
}

func main3() {
	sVals := map[int]S{1: {data: "A"}}

	// You can only call Read using a value
	sVals[1].Read()

	// This will not compile:
	//  sVals[1].Write("test")

	sPtrs := map[int]*S{1: {data: "A"}}

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

func main8() {
	a := getAdder()
	a.CallMe()
}

type MainSchema struct {
	UserID      int64
	FirstName   string
	LastName    string
	Email       string
	Mobile      string
	DisplayName string
	UserData    string `json:"-"`
}

type SubSchema struct {
	FirstName string
	LastName  string
}

func copy(unknown interface{}) interface{} {
	fmt.Println(unknown)
	small := &SubSchema{}
	copier.Copy(small, unknown)
	fmt.Println(small)
	small.FirstName = "Suvarna"
	small.LastName = "Puranik"
	copier.Copy(unknown, small)
	fmt.Println(unknown)
	return unknown
}

func main9() {
	u := MainSchema{
		UserID:    1,
		FirstName: "Sachin",
		LastName:  "Kumar",
		Email:     "abc@xyz.com",
	}
	res := copy(u)
	fmt.Println(res)
}

func main() {
	coursesHandler := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from the other side",
		})
	}
	router := gin.Default()

	// Handler for "/website/courses"
	router.GET("/website/courses", coursesHandler)

	// Serve static files for other paths
	router.Static("/", "./website")

	// Run the server
	router.Run(":8080")
}
