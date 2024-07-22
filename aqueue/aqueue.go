package aqueue

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type RequestQueue struct {
	nbWorkers int
	queue chan string
	wg sync.WaitGroup
}


func NewRequestQueue(nbWorkers int) *RequestQueue {
	return &RequestQueue {
		nbWorkers: nbWorkers,
		queue: make(chan string),
	}
}

func (rq *RequestQueue) addRequest(url string){
	rq.wg.Add(1)
	rq.queue <- url
}

func (rq *RequestQueue) worker(){
	for url := range rq.queue {
		rq.fetch(url)
		rq.wg.Done()
	}
}

func (rq *RequestQueue) fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error reading response from %s: %v\n", url, err)
		return
	}
	fmt.Printf("Fetched data from %s: %s...\n\n", url, string(body)[:100])
}


func (rq *RequestQueue) Run(urls []string){
	for i := 0; i < rq.nbWorkers; i++ {
		go rq.worker()
	}
	
	for _, url := range urls {
		rq.addRequest(url)
	}
	rq.wg.Wait()
	close(rq.queue)
}