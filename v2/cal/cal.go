package cal

// Add Fucntion to Perform Addition of Two numbers
func Add(args ...int) (sum int) {
	for _, i := range args {
		sum += i
	}
	return
}

// Sub Fucntion to Perform Addition of Two numbers
func Sub(a, b int) int {

	return a - b
}
