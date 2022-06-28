package movie

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Movie struct {
	Id    int64   `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Rate  float64 `json:"rate,omitempty"`
	Quote string  `json:"quote,omitempty"`
	Info  string  `json:"info,omitempty"`
}

const (
	Header = `# 豆瓣 TOP Movie 250

> use go native package html to achieve this.

🎁 you can also read [DoubanTOP](https://github.com/Z-Spring/DoubanTOP) which achieves the same feature but use goquery.


Douban top movies from %d to %d.

| Id | Title | Rate | Info | Quote |
| --- | ----- | ---- | ---- | ----- |
`
	Footer = "\n*Last update Time: %v*"
)

func GetMovie(start int) []Movie {
	body := GetMovieBody(start)
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}

	var f func(node *html.Node)
	var (
		movie  []Movie
		movie2 Movie
	)
	pp := make(map[string]bool)

	f = func(n *html.Node) {
		if n.Type == html.TextNode && n.Parent.Type == html.ElementNode && n.Parent.Data == "span" ||
			n.Type == html.TextNode && n.Parent.Type == html.ElementNode && n.Parent.Data == "p" {
			for _, i := range n.Parent.Attr {
				// 获取电影名称
				if i.Val == "title" {
					err := GetZhTitle(n.Data)
					if err != nil {
						continue
					}
					movie2.Name = strings.TrimPrefix(n.Data, " / ")
					//log.Println(strings.TrimPrefix(n.Data, " / "))
				}
				// 获取评分
				if i.Val == "rating_num" {
					data, err := strconv.ParseFloat(n.Data, 10)
					if err != nil {
						log.Println(err)
						return
					}
					movie2.Rate = data
				}
				// 获取评论引用
				if i.Val == "inq" {
					movie2.Quote = n.Data
				}
				if i.Key == "class" && i.Val == "" {
					s := strings.ReplaceAll(n.Data, "\n", "")
					s2 := strings.TrimSpace(s)
					movie2.Info = s2
				}

				// 利用map来达到set集合的效果
				if movie2.Name != "" && movie2.Rate != 0 && movie2.Quote != "" {
					if ok := pp[movie2.Name]; !ok {
						if ok := pp[movie2.Quote]; !ok {
							pp[movie2.Name] = true
							pp[movie2.Quote] = true
							movie = append(movie, movie2)
						}

					}
				}
			}
		}
		// 获取id
		if n.Type == html.TextNode && n.Parent.Type == html.ElementNode && n.Parent.Data == "em" {
			for _, i := range n.Parent.Attr {
				if i.Key == "class" && i.Val == "" {
					id, _ := strconv.ParseInt(n.Data, 10, 32)
					movie2.Id = id
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}
	f(doc)
	return movie
	/*marshal, err := json.MarshalIndent(movie, "\t\t", " ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(marshal))*/
}

func GetZhTitle(title string) error {
	for _, v := range title {
		if !unicode.Is(unicode.Han, v) && !unicode.IsPunct(v) && !unicode.IsSpace(v) && !unicode.IsNumber(v) {
			return errors.New("title只能是中文")
		}
	}
	return nil
}

var ids []int64

func WriteToFile(movie []Movie) error {
	// change path here
	file, err := os.OpenFile("README.md", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	//file, err := os.OpenFile("README.md", os.O_RDWR|os.O_TRUNC, 0666)
	//file, err := os.OpenFile("C:\\Users\\Murphy\\Desktop\\Movie.md", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, movie := range movie {
		ids = append(ids, movie.Id)
	}
	IdStart := ids[0]
	IdEnd := ids[len(ids)-1]

	_, err = file.WriteString(fmt.Sprintf(Header, IdStart, IdEnd))
	if err != nil {
		return err
	}
	for _, movie := range movie {
		_, err := file.WriteString(fmt.Sprintf("| %d | %s | %v/10.0 | %v | %s |\n", movie.Id, movie.Name, movie.Rate, movie.Info, movie.Quote))
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString(fmt.Sprintf(Footer, time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		return err
	}
	return nil
}
