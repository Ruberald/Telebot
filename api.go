package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"encoding/json"
)

var key, _ = os.ReadFile(".apikey")
var reader = bufio.NewReader(os.Stdin)

type Details struct {
   	Title, 
   	Year, 
   	Rated, 
   	Released,
  	Runtime,
  	Genre,
  	Director,
  	Writer,
  	Actors string
 }

func GetInfo(name string) ( string ) {
	// fmt.Println("Enter the name of a film or show:")

	// name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(strings.Replace(name, " ", "+", -1))
	fmt.Println(name)

	resp, _ := http.Get("http://www.omdbapi.com/?apikey=" + string(key) + "&t=" + name + "&r=json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  log.Fatalln(err)
	}


	sb := string(body)
	// sb = strings.ReplaceAll(sb, (sb[strings.Index(sb, "\"Ratings\"") : strings.Index(sb, "],")+2]), "")
	// log.Printf(sb)

	var details Details
	json.Unmarshal([]byte(sb), &details)
	// fmt.Printf("%+v\n", details)

	return fmt.Sprintf("%+v\n", details)
}