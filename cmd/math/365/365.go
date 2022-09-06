package _365

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity+jug2Capacity == targetCapacity {
		return true
	}

	var curr int
	for i := 0; i < jug1Capacity; i++ {
		curr = (jug2Capacity * i) % jug1Capacity
		if curr == targetCapacity || curr+jug2Capacity == targetCapacity {
			return true
		}
	}
	for i := 0; i < jug2Capacity; i++ {
		curr = (jug1Capacity * i) % jug2Capacity
		if curr == targetCapacity || curr+jug1Capacity == targetCapacity {
			return true
		}
	}

	return false
}
