package main

func topoSort(graph [][]int, inDegree []int, items []int) (res []int) {
	s := []int{}
	for _, item := range items {
		if inDegree[item] == 0 {
			s = append(s, item)
		}
	}
	for len(s) > 0 {
		item := s[0]
		s = s[1:]
		res = append(res, item)
		for _, i := range graph[item] {
			inDegree[i]--
			if inDegree[i] == 0 {
				s = append(s, i)
			}
		}
	}
	return
}

func sortItems(n int, m int, group []int, beforeItems [][]int) (ans []int) {
	groupItems := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 { // 给没有分配小组的项目建立一个新组
			group[i] = m + i
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupGraph := make([][]int, m+n)
	groupInDeg := make([]int, m+n)
	itemGraph := make([][]int, n)
	itemInDeg := make([]int, n)

	// 建立组依赖图和任务依赖图
	for cur, preList := range beforeItems {
		curGroup := group[cur]
		for _, pre := range preList {
			preGroup := group[pre]
			if preGroup != curGroup {
				groupGraph[preGroup] = append(groupGraph[preGroup], curGroup)
				groupInDeg[curGroup]++
			} else {
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemInDeg[cur]++
			}
		}
	}

	// 对组进行拓扑排序
	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}
	groupOrder := topoSort(groupGraph, groupInDeg, items)
	if len(groupOrder) < m+n {
		return nil // 拓扑排序后的结果小于原来的总数，代表不能完成拓扑排序
	}

	// 对组内排序
	for _, group := range groupOrder {
		itemOrder := topoSort(itemGraph, itemInDeg, groupItems[group])
		if len(itemOrder) < len(groupItems[group]) {
			return nil
		}
		ans = append(ans, itemOrder...)
	}
	return
}
