package leetcode

type WordDictionary struct {
	children [26]*WordDictionary
	isWord   bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (d *WordDictionary) AddWord(word string) {
	curr := d
	for i := range word {
		if curr.children[word[i]-'a'] == nil {
			curr.children[word[i]-'a'] = &WordDictionary{}
		}
		curr = curr.children[word[i]-'a']
	}
	curr.isWord = true
}

func (d *WordDictionary) Search(word string) bool {
	if word == "" {
		return d.isWord
	}

	if word[0] == '.' {
		for i := 0; i < 26; i++ {
			if d.children[i] != nil {
				if d.children[i].Search(word[1:]) {
					return true
				}
			}
		}
		return false
	}
	find := word[0] - 'a'

	if d.children[find] == nil {
		return false
	}
	return d.children[find].Search(word[1:])
}
