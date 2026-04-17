// Найти все ошибки,даже если просто не нравится как написано
// Задача от компании ВБ
package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething() ([]string, error) {
    var res []string 
	err := runFunc(func() {
		for i := 0; i < 5; i++ {
			res[i] = fmt.Sprintf("result %d", i)
		}
	})
	return res, err
}

func runFunc(f func()) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan bool, 1)

	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errCh <- fmt.Errorf("panic: %v", r)
			}
		}()

		f()
		done <- true
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case err := <-errCh:
		return err
	case <-time.After(3 * time.Second):
		return fmt.Errorf("timeout")
	}
}

func main() {
	res, err := doSomething()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(res)
}