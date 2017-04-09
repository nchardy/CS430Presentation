package main

import (
	"fmt"
	"time"
	"sync"
)

type Job struct {
  title string
  finished bool
}

var items = [...]string{"Milk","Cheese","Napkins","PopTarts","Chips","Salad","Pizza","Rice","Chicken","Soap"}
var children = [...]string{"Frank", "Jill", "Constantine"}
var wg sync.WaitGroup

// START1 OMIT
func worker(name string, jobs <-chan Job, completed chan<- Job){ // HL
	for j := range jobs { // HL
		fmt.Println("-->", name, "Went to get the", j.title)
		j.finished = true 
		time.Sleep(time.Second)
		fmt.Println("<--", name, "Finished getting the", j.title)
		completed <- j // HL
		wg.Done() // HL
	} // HL
} // HL
// END1 OMIT

// START2 OMIT
func main(){
	jobs := make(chan Job, len(items)) // HL
	completed := make(chan Job, len(items)) // HL
	for w:=0; w<=len(children)-1; w++ {
		go worker(children[w], jobs, completed) // HL
	}
	for j:=0; j<=len(items)-1; j++ {
		wg.Add(1) // HL
		job := Job{items[j], false}
		jobs<-job // HL
	}
	wg.Wait() // HL
	close(jobs) // HL
	fmt.Println("All Done.  Let's do inventory now:")
	time.Sleep(time.Second)
	for a:=0; a<=len(items)-1; a++ {
		item := <-completed
		if item.finished{
			fmt.Println("✓", item.title)
		}else{
			fmt.Println("x", item.title)
		}
	}
}
// END2 OMIT