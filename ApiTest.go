// ApiTest
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	_"strings"
	_"text/scanner"
)

func main() {
	req, err := http.NewRequest("GET", "http://football-api.com/api/?Action=standings&APIKey=76fa3dd7-dbc9-984e-073295e09af1&comp_id=1204", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("StatusCode:", res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	//	fmt.Println(data)
	fmt.Println(reflect.TypeOf(body))
	fmt.Println(reflect.TypeOf(data))
	//fmt.Println(data["stand_position"])
	//if(reflect.TypeOf(body))
	/*
	const src = data
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		fmt.Println("At position", s.Pos(), ":", s.TokenText())
	}
	*/
}
