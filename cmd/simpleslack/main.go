package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/justone/simpleslack"
)

type SliceValue struct {
	contents []string
}

func (sv *SliceValue) String() string {
	return "slice of values"
}

func (sv *SliceValue) Set(newval string) error {
	sv.contents = append(sv.contents, newval)
	return nil
}

func main() {
	var params SliceValue
	var token string
	var method string
	flag.Var(&params, "v", "parameter values")
	flag.StringVar(&token, "t", os.Getenv("SLACK_TOKEN"), "slack token")
	flag.StringVar(&method, "m", "", "api method")

	flag.Parse()
	s := simpleslack.Client{token}

	values, err := url.ParseQuery(strings.Join(params.contents, "&"))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result, err := s.Post(method, values)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	byt, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(string(byt))
}
