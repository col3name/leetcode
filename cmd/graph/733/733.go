package _733

type Node struct {
	sr int
	sc int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]

	if color == newColor {
		return image
	}

	q := make([]Node, 0)
	q = append(q, Node{
		sr: sr,
		sc: sc,
	})
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		sr = current.sr
		sc = current.sc
		if image[sr][sc] == color {
			image[sr][sc] = newColor
			if sr >= 1 {
				q = append(q, Node{sr: sr - 1, sc: sc})
			}
			if sc >= 1 {
				q = append(q, Node{sr: sr, sc: sc - 1})
			}
			if sr+1 < len(image) {
				q = append(q, Node{sr: sr + 1, sc: sc})
			}
			if sc+1 < len(image[0]) {
				q = append(q, Node{sr: sr, sc: sc + 1})
			}
		}
	}

	return image
}
