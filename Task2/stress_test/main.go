package main

import (
	"fmt"
	"net/http"
)

func main() {
	go func() {
		for _ = range 1000 {
			_, err := http.Get("http://insuretech")
			if err != nil {
				fmt.Println(err)
				return
			}
			//
		}
	}()
	go func() {
		for _ = range 1000 {
			_, err := http.Get("http://insuretech")
			if err != nil {
				fmt.Println(err)
				return
			}
			//
		}
	}()
	go func() {
		for _ = range 1000 {
			_, err := http.Get("http://insuretech")
			if err != nil {
				fmt.Println(err)
				return
			}
			//
		}
	}()
	go func() {
		for _ = range 1000 {
			_, err := http.Get("http://insuretech")
			if err != nil {
				fmt.Println(err)
				return
			}
			//
		}
	}()
	go func() {
		for _ = range 1000 {
			_, err := http.Get("http://insuretech")
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}()
	go func() {
		for _ = range 1000 {
			_, err := http.Get("http://insuretech")
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}()

	for _ = range 1000 {
		_, err := http.Get("http://insuretech")
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
