package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Xkcd struct {
	Title string
	Url   string
	Img   string
}

const baseUrl = "https://xkcd.com/"

func getOfflineIndex() {
	out, err := os.Create("offlineIndex.txt")
	if err != nil {
		fmt.Println("error !!!")
		return
	}
	defer out.Close()

	write := csv.NewWriter(out)
	write.Write([]string{"title", "link", "img"})

	for i := 1; i <= 120; i++ {
		xkcd := &Xkcd{}

		url := fmt.Sprintf("%s%d/info.0.json", baseUrl, i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return
		}
		if err := json.NewDecoder(resp.Body).Decode(xkcd); err != nil {
			resp.Body.Close()
			fmt.Println(err)
		}

		resp.Body.Close()
		write.Write([]string{xkcd.Title, xkcd.Url, xkcd.Img})
		write.Flush()
	}

}

func resLink(target string) {
	f, err := os.Open("./offlineIndex.txt")

	if err != nil {
		fmt.Println("can't find offline index")
		return
	}
	defer f.Close()
	reader := csv.NewReader(f)

	for {
		res, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if target == res[0] {
			fmt.Println(res[2])
			return
		}
	}
	fmt.Println("can't find !!!")
}

func main() {
	f, err := os.Open("./offlineIndex.txt")
	if err != nil {
		getOfflineIndex()
	}
	f.Close()

	for _, val := range os.Args[1:] {
		resLink(val)
	}
}
