package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("Image")
)

type Search func(query string) Result
type Result struct {
	Title, URL string
}

func main() {
	fmt.Println("sync")
	syncResult := CallWebpage("hi")
	fmt.Println(syncResult)

	fmt.Println("goResult")
	goResult := GoRoutineWebpage("hi")
	fmt.Println(goResult)

}
func GoRoutineWebpage(query string) (results []Result) {

	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()

	for i := 0; i < 1; i++ {
		result := <-c
		results = append(results, result)
	}

	return results
}

func CallWebpage(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	return results
}
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("%s result for %q \n", kind, query)
		return Result{kind, "url" + kind}
	}
}
