package main

// import "fmt"

// func by2(num int) string {
// 	if num % 2 == 0 {
// 		return "ok"
// 	} else {
// 		return "no"
// 	}
// }

// func main() {
// 	// if result := by2(10); result == "ok" {
// 	// 	fmt.Println("great")
// 	// }

// 	// for i := 0;i < 10; i++{
// 	// 	if i == 3 {
// 	// 		continue
// 	// 	}
// 	// 	if i == 6 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(i)
// 	// }

// 	// sum :=1
// 	// for sum < 10 {
// 	// 	sum += sum
// 	// 	fmt.Println(sum)
// 	// }
// 	// fmt.Println(sum)

// 	// for {
// 	// 	fmt.Println("Loop")
// 	// }

// 	// l := []string{"python", "go", "java"}
// 	// for i := 0; i < len(l); i++ {
// 	// 	fmt.Println(i, l[i])
// 	// }
// 	// for i, v := range l {
// 	// 	fmt.Println(i, v)
// 	// }
// 	// for _, v := range l {
// 	// 	fmt.Println(v)
// 	// }

// 	// m := map[string]int{"apple": 100, "banana": 200}
// 	// for k, v := range m {
// 	// 	fmt.Println(k, v)
// 	// }
// 	// for k := range m {
// 	// 	fmt.Println(k)
// 	// }
// 	// for _,v := range m {
// 	// 	fmt.Println(v)
// 	// }
// }

// import (
// 	"fmt"
// 	"time"
// )

// func getOsName() string {
// 	return "mac"
// }

// func main() {
// 	switch os := getOsName(); os {
// 	case "mac":
// 		fmt.Println("mac")
// 	case "windows":
// 		fmt.Println("windows")
// 	default:
// 		fmt.Println("default")
// 	}

// 	t := time.Now()
// 	fmt.Println(t.Hour())
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Morning")
// 	case t.Hour() < 17:
// 		fmt.Println("Afternoon")
// 	}
// }

// import (
// 	"fmt"
// 	"os"
// )

// func foo() {
// 	defer fmt.Println("world foo")
// 	fmt.Println("hello foo")
// }
// func main() {
// 	// defer fmt.Println("world")
// 	// foo()
// 	// fmt.Println("hello")

// 	// fmt.Println("run")
// 	// defer fmt.Println(1)
// 	// defer fmt.Println(2)
// 	// defer fmt.Println(3)
// 	// fmt.Println("success")

// 	file, _ := os.Open("./lesson.go")
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	file.Read(data)
// 	fmt.Println(string(data))
// }
// import (
// 	"io"
// 	"log"
// 	"os"
// )

// func loggingSettings(logFileName string) {
// 	logFile, _ := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	multiLogFile := io.MultiWriter(os.Stdout, logFile)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
// 	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	log.SetOutput(multiLogFile)
// }
// func main() {
// 	loggingSettings("test.log")
// 	_, err := os.Open("./lesson.go")
// 	if err != nil {
// 		log.Fatalln("Exit", err)
// 	}
// 	log.Println("logging!")
// 	log.Printf("%T %v", "test", "test")
// 	log.Fatalf("%T %v", "test", "test")
// 	log.Fatalln("error")
// 	log.Println("ok!")
// }

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )
// func main() {
// 	file, err := os.Open("lesson.go")
// 	if err != nil {
// 		log.Fatalln("Error")
// 	}
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	count, err := file.Read(data)
// 	if err != nil {
// 		log.Fatalln("Error")
// 	}
// 	fmt.Println(count,string(data))
// }

// import "fmt"
// func thirdPartyConnectDB() {
// 	// panicは原則非推奨。エラー処理を行うべき
// 	panic("Unable to connect database!")
// }
// func save() {
// 	defer func() {
// 		s := recover()
// 		fmt.Println(s)
// 	}()
// 	thirdPartyConnectDB()
// }
// func main() {
// 	save()
// 	fmt.Println("OK?")
// }

import "fmt"
func main() {
	l := []int{100, 300,23,11,23,2,4,6,4}
	max := l[0]
	for _,v := range l{
			if v > max {
				max = v
			}
	}
	fmt.Println(max)

	m := map[string]int{"apple": 200, "banana": 300, "lemon": 150}
	sum := 0
	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)
}
