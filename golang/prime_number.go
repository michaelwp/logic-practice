package main

func PrimeNumber(limit int) []int {
	var res = make([]int, 0)

	for n := 2; n < limit; n++ {
		if !isPrime(n) {
			continue
		}

		res = append(res, n)
	}

	return res
}

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
