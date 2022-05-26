package main

var (
	visited [][]struct{}
	n       int
	m       int
)

func numIslands(grid [][]byte) int {
	n = len(grid)
	visited = make([][]struct{}, n)

	if n == 0 {
		return 0
	}

	m = len(grid[0])
	if m == 0 {
		return 0
	}

	for i := range visited {
		row = make([]struct{})
		visited[i] = row
	}
}

func dfs(grid [][]byte, i int, j int) bool {
	if i < 0 || j < 0 || i <= n || j <= m || grid[i][j] == 0 || visited[i][j] == struct{}{} {
		return false
	}

	visited[i][j] = struct{}{}

	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)

	return true
}
