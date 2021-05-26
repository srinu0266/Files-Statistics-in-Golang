package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

//response file structure
type FileStats struct {
	FilesReceived   int
	MaxFileSize     int
	AvgFileSize     float64
	FileExt         []string
	FrequentFile    map[string]int
	LatestFilePaths []string
}

//request file structure
type FileInfo struct {
	Name     string      `json:"name"`
	Size     int64       `json:"size"`
	Mode     os.FileMode `json:"mode"`
	Modetime time.Time   `json:"modetime"`
}

var filestats FileStats
var lock sync.Mutex
var frequentfile = make(map[string]int)

func ProcessFile(w http.ResponseWriter, r *http.Request) {

	var p FileInfo

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		log.Printf("Error reading Body %s ", err.Error())
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	UpdateStats(p)

	fmt.Fprintf(w, "Record Updated")

}

func UpdateStats(p FileInfo) {
	lock.Lock()
	defer lock.Unlock()

	filestats.FilesReceived += 1

	//set max file size
	if filestats.MaxFileSize < int(p.Size) {
		filestats.MaxFileSize = int(p.Size)
	}

	//avg file size
	filestats.AvgFileSize = (filestats.AvgFileSize+float64(p.Size))/float64(filestats.FilesReceived) + 1

	//findout file extension
	ss := strings.Split(p.Name, ".")
	ext := ss[len(ss)-1]

	set := make(map[string]struct{})
	for _, v := range filestats.FileExt {
		set[v] = struct{}{}
	}

	if _, ok := set[ext]; ok {

	} else {
		filestats.FileExt = append(filestats.FileExt, ext)
	}
	frequentfile[ext] += 1

	filestats.FrequentFile = make(map[string]int)

	for k := range filestats.FrequentFile {
		delete(filestats.FrequentFile, k)
	}

	count := 0
	fileext := ext
	for k, v := range frequentfile {
		if v > count {
			count = v
			fileext = k
		}
	}
	filestats.FrequentFile[fileext] = count

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	} else {

		if len(filestats.LatestFilePaths) == 0 {
			filestats.LatestFilePaths = append(filestats.LatestFilePaths, mydir)
		}
	}

}

//get all file statistics
func GetStatistics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filestats)

}
