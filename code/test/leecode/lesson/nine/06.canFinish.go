package nine

/*
https://leetcode-cn.com/problems/course-schedule/

课程表

你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，
表示如果要学习课程ai 则 必须 先学习课程 bi 。
例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
	输入：numCourses = 2, prerequisites = [[1,0]]
	输出：true
	解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
	输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
	输出：false
	解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：
	1 <= numCourses <= 105
	0 <= prerequisites.length <= 5000
	prerequisites[i].length == 2
	0 <= ai, bi < numCourses
	prerequisites[i] 中的所有课程对 互不相同
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	/*
		题目的实质就是在检测环(拓扑排序,检测环)
		可以采用 kahn 算法 或者 dfs 来解决
	*/
	adjs := make([][]int, numCourses) // 构造成临接矩阵
	for i := 0; i < numCourses; i++ {
		adjs[i] = make([]int, 0)
	}
	indegrees := make([]int, numCourses) // 统计每个课程出现的次数
	for i := 0; i < len(prerequisites); i++ {
		adjs[prerequisites[i][1]] = append(adjs[prerequisites[i][1]], prerequisites[i][0])
		indegrees[prerequisites[i][0]]++
	}

	zeroInDegrees := make([]int, 0) // 度为 0 的集合
	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			zeroInDegrees = append(zeroInDegrees, i)
		}
	}
	zeroInDegreesCount := 0
	for len(zeroInDegrees) > 0 {
		coursei := zeroInDegrees[0]
		zeroInDegrees = zeroInDegrees[1:]
		zeroInDegreesCount++
		for _, coursej := range adjs[coursei] {
			indegrees[coursej]--
			if indegrees[coursej] == 0 {
				zeroInDegrees = append(zeroInDegrees, coursej)
			}
		}
	}
	return zeroInDegreesCount == numCourses
}
