package processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func Thread(url string, ch chan FileInfo, wg *sync.WaitGroup) {

L:
	for {
		select {
		case fileinfo, ok := <-ch:
			if !ok {
				fmt.Println("channel hasbeen closed!")
				break L
			}
			Task(fileinfo, url)
		}
	}
	wg.Done()

}

func Task(fileinfo FileInfo, url string) {
	response, err := json.Marshal(fileinfo)
	if err != nil {
		fmt.Println("Error While Marshaling", fileinfo)
		return
	}

	responseBody := bytes.NewBuffer(response)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		fmt.Println("Error While Post Request", fileinfo)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error While Reading Post Response", fileinfo)
		return
	}
	sb := string(body)
	fmt.Printf("Response from  %s===>%s\n", url, sb)

}
