package _994

type Point struct {
	X int
	Y int
}

type Item struct {
	Time  int
	Point Point
}
type Queue []Item

func (q *Queue) Push(value Item) {
	*q = append(*q, value)
}

func (q *Queue) Pop() Item {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func orangesRotting(grid [][]int) int {
	queue := make(Queue, 0)
	fresh := 0
	time := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fresh = fresh + 1
			} else if grid[i][j] == 2 {
				queue.Push(Item{
					Time:  0,
					Point: Point{X: i, Y: j},
				})
			}
		}
	}

	directions := []Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for !queue.Empty() {
		var data = queue.Pop()
		time = data.Time
		for i := 0; i < len(directions); i++ {
			var direction = directions[i]
			var row = data.Point.X + direction.X
			var col = data.Point.Y + direction.Y

			if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] != 1 {
				continue
			}

			fresh = fresh - 1
			grid[row][col] = 2
			queue.Push(Item{
				Time:  data.Time + 1,
				Point: Point{X: row, Y: col},
			})
		}
	}

	if fresh != 0 {
		return -1
	}
	return time
}
