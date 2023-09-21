package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/title":
		click("Title")
	case "/artist":
		click("Artist")
	case "/album":
		click("Album")
	case "/year":
		click("Year")
	case "/length":
		click("Length")

	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, &tracks); err != nil {
		log.Println(err)
	}
}

func click(op string) {
	switch op {
	case "Title":
		sort.Stable(sortOper{
			tracks,
			func(i, j *Track) bool {
				return i.Title < j.Title
			},
			func(x, y *Track) {
				*x, *y = *y, *x
			},
		})
	case "Artist":
		sort.Stable(sortOper{
			tracks,
			func(i, j *Track) bool {
				return i.Artist < j.Artist
			},
			func(x, y *Track) {
				*x, *y = *y, *x
			},
		})
	case "Album":
		sort.Stable(sortOper{
			tracks,
			func(i, j *Track) bool {
				return i.Album < j.Album
			},
			func(x, y *Track) {
				*x, *y = *y, *x
			},
		})
	case "Year":
		sort.Stable(sortOper{
			tracks,
			func(i, j *Track) bool {
				return i.Year < j.Year
			},
			func(x, y *Track) {
				*x, *y = *y, *x
			},
		})
	case "Length":
		sort.Stable(sortOper{
			tracks,
			func(i, j *Track) bool {
				return i.Length < j.Length
			},
			func(x, y *Track) {
				*x, *y = *y, *x
			},
		})

	}
}

type sortOper struct {
	t    []*Track
	less func(i, j *Track) bool
	swap func(x, y *Track)
}

func (s sortOper) Len() int           { return len(s.t) }
func (s sortOper) Less(i, j int) bool { return s.less(s.t[i], s.t[j]) }
func (s sortOper) Swap(i, j int)      { s.swap(s.t[i], s.t[j]) }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Abc", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"bcd", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"efc", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Efc", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
