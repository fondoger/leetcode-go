package main

/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */

// 字典树
// 注意：字典树的单词结点存一个完成的单词，而非一个flag标志

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	root := buildTrieTree(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs212(board, root, i, j, &res)
		}
	}
	return res
}

func dfs212(board [][]byte, p *TrieNode, i, j int, res *[]string) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if board[i][j] == 0 {
		return
	}
	next := p.Childs[board[i][j]-'a']
	if next == nil {
		return
	}
	if next.Word != "" {
		*res = append(*res, next.Word)
		next.Word = ""
	}

	tmp := board[i][j]
	board[i][j] = 0
	dfs212(board, next, i-1, j, res)
	dfs212(board, next, i+1, j, res)
	dfs212(board, next, i, j-1, res)
	dfs212(board, next, i, j+1, res)
	board[i][j] = tmp
}

func buildTrieTree(words []string) *TrieNode {
	root := NewTrieNode()
	for _, w := range words {
		p := root
		for i := 0; i < len(w); i++ {
			if p.Childs[w[i]-'a'] == nil {
				p.Childs[w[i]-'a'] = NewTrieNode()
			}
			p = p.Childs[w[i]-'a']
		}
		p.Word = w
	}
	return root
}

// TrieNode :
type TrieNode struct {
	Childs []*TrieNode // 固定大小的数组
	Word   string
}

// NewTrieNode :
func NewTrieNode() *TrieNode {
	return &TrieNode{
		Childs: make([]*TrieNode, 26),
		Word:   "",
	}
}
