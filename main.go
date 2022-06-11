package main

import (
	"awesomeProject3/movie"
	"fmt"
	"io"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

type Student struct {
	Name string
	Age  int
}

type Node struct {
	Data     int
	NextNode *Node
}

var head = new(Node)
var l, r []int

func main() {
	movies := movie.GetMovie(200)
	err := movie.WriteToFile(movies)
	if err != nil {
		log.Println(err)
		return
	}
	//movie.GetMovie()
	/*s := "哈利·波特与阿兹卡班的囚徒"
	for _, i := range s {

		fmt.Println(unicode.IsPunct(i))
	}*/
	/*for _, i := range s {


	}*/
}

func MaoPao(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}

	}

}

func QuickSort(a []int) {
	a1 := a[0]
	for i := 1; i < len(a); i++ {
		if a1 > a[i] {
			l = append(l, a[i])
			fmt.Println(l)
		} else if a1 < a[i] {
			r = append(r, a[i])
		}

	}

	//QuickSort(l)
	/*	fmt.Println(l)
		fmt.Println(r)*/

}

func addNode(t *Node, v int) int {
	if head == nil {
		head = &Node{Data: v, NextNode: nil}
		return 0
	}
	if v == t.Data {
		fmt.Println("节点已经存在", v)
		return -2
	}

	if t.NextNode == nil {
		t.NextNode = &Node{Data: v, NextNode: nil}
		return -1
	}
	return addNode(t.NextNode, v)
}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
		next.ServeHTTP(w, r)
	}
}

func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	middle := array[mid]
	if middle == target {
		return mid
	} else if middle > target {
		return BinarySearch(array, target, 0, mid-1)
	} else if middle < target {
		return BinarySearch(array, target, mid+1, r)
	}
	return -1

}
