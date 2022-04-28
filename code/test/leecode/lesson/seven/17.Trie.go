package seven
/*
https://leetcode-cn.com/problems/implement-trie-prefix-tree/

实现 Trie (前缀树)

Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：
	Trie() 初始化前缀树对象。
	void insert(String word) 向前缀树中插入字符串 word 。
	boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
	boolean startsWith(String prefix) 如果之前已经插入的字符串word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

示例：
	输入
	["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
	[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
	输出
	[null, null, true, false, true, null, true]

解释
	Trie trie = new Trie();
	trie.insert("apple");
	trie.search("apple");   // 返回 True
	trie.search("app");     // 返回 False
	trie.startsWith("app"); // 返回 True
	trie.insert("app");
	trie.search("app");     // 返回 True

提示：
	1 <= word.length, prefix.length <= 2000
	word 和 prefix 仅由小写英文字母组成
	insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
*/
type TrieNode struct {
	data byte
	isEnding bool
	children [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func ConstructorTrie() Trie {
	return Trie{
		root: &TrieNode{
			data: '/',
			children: [26]*TrieNode{},
		},
	}
}

func (this *Trie) Insert(word string)  {
	p := this.root
	for i := 0; i < len(word); i++ {
		tmp := word[i]
		if p.children[tmp-'a'] == nil {
			p.children[tmp-'a'] = &TrieNode{data: tmp}
		}
		p = p.children[tmp-'a']
	}
	p.isEnding = true
}

func (this *Trie) Search(word string) bool {
	p := this.root
	for i := 0; i < len(word); i++ {
		tmp := word[i]
		if p.children[tmp-'a'] == nil {
			return false
		}
		p = p.children[tmp-'a']
	}
	return p.isEnding
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for i := 0; i < len(prefix); i++ {
		tmp := prefix[i]
		if p.children[tmp-'a'] == nil {
			return false
		}
		p = p.children[tmp-'a']
	}
	return true
}