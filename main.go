package main

import (
	"fmt"
	q "example.com/aqueue/aqueue"
)



func main(){
	fmt.Print("hello world\n");
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/posts",
	}
	
	requestQueue := q.NewRequestQueue(3)
	requestQueue.Run(urls)
	fmt.Print("The end.\n")
	
}