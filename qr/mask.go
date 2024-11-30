package qr

import "math"

// getScoreRule1 вычисляем балы для выбора маски по Правилу 1:
// горизонтальный подсчет
func getScoreRule1(canvas *[][]bool) int {
	lengthCanvas := len(*canvas)
	rule1Score := 0
	for i := 0; i < lengthCanvas; i++ {
		color := (*canvas)[i][0]
		cnt := 1
		for j := 1; j < lengthCanvas; j++ {
			if color == (*canvas)[i][j] {
				cnt++
			} else {
				color = (*canvas)[i][j]
				if cnt >= 5 {
					rule1Score += (cnt - 2)
				}
				cnt = 1
			}
		}
		if cnt >= 5 {
			rule1Score += cnt - 2
		}
	}
	return rule1Score
}

// getScoreRule2 вычисляем балы для выбора маски по Правилу 2:
// вертикальный подсчет
func getScoreRule2(canvas *[][]bool) int {
	lengthCanvas := len(*canvas)
	rule2Score := 0
	for j := 0; j < lengthCanvas; j++ {
		color := (*canvas)[0][j]
		cnt := 1
		for i := 1; i < lengthCanvas; i++ {
			if color == (*canvas)[i][j] {
				cnt++
			} else {
				color = (*canvas)[i][j]
				if cnt >= 5 {
					rule2Score += (cnt - 2)
				}
				cnt = 1
			}
		}
		if cnt >= 5 {
			rule2Score += cnt - 2
		}
	}
	return rule2Score
}

// getScoreRule3 вычисляем балы для выбора маски по Правилу 3:
// За каждый квадрат модулей одного цвета размером 2 на 2 начисляется по 3 очка.
func getScoreRule3(canvas *[][]bool) int {
	lengthCanvas := len(*canvas)
	rule3Score := 0
	for i := 0; i < lengthCanvas-1; i++ {
		for j := 0; j < lengthCanvas-1; j++ {
			if (*canvas)[i][j] == (*canvas)[i+1][j] &&
				(*canvas)[i][j] == (*canvas)[i][j+1] &&
				(*canvas)[i][j] == (*canvas)[i+1][j+1] {
				rule3Score += 3
			}
		}
	}
	return rule3Score
}

// getScoreRule4 вычисляем балы для выбора маски по Правилу 4:
// За каждую последовательность по горизонтали модулей ЧБЧЧЧБЧ,
// с 4-мя белыми модулями с одной из сторон (или с 2-х сразу),
// добавляется 40 очков
func getScoreRule4(canvas *[][]bool) int {
	lengthCanvas := byte(len(*canvas))
	bitRule := 0b11111111111
	ruleL := 0b00001011101
	ruleR := 0b10111010000
	m := map[[2]byte]struct{}{}
	var i, j byte
	for i = 0; i < lengthCanvas; i++ {
		x := 0
		for j = 0; j < lengthCanvas; j++ {
			x <<= 1
			if (*canvas)[i][j] {
				x += 1
			}
			if j > 9 && x&bitRule == ruleL {
				m[[2]byte{i, j - 3}] = struct{}{}
			}
			if x&bitRule == ruleR {
				m[[2]byte{i, j - 7}] = struct{}{}
			}
		}
	}
	return len(m) * 40
}

// getScoreRule5 вычисляем балы для выбора маски по Правилу 4:
// За каждую последовательность по вертикали модулей ЧБЧЧЧБЧ,
// с 4-мя белыми модулями с одной из сторон (или с 2-х сразу),
// добавляется 40 очков
func getScoreRule5(canvas *[][]bool) int {
	lengthCanvas := byte(len(*canvas))
	bitRule := 0b11111111111
	ruleL := 0b00001011101
	ruleR := 0b10111010000
	m := map[[2]byte]struct{}{}
	var i, j byte
	for j = 0; j < lengthCanvas; j++ {
		x := 0
		for i = 0; i < lengthCanvas; i++ {
			x <<= 1
			if (*canvas)[i][j] {
				x += 1
			}
			if i > 9 && x&bitRule == ruleL {
				m[[2]byte{j, i - 3}] = struct{}{}
			}
			if x&bitRule == ruleR {
				m[[2]byte{j, i - 7}] = struct{}{}
			}
		}
	}
	return len(m) * 40
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

// getScoreRule6 вычисляем балы по соотношения количества чёрных и белых модулей
func getScoreRule6(canvas *[][]bool) int {
	lengthCanvas := byte(len(*canvas))
	cntI := 0
	var i, j byte
	for i = 0; i < lengthCanvas; i++ {
		for j = 0; j < lengthCanvas; j++ {
			if (*canvas)[i][j] {
				cntI++
			}
		}
	}
	return abs(int(math.Floor(float64(cntI*100)/float64(lengthCanvas)/float64(lengthCanvas)-50.0)) * 2)
}

// getScore вычисляем балы для выбора маски по горизонтали
func getScore(canvas *[][]bool) int {
	return getScoreRule1(canvas) +
		getScoreRule2(canvas) +
		getScoreRule3(canvas) +
		getScoreRule4(canvas) +
		getScoreRule5(canvas) +
		getScoreRule6(canvas)
}
