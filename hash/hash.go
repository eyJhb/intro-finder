package hash

func Diff(hash1, hash2 uint64) uint8 {
	// holds the distance between the two
	var distance uint8

	// init our mask, and set initial value of 1
	var mask uint64
	mask |= 1

	for i := 0; i < 64; i++ {
		// mask our hash1 and hash2, see if they are different
		// if se increment the distance
		if hash1&mask != hash2&mask {
			distance++
		}

		// move our mask the the left
		mask <<= 1
	}

	return distance
}
