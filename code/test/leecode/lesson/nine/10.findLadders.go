package nine

/*
https://leetcode-cn.com/problems/word-transformer-lcci/

单词转换

给定字典中的两个词，长度相等。写一个方法，把一个词转换成另一个词， 但是一次只能改变一个字符。每一步得到的新词都必须能在字典中找到。
编写一个程序，返回一个可能的转换序列。如有多个可能的转换序列，你可以返回任何一个。

示例 1:
	输入:
		beginWord = "hit",
		endWord = "cog",
		wordList = ["hot","dot","dog","lot","log","cog"]
	输出:
		["hit","hot","dot","lot","log","cog"]

示例 2:
	输入:
		beginWord = "hit"
		endWord = "cog"
		wordList = ["hot","dot","dog","lot","log"]
	输出: []
	解释:endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
*/
var visitedFindLadders map[string]bool
var resultPath []string
var foundTag bool
func findLadders(beginWord string, endWord string, wordList []string) []string {
	/*
		就是在 wordList 种利用已有的单词进行排列组合
		这里可以使用 dfs 算法

		回溯的过程要 remove 掉已经已经加入的节点
	*/
	visitedFindLadders = make(map[string]bool, 0)
	resultPath = make([]string, 0)
	foundTag = false
	dfsFindLadders(beginWord, endWord, []string{}, wordList)
	return resultPath
}
func dfsFindLadders(curWord, endWord string, path, wordList []string) {
	if foundTag {return}
	path = append(path, curWord)
	visitedFindLadders[curWord] = true // out of memory
	if curWord == endWord {
		resultPath = append(resultPath, path...)
		foundTag =  true
		return
	}
	for i := 0; i < len(wordList); i++ {
		nextWord := wordList[i]
		if !visitedFindLadders[nextWord] && isValid(curWord, nextWord) {
			dfsFindLadders(nextWord, endWord, path, wordList)
		}
	}
	path = path[:len(path)-1]
}
func isValid(word1, word2 string) bool {
	diff := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diff++
		}
	}
	return diff == 1
}