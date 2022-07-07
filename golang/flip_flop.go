package main

func flipFlop(number, key1, key2 int) (res map[int]string) {
	res = map[int]string{}

	// steps :
	// loops in number of times
	// validation :
	// - if n mod (key1 * key2) == 0 add value  "flip flop" with key `n` to res
	// - if n mod key1 == 0 add value "flip" with key `n` to res
	// - if n mod key2 == 0 add value  "flop" with key `n` to res
	for n := 1; n <= number; n++ {
		if n%(key1*key2) == 0 {
			res[n] = "flip flop"
		} else if n%(key1) == 0 {
			res[n] = "flip"
		} else if n%(key2) == 0 {
			res[n] = "flop"
		}
	}

	return res
}
