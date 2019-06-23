package main

func sampleStats(count []int) []float64 {
	if len(count) == 0 {
		return []float64{0, 0, 0, 0, 0}
	}
	values, sum := 0, 0
	maximum, minimum := -256, 256
	maxOccure, mode := 0, 0
	for i := 0; i <= 255; i++ {
		if count[i] > 0 {
			values += count[i]
			sum += i * count[i]
			if count[i] > 0 {
				maximum = i
			}
			if minimum == 256 {
				minimum = i
			}
			if count[i] > maxOccure {
				maxOccure = count[i]
				mode = i
			}
		}
	}
	mean := float64(sum) / float64(values)
	// find median
	// 1 2 3 4 5
	// 1 2 3 4
	midA, midB := (values+1)/2, (values)/2+1
	A, B := 256, 256
	count2 := 0
	for i := 0; i <= 255; i++ {
		if count[i] > 0 {
			count2 += count[i]
			if A == 256 && count2 >= midA {
				A = i
			}
			if B == 256 && count2 >= midB {
				B = i
			}
		}
	}
	median := (float64(A) + float64(B)) / 2
	return []float64{float64(minimum), float64(maximum), mean, median, float64(mode)}
}
