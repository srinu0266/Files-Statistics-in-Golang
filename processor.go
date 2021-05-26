package processor

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

type FileInfo struct {
	Name     string      `json:"name"`
	Size     int64       `json:"size"`
	Mode     os.FileMode `json:"mode"`
	Modetime time.Time   `json:"modetime"`
}

func Process(foldername, url string, threads int) {

	regex, err := regexp.Compile("^.*\\.[a-zA-Z0-9]+$")
	if err != nil {
		log.Fatal(err.Error())
	}

	var wg *sync.WaitGroup
	wg = &sync.WaitGroup{}

	ch := make(chan FileInfo)

	wg.Add(1)
	go Thread(url, ch, wg)

	/*
		for i := 0; i < threads; i++ {
			go Thread(url, ch)
		}
	*/

	err = filepath.Walk(foldername, func(path string, info os.FileInfo, err error) error {
		if err == nil && regex.MatchString(info.Name()) {

			fmt.Println("File=", info.Name())
			fileinfo := FileInfo{
				Name:     info.Name(),
				Size:     info.Size(),
				Mode:     info.Mode(),
				Modetime: info.ModTime(),
			}

			ch <- fileinfo

		}
		return nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	close(ch)
	wg.Wait()

}

func PProcess(foldername, url string, threads int) {

	regex, err := regexp.Compile("^.*\\.[a-zA-Z0-9]+$")
	if err != nil {
		log.Fatal(err.Error())
	}

	var wg *sync.WaitGroup
	wg = &sync.WaitGroup{}

	ch := make(chan FileInfo)

	wg.Add(1)
	go Thread(url, ch, wg)

	/*
		for i := 0; i < threads; i++ {
			go Thread(url, ch)
		}
	*/

	err = filepath.Walk(foldername, func(path string, info os.FileInfo, err error) error {
		if err == nil && regex.MatchString(info.Name()) {

			fmt.Println("File=", info.Name())
			fileinfo := FileInfo{
				Name:     info.Name(),
				Size:     info.Size(),
				Mode:     info.Mode(),
				Modetime: info.ModTime(),
			}

			ch <- fileinfo

		}
		return nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	close(ch)
	wg.Wait()

}
