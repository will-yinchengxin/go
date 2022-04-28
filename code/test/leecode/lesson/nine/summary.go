package nine

/*
- 图
	1) 常用概念: 顶点 边 无向图 有向图 有权图 度 入度 出度
		- 无向无权图
		- 无向有权图
		- 有向无权图
		- 有向有权图
		图.png

	树是图的一种特殊情况

	2) 图的存储方式
		存是为了方便用
			- 检查两顶点之间是否有边
			- 获取两顶点之间边的权重
			- 获取某个顶点相连的所有 边/顶点
		图的存储方式.png

		邻接矩阵: 存储稀疏图比较浪费空间, 数据的访问效率比较高 (空间换时间)
		邻接表:	 适合存储稀疏图, 访问两个顶点之间是否有边, 需要遍历链表 (时间换空间)

		//--------------------- 有向无权图(邻接矩阵) ---------------------
		import "container/list"

		var v int
		var matrix[][]bool
		func graph(val int) {
			v = val
			for i := 0; i < val; i++ {
				var tmp []bool
				for j := 0; j < val; j++ {
					tmp = append(tmp, false)
				}
				matrix = append(matrix, tmp)
			}
		}
		func addEdge(s, t int) {
			matrix[s][t] = true
		}
		//--------------------- 有向无权图(邻接表) ---------------------
		import "container/list"

		var v int
		var matrix[]list.List
		func graph(val int) {
			v = val
			for i := 0; i < val; i++ {
				tmp := list.List{}
				matrix = append(matrix, tmp)
			}
		}
		func addEdge(s, t int) {
			matrix[s].PushBack(t)
		}

	3) 图上各种算法
		3.1) 搜索 OR 遍历
			- BFS
			- DFS
		3.2) 最短路径
			- Dijkstra: 针对有权图的单源最短路径算法,并且要求没有负权边
			- Bellman-Ford: 针对有权图的单源最短路径算法,允许存在负权边
			- Floyd: 针对没有权图的多源最短路径算法,允许存在负权边,但不允许负权环
			- A*算法: 启发式搜索算法,求有权图的次优最短路线
		3.3) 最小生成树
			- Prim算法
			- Kruskal算法
		3.4) 最大流,二分分配
			- Ford-Fulkerson
			- Edmonds-Karp

	4) 广度优先搜索/遍历算法
		树是图的一种特殊情况,二又树的按层遍历，实际上就是广度优先搜系，从根节点开始，一层一层的从上往下遍历，
		先遍历与根节点近的，再逐层遍历与根节点远的。
		图上的广度优先搜索（或遍历）跟树上的按层遍历很像，先查找离起始顶点s最近的，然后是次近的，依次往外搜索，
		直到找到终止顶点t（或所有顶点都遍历了一遍）。
		树的按层遍历需要用到队列，同理，图的广度优先搜索（或遍历）也要用到队列。除此之外，对于图的按层遍历，
		需要用一个visited数组，记录已经遍历过的顶点，防止图中存在环，出现循环遍历多次的情况。

		广度优先搜索-遍历算法.png

		广度优先搜索处理是无权图，实际上，通过广度优先搜索找到的源点到终点的路径也是顶点s到顶点t的最短路径。

		var v int
		var matrix[]list.List
		func Graph(val int) {
			v = val
			for i := 0; i < val; i++ {
				tmp := list.List{}
				matrix = append(matrix, tmp)
			}
		}
		func addEdge(a, t int) {
			matrix[a].PushBack(t)
			matrix[t].PushBack(a)
		}

		//----------------------- 广度优先搜索代码实现(查找 s -> t 之间是否存在有效路径) -----------------------
		var v int
		var matrix []list.List
		func graph(val int) {
			v = val
			for i := 0; i < val; i++ {
				tmp := list.List{}
				matrix = append(matrix, tmp)
			}
		}
		func addEdge(s, t int) {
			matrix[s].PushBack(t)
		}

		// 搜索代码
		func bfs(s, t int) {
			visited := make([]bool, v) // 记录是否被方位过, 避免重复访问
			queue := list.List{}       // 用来记录节点, 访问拓展(类似二叉树的遍历)
			queue.PushBack(s)          // 向队列种插入第一个元素
			visited[s] = true
			for queue.Len() > 0 {
				lastVal := queue.Back().Value.(int)

				if lastVal == t { // 找到了结果(加入了此处判断就为搜索, 不加则为遍历)
					return
				}

				tmp := matrix[lastVal].Front()
				for i := 0; i < matrix[lastVal].Len(); i++ {
					if i != 0 {
						tmp = tmp.Next()
					}
					tmpVal := tmp.Value.(int)
					if !visited[tmpVal] {
						visited[tmpVal] = true
						queue.PushBack(tmpVal)
					}
				}
			}
		}

		//----------------------- 广度优先搜索代码实现(查找 s -> t 之间是否存在有效路径, 并打印路径) -----------------------
		var v int
		var matrix []list.List // 无向无权图
		func graph(val int) {
			v = val
			for i := 0; i < val; i++ {
				tmp := list.List{}
				matrix = append(matrix, tmp)
			}
		}
		func addEdge(s, t int) {
			matrix[s].PushBack(t)
		}

		// 搜索代码
		func bfs(s, t int) {
			visited := make([]bool, v) // 记录是否被方位过, 避免重复访问
			queue := list.List{}       // 用来记录节点, 访问拓展(类似二叉树的遍历)
			queue.PushBack(s)          // 向队列种插入第一个元素
			visited[s] = true

			prev := make([]int, v) 	// ##### 新增切片记录路径
			for i := 0; i < v; i++ {
				prev[i] = -1
			}

			for queue.Len() > 0 {
				lastVal := queue.Back().Value.(int)

				if lastVal == t { // 找到了结果(加入了此处判断就为搜索, 不加则为遍历)
					print(prev, s, t) // ###### 打印出路径
					return
				}

				tmp := matrix[lastVal].Front()
				for i := 0; i < matrix[lastVal].Len(); i++ {
					if i != 0 {
						tmp = tmp.Next()
					}
					tmpVal := tmp.Value.(int)
					if !visited[tmpVal] {
						prev[tmpVal] = lastVal // #### 修改状态
						visited[tmpVal] = true
						queue.PushBack(tmpVal)
					}
				}
			}
		}
		func print(prev []int, s, t int) { // 类似一个链表, 倒叙输出
			if prev[t] != -1 && s != t {
				print(prev, s, prev[t]) // 通过递归实现倒叙输出的过程
			}
			fmt.Println(t , " ") // 打印出路径
		}

	5) 深度优先搜索/遍历算法
		广度优先搜索是一种“地毯式”的搜索策略，那么深度优先搜索（DFS）就是一种“不撞南墙不回头”的搜索策略。DFS是图上的回溯。
		沿着一条路一股脑往前走，当走到无路可走时，再回退到上一个岔路口，选择另一条路继续前进。

		树的前中后序遍历就是深度优先遍历。前中后序的区别仅仅在于处理节点的时机不同。
		换句话说：树上的深度优先遍历又分为三类，前/中/后 序遍历

		type Node struct {
			Val int
			Left *Node
			Righ *Node
		}
		func dfs(root *Node) {
			if root == nil {
				return
			}
			fmt.Println(root.Val)
			dfs(root.Left)
			dfs(root.Righ)
		}
		//----------------------- 深度优先搜索, 代码实现 -----------------------
		var found bool
		var visited []bool
		var v int
		var matrix []list.List
		func dfs_simple(s, t int) bool {
			found = false
			visited = make([]bool, v)
			dfs_simple_r(s, t)
			return found
		}
		func dfs_simple_r(s, t int) {
			if found {
				return
			}
			visited[s] = true
			if s == t {
				found = true
				return
			}
			tmp := matrix[s].Front()
			for i := 0; i < matrix[s].Len(); i++ {
				if i != 0 {
					tmp = tmp.Next()
				}
				q := tmp.Value.(int)
				if !visited[q] {
					dfs_simple_r(q, t)
				}
			}
		}
		//----------------------- 深度优先搜索, 代码实现(支持打印路径) -----------------------
		var visited []bool
		var v int
		var matrix []list.List
		var resultPath []int
		func dfs(s, t int) []int {
			visited = make([]bool, v)
			resultPath = []int{}

			path := []int{}
			dfs_r(s, t, path)
			return resultPath
		}
		func dfs_r(s, t int, path []int) {
			if s == t {
				copy(resultPath, path)
				return
			}
			visited[s] = true
			path = append(path, s)
			tmp := matrix[s].Front()
			for i := 0; i < matrix[s].Len(); i++ {
				if i != 0 {
					tmp = tmp.Next()
				}
				q := tmp.Value.(int)
				if !visited[q] {
					dfs_r(q, t, path)
				}
			}
			path = path[:len(path)-1] // 这里是关键
		}

	6) 深度和广度优先搜索
		实际上，DFS也是一种回溯算法，也可以看做多阶段决策模型，用回溯来解决。
			·每个阶段都是基于当前节点移动到下一个节点。
			·可选列表是：相邻并没有被访问过的节点。
			·当前阶段做不同的选择，对应下一个阶段是不同的。
			·回溯的结束条件是：所有节点都已经访问完或找到了终止节点。
			·在回溯的过程中，我们用visited数组，记录已经遍历过的顶点，以免循环重复遍历。

		关系图.png

		// ----------------- 理由回溯解决问题 ---------------------
		var val int
		var visited []bool
		var resultPath []int
		var matrix []list.List
		func dfs2(s, t int) []int {
			visited = make([]bool, val)
			path := []int{}
			path = append(path, s)
			visited[s] = true
			backtrackDFS(s, t, path)
			return resultPath
		}
		func backtrackDFS(s, t int, path []int) {
			if s == t { // 结束条件
				tmp := make([]int, len(path))
				copy(tmp, path)
				resultPath = append(resultPath, tmp...)
				return
			}

			tmp := matrix[s].Front()
			for i := 0; i < matrix[s].Len(); i++ {
				if i != 0 {
					tmp = tmp.Next()
				}
				q := tmp.Value.(int)
				if !visited[q] {
					path = append(path, q)
					visited[q] = true
					backtrackDFS(q, t, path)
					path = path[:len(path)-1]
				}
			}
		}

	拓扑排序算法
		- Kahn 算法
		- DFS 算法

		Kahn 算法(拓扑排序.png):
			定义数据结构：如果s需要先于t执行，那就添加一条s指向t的边。所以，每个顶点的入度表示这个顶点依赖多少个其他顶点。
			如果某个顶点的入度变成了0，就表示这个顶点没有依赖的顶点了，或者说这个顶点依赖的顶点都已执行。

			我们从图中找出一个入度为0的顶点，将其输出到拓扑排序的结果序列中，这里的输出就表示被执行。既然这个顶点已经被执行了，
			那么所有依赖它的顶点的入度都可以减1，反映到图上就是把这个顶点的可达顶点的入度都减1。我们循环执行上面的过程，直到所有的顶点都被输出。
			最后的结果序列就是满足所有局部依赖关系的一个拓扑排序。
			//----------------------- Kahn(code) ----------------------------------
			// numCount 元素个数
			// preArr   待处理数组
			func toSortByKahn(numCount int, preArr [][]int) {
				// 构造成临接矩阵
				adj := make([][]int, numCount)
				for i := 0; i < numCount; i++ {
					adj[i] = make([]int, 0)
				}
				for i := 0; i < len(preArr); i++ {
					adj[preArr[i][1]] = append(adj[preArr[i][1]], preArr[i][0])
				}

				// 统计每个顶点的入度
				inDegree := make([]int, numCount)
				for i := 0; i < numCount; i++ {
					for j := 0; j < len(adj[i]); j++ {
						w := adj[i][j]
						inDegree[w]++
					}
				}

				// 记录入度为 0 的顶点
				zeroSet := make([]int, 0) // 度为 0 的集合
				for i := 0; i < len(inDegree); i++ {
					if inDegree[i] == 0 {
						zeroSet = append(zeroSet, i)
					}
				}

				// 进行拓扑排序的输出等操作
				for len(zeroSet) != 0 {
					i := zeroSet[0]
					zeroSet = zeroSet[1:]
					fmt.Println(i)
					for j := 0; j < len(adj[i]); j++ {
						k := adj[i][j]
						inDegree[k]--
						if inDegree[k] == 0 {
							zeroSet = append(zeroSet, k)
						}
					}
				}
			}
**/