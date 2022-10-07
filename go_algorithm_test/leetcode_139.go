package main

type trie struct {
	isNode bool
	val    byte
	list   map[byte]*trie
}

func wordBreak(s string, wordDict []string) bool {
	//制作字典树
	root := &trie{false, ' ', map[byte]*trie{}}
	for _, word := range wordDict {
		p := root
		//先处理头部，使其尽量保持一致
		for i := 0; i < len(word); i++ {
			if node, ok := p.list[word[i]]; ok {
				p = node
			} else {
				p.list[word[i]] = &trie{false, word[i], map[byte]*trie{}}
				p = p.list[word[i]]
			}
			if i == len(word)-1 {
				p.isNode = true
			}

		}
	}
	//搜索函数
	var search func(s string) bool
	search = func(s string) bool {
		if s == "" {
			return true
		}
		location := 0
		for ; location < len(s); location++ {
			p := root
			update := false
			for j := location; j < len(s); j++ {
				if node, ok := p.list[s[j]]; ok {
					p = node
					if p.isNode == true {
						if search(s[j+1:]) {
							return true
						}
						location = j
						update = true
					}
				} else {
					break
				}
			}
			if update == false {
				return false
			}
		}
		if location == len(s) {
			return true
		} else {
			return false
		}
	}
	/*


		location := 0
		for ; location < len(s); location++ {
			p := root
			update := false
			for j := location; j < len(s); j++ {
				if node, ok := p.list[s[j]]; ok {
					p = node
					if p.isNode == true {
						location = j
						update = true
					}
				} else {
					break
				}
			}
			if update == false {
				return false
			}
		}
		if location == len(s) {
			return true
		} else {
			return false
		}
	*/
	return search(s)
}
