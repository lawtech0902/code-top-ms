package game

import "math"

const (
	TARGET                          = 24
	EPSILON                         = 1e-6 // 误差
	ADD, MULTIPLY, SUBTRACT, DIVIDE = 0, 1, 2, 3
)

// 回溯
func judgePoint24(cards []int) bool {
	float64Cards := make([]float64, 0)
	for _, card := range cards {
		float64Cards = append(float64Cards, float64(card))
	}

	return backTracking(float64Cards)
}

func backTracking(cards []float64) bool {
	if len(cards) == 0 {
		return false
	}

	if len(cards) == 1 {
		return math.Abs(cards[0]-TARGET) < EPSILON
	}

	size := len(cards)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i != j {
				track := make([]float64, 0)
				for k := 0; k < size; k++ {
					if k != i && k != j {
						track = append(track, cards[k])
					}
				}

				for k := 0; k < 4; k++ {
					if k < 2 && i < j {
						continue
					}

					switch k {
					case ADD:
						track = append(track, cards[i]+cards[j])
					case MULTIPLY:
						track = append(track, cards[i]*cards[j])
					case SUBTRACT:
						track = append(track, cards[i]-cards[j])
					case DIVIDE:
						if math.Abs(cards[j]) < EPSILON {
							continue
						} else {
							track = append(track, cards[i]/cards[j])
						}
					}

					if backTracking(track) {
						return true
					}

					track = track[:len(track)-1]
				}
			}
		}
	}

	return false
}
