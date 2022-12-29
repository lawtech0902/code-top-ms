package makingalargeisland

func largestIsland(grid [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // 4个方向
	m, n := len(grid), len(grid[0])
	gridNum := make(map[int]int)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	mark := 2         // 记录每个island的编号
	isAllGrid := true // 标记是否整个grid都是陆地
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				isAllGrid = false
			}

			if !visited[i][j] && grid[i][j] == 1 {
				count = 0
				dfs(&grid, &visited, i, j, mark, &count, dirs) // 将与其连接的陆地标记为true
				gridNum[mark] = count                          // 记录每个island的面积
				mark++                                         // 记录下一个岛屿编号
			}
		}
	}

	if isAllGrid {
		return m * n // 如果全是陆地，则返回全面积
	}

	// 以下逻辑是根据添加陆地的位置，计算周边岛屿面积之和
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count = 1 // 记录连接之后的岛屿数量
			if grid[i][j] == 0 {
				visitGrid := make(map[int]bool) // 标记访问过的岛屿
				for k := 0; k < 4; k++ {
					newX := i + dirs[k][0]
					newY := j + dirs[k][1]
					// 添加过的岛屿不要重复添加
					if newX < 0 || newX >= m || newY < 0 || newY >= n || visitGrid[grid[newX][newY]] {
						continue
					}

					// 把相邻四面的岛屿数量加起来
					count += gridNum[grid[newX][newY]]
					// 标记该岛屿已经添加过
					visitGrid[grid[newX][newY]] = true
				}
				res = max(res, count)
			}
		}
	}

	return res
}

func dfs(grid *[][]int, visited *[][]bool, x, y, mark int, count *int, dirs [][]int) {
	// 终止条件：访问过的节点或遇到海水
	if (*visited)[x][y] || (*grid)[x][y] == 0 {
		return
	}

	(*visited)[x][y] = true // 标记访问过
	(*grid)[x][y] = mark    // 给陆地标记新标签
	(*count)++
	for i := 0; i < 4; i++ {
		nextX := x + dirs[i][0]
		nextY := y + dirs[i][1]
		if nextX < 0 || nextX >= len(*grid) || nextY < 0 || nextY >= len((*grid)[0]) {
			continue
		}

		dfs(grid, visited, nextX, nextY, mark, count, dirs)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
