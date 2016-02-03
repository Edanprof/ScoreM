package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	team   = make([]string, 1000)
	score  = make([]string, 1000)
	length int
)

func main() {
	extracter("http://www.espnfc.com/scores", ".team-name", ".team-score", 20160130)
	for i := 0; i <= length; i += 2 {
		fmt.Printf(team[i] + " vs " + team[i+1])
		fmt.Printf("\t" + score[i] + ":" + score[i+1] + "\n")
	}
}

func extracter(url string, team_class string, score_class string, date int) {
	doc, err := goquery.NewDocument(url + "?date=" + strconv.Itoa(date))
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
