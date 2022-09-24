package main

const N = 10

func main() {

}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		select {
		case x := <-firstChan:
			result <- x * x
		case x := <-secondChan:
			result <- x * 3
		case <-stopChan:
		}
	}()
	return result
}
