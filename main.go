package main

import (
	"awesomeProject3/movie"
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	start := time.Now()
	movies := movie.GetMovie(0)
	err := movie.WriteToFile(movies)
	if err != nil {
		log.Println(err)
		return
	}
	end := time.Since(start)
	fmt.Println(end)
}
