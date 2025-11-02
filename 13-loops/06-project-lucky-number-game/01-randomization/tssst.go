	rand.Seed(time.Now().UnixNano())
	 

	
	upper := 999999999999999999
	target := 97389682632
	i := 1
	n := 0

	for {
		mid := (upper + n) / 2
		if mid == target {
			fmt.Printf("Success, target found! Guess: %d Target: %d Turn: %d \n", n, target, i)
			break
		}
		if mid <= target {
			n = mid
		} else {
			upper = mid
		}
		i++
	
