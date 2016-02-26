package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	team     = make([]string, 1000)
	score    = make([]string, 1000)
	length   int
	today    string
	_team_no int
)

func main() {
	today = today_ymd(time.Now())
	//team_no = 59
	extracter_rank("https://godoc.org/github.com/PuerkitoBio/goquery#Selection", today)
	/*	for i := 0; i <= length; i += 2 {
			fmt.Printf(team[i] + " vs " + team[i+1])
			fmt.Printf("\t" + score[i] + ":" + score[i+1] + "\n")
		}
	*/
}

func extracter_rank(url string, today string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	//doc_parse, err := html.Parse(strings.NewReader(doc))
	doc.Find(".funcdecl decl").FindNodes().Each(func(i int, s *goquery.Selection) {
		fmt.Printf(strings.Replace(s.Text(), "\n", "", -1))
	})

	/*
		doc.Find().Each(func(i int, s *goquery.Selection) {
			team[i] = strings.Replace(s.Text(), "\n", "", -1)
			team[i] = strings.Replace(team[i], "\t", "", -1)
		})
	*/

}

func today_ymd(t time.Time) string {
	year := strconv.Itoa(t.Year())
	m := int(t.Month())
	d := t.Day()
	var month string
	var day string

	if m < 10 {
		month = "0" + strconv.Itoa(m)
	} else {
		month = strconv.Itoa(m)
	}
	if d < 10 {
		day = "0" + strconv.Itoa(d)
	} else {
		day = strconv.Itoa(d)
	}
	return year + month + day
}
