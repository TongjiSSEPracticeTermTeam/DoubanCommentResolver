package main

import (
	"bufio"
	"douban_comment_resolver/pkg/resolver"
	"flag"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
)

var inPath = flag.String("in", "", "The path to the Douban comment page")
var outPath = flag.String("out", "", "The path to write the JSON")
var movieName = flag.String("name", "", "The name of the movie")
var movieId = flag.String("id", "", "The id of the movie")

func main() {
	flag.Parse()

	f, err := os.OpenFile(*inPath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}(f)
	r := bufio.NewReader(f)

	o, err := resolver.ResolveComments(r)
	if err != nil {
		log.Fatalln(err.Error())
	}

	out := make(map[string]interface{})
	out["name"] = movieName
	out["id"] = movieId
	out["data"] = o

	json, err := jsoniter.Marshal(out)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = os.WriteFile(*outPath, []byte(json), 0666)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
