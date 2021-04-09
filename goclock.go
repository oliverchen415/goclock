package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/araddon/dateparse"
	"github.com/cheynewallace/tabby"
)

type Clock struct {
	Datetime string 
	Timezone string
}

func main() {
	area := os.Args[1]
	location := os.Args[2]

	url := fmt.Sprintf("http://worldtimeapi.org/api/timezone/%s/%s", area, location)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var clock Clock
	json.Unmarshal([]byte(body), &clock)
	layout, err := dateparse.ParseAny(clock.Datetime)
	if err != nil {
		log.Fatalln(err)
	}

	t := tabby.New()
	t.AddHeader("DATETIME", "TIMEZONE")
	t.AddLine(layout, clock.Timezone)
	t.Print()
}
