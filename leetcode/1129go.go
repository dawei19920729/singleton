package main

/*
在一个有向图中，节点分别标记为   0, 1, ..., n-1。图中每条边为红色或者蓝色，且存在自环或平行边。

red_edges   中的每一个   [i, j]   对表示从节点 i 到节点 j 的红色有向边。类似地，blue_edges   中的每一个   [i, j]   对表示从节点 i 到节点 j 的蓝色有向边。

返回长度为 n 的数组   answer，其中   answer[X]   是从节点   0   到节点   X   的红色边和蓝色边交替出现的最短路径的长度。如果不存在这样的路径，那么 answer[x] = -1。

输入：n = 3, red_edges = [[0,1],[1,2]], blue_edges = []
输出：[0,1,-1]
输入：n = 3, red_edges = [[0,1]], blue_edges = [[2,1]]
输出：[0,1,-1]
n = 3, red_edges = [[0,1]], blue_edges = [[1,2]]
输出：[0,1,2]
输入：n = 3, red_edges = [[0,1],[0,2]], blue_edges = [[1,0]]
输出：[0,1,1]
*/

func main() {

}

/*
使用 00 表示红色，11 表示蓝色，对于某一节点 xx，从节点 00 到节点 xx 的红色边和蓝色边交替出现的路径有两种类型：
类型 00：路径最终到节点 xx 的有向边为红色；
类型 11：路径最终到节点 xx 的有向边为蓝色。
如果存在从节点 00 到节点 xx 的类型 00 的颜色交替路径，并且从节点 xx 到节点 yy 的有向边为蓝色，那么该路径加上该有向边组成了从节点 00 到节点 yy 的类型 11 的颜色交替路径。
类似地，如果存在从节点 00 到节点 xx 的类型 11 的颜色交替路径，并且从节点 xx 到节点 yy 的有向边为红色，那么该路径加上该有向边组成了从节点 00 到节点 yy 的类型 00 的颜色交替路径。
使用广度优先搜索获取从节点 00 到某一节点的两种类型的颜色交替最短路径的长度，广度优先搜索的队列元素由节点编号和节点路径类型组成，
初始时节点 00 到节点 00 的两种类型的颜色交替最短路径的长度都是 00，将两个初始值入队。对于某一个队列元素，节点编号为 xx，节点路径类型为 tt，
那么根据类型 tt 选择颜色为 1-t1−t 的相邻有向边，如果有向边的终点节点 yy 对应类型 1-t1−t 没有被访问过，
那么更新节点 yy 的类型 1-t1−t 的颜色交替最短路径的长度为节点 xx 的类型 tt 的颜色交替最短路径的长度加 11，并且将它入队。
从节点 00 到某一节点的颜色交替最短路径的长度为两种类型的颜色交替最短路径长度的最小值。
*/
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	type pair struct{ x, color int }
	g := make([][]pair, n)
	for _, e := range redEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 0})
	}
	for _, e := range blueEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 1})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	vis := make([][2]bool, n)
	vis[0] = [2]bool{true, true}
	q := []pair{{0, 0}, {0, 1}}
	for level := 0; len(q) > 0; level++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			x := p.x
			if dis[x] < 0 {
				dis[x] = level
			}
			for _, to := range g[x] {
				if to.color != p.color && !vis[to.x][to.color] {
					vis[to.x][to.color] = true
					q = append(q, to)
				}
			}
		}
	}
	return dis
}
