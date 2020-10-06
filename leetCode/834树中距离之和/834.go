package main

func main() {

}

func sumOfDistancesInTree(n int, edges [][]int) (ans []int) {
	graph := make([][]int, n)
	ans = make([]int, n)

	for _, e := range edges {
		p, q := e[0], e[1]
		graph[p] = append(graph[p], q)
		graph[q] = append(graph[q], p)
	}

	childsNum := make([]int, n)
	nodesDist := make([]int, n)

	var dfs func(curNode, parent int)
	dfs = func(curNode, parent int) {
		childsNum[curNode] = 1
		for _, neighbor := range graph[curNode] {
			if neighbor == parent {
				continue
			}
			dfs(neighbor, curNode)
			childsNum[curNode] += childsNum[neighbor]
			nodesDist[curNode] += nodesDist[neighbor] + childsNum[neighbor]
		}
	}

	var calcNodeDist func(curNode, parent int)
	calcNodeDist = func(curNode, parent int) {
		ans[curNode] = nodesDist[curNode]

		for _, neighbor := range graph[curNode] {
			if neighbor == parent {
				continue
			}
			cn, cd := childsNum[curNode], nodesDist[curNode]
			nn, nd := childsNum[neighbor], nodesDist[neighbor]

			childsNum[curNode] -= childsNum[neighbor]
			nodesDist[curNode] -= nodesDist[neighbor] + childsNum[neighbor]

			childsNum[neighbor] += childsNum[curNode]
			nodesDist[neighbor] += nodesDist[curNode] + childsNum[curNode]

			calcNodeDist(neighbor, curNode)

			childsNum[curNode], nodesDist[curNode] = cn, cd
			childsNum[neighbor], nodesDist[neighbor] = nn, nd
		}
	}

	dfs(0, -1)
	calcNodeDist(0, -1)

	return
}
