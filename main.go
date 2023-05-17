package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/pterm/pterm"
)

// Slice of strings with placeholder text.
var fakeInstallList = os.Args[1:]

func getFile(url string) error {
	clinet := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := clinet.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	r := regexp.MustCompile(`\/(?P<filename>[^/]+)$`)
	filename := r.FindStringSubmatch(url)[1]
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func getDownload() {
	// Create progressbar as fork from the default progressbar.
	p, _ := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("下载任务开始。。。").Start()

	for i := 0; i < p.Total; i++ {
		if i == 6 {
			time.Sleep(time.Second * 3) // Simulate a slow download.
		}
		getFile(fakeInstallList[i])
		p.UpdateTitle("下载成功 " + fakeInstallList[i])         // Update the title of the progressbar.
		pterm.Success.Println("下载成功 " + fakeInstallList[i]) // If a progressbar is running, each print will be printed above the progressbar.
		p.Increment()                                       // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
		time.Sleep(time.Millisecond * 350)                  // Sleep 350 milliseconds.
	}
}

func main() {
	getDownload()
}
