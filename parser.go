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
	team   = make([]string, 1000)
	score  = make([]string, 1000)
	length int
	today  string
)

func main() {
	today = today_ymd(time.Now())
	extracter("http://www.espnfc.com/scores", ".team-name", ".team-score", today)
	for i := 0; i <= length; i += 2 {
		fmt.Printf(team[i] + " vs " + team[i+1])
		fmt.Printf("\t" + score[i] + ":" + score[i+1] + "\n")
	}
}

func extracter(url string, team_class string, score_class string, date string) {
	doc, err := goquery.NewDocument(url + "?date=" + date)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(team_class).Each(func(i int, s *goquery.Selection) {
		//class, _ := s.Attr("class")
		//fmt.Println(i, s.Text())
		length = i
		team[i] = strings.Replace(s.Text(), "\n", "", -1)
		team[i] = strings.Replace(team[i], "\t", "", -1)
		//fmt.Println(reflect.TypeOf(s.Text()))
	})

	doc.Find(score_class).Each(func(i int, s *goquery.Selection) {
		score[i] = strings.Replace(s.Text(), "\n", "", -1)
		score[i] = strings.Replace(score[i], "\t", "", -1)
		//fmt.Println(i, s.Text())
	})
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
