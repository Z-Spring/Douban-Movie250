package movie

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetMovieBody(start int) []byte {
	//url := "https://movie.douban.com/top250"
	url := ChooseStartId(start)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)

	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}
	return body
}

func ChooseStartId(start int) string {
	if start >= 0 && start <= 249 {
		url := "https://movie.douban.com/top250?start=%d&filter="
		url2 := fmt.Sprintf(url, start)
		return url2
	}
	return "The start is from 0~249"
}
