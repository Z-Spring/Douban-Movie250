package main

import (
	"awesomeProject3/movie"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	movies := movie.GetMovie(0)
	err := movie.WriteToFile(movies)
	if err != nil {
		log.Println(err)
		return
	}
}
