package main

import (
	"fmt"
	"sync"
)

// Слить N каналов в один. Даны n каналов типа chan int. Надо написать функцию, которая смерджит все данные из этих
// каналов в один и вернет его.

func merge(chs ...<-chan int) <-chan int {
	ret := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for item := range ch {
				ret <- item
			}
		}()
	}

	go func() {
		wg.Wait()

		close(ret)
	}()

	return ret
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, item := range []int{1, 2, 3} {
			a <- item
		}

		close(a)
	}()

	go func() {
		for _, item := range []int{20, 10, 30} {
			b <- item
		}

		close(b)
	}()

	go func() {
		for _, item := range []int{300, 200, 100} {
			c <- item
		}

		close(c)
	}()

	for num := range merge(a, b, c) {
		fmt.Println(num)
	}
}
