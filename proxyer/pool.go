package proxyer

func Pooldigger(concurrent int) chan int {

	pool := make(chan int, concurrent)

	for x := 0; x < concurrent; x++ {
		pool <- 1
	}

	return pool

}
