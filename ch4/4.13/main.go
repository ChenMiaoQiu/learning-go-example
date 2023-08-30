// go run .\main.go Minions
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseUrl = "http://www.omdbapi.com/?i=tt3896198&apikey=ffa28a90"

func downLoadImg(imgUrl string) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("download error !!!")
		return
	}
	defer resp.Body.Close()

	flie, err := os.Create("output.jpg")
	if err != nil {
		fmt.Println("download error !!!")
		return
	}
	defer flie.Close()

	_, err = io.Copy(flie, resp.Body)
	if err != nil {
		return
	}
}

func getPoster(name string) {
	url := baseUrl + "&t=" + name

	var posterUrl struct {
		Poster string
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("can't find")
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("cant't find")
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&posterUrl); err != nil {
		fmt.Println(err)
		return
	}

	imageUrl := posterUrl.Poster
	downLoadImg(imageUrl)
}

func main() {
	for _, name := range os.Args[1:] {
		getPoster(name)
	}
}
