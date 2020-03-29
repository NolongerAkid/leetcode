package main
type Trie struct {
	val string
	children []*Trie
	end bool

}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	cur := this
	for i :=0;i<len(word);i++{
		j :=0
		l := len(cur.children)
		for ;j<l;j++{
			if cur.children[j].val == string(word[i]){
				break
			}
		}
		if j == l{
			cur.children = append(cur.children,&Trie{
				val:      string(word[i]),
				children: nil,
				end:      false,
			})
		}
		cur = cur.children[j]
		if i == len(word)-1{
			cur.end = true
		}

	}

}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i :=0;i<len(word);i++{
		j :=0
		l := len(cur.children)
		for ;j<l;j++{
			if cur.children[j].val == string(word[i]){
				cur = cur.children[j]
				break
			}

		}
		if j == l{
			return false
		}
	}
	if cur.end{
		return true
	}else{
		return false
	}

}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this == nil {
		return false
	}
	cur := this
	for i :=0;i<len(prefix);i++{
		j :=0
		l := len(cur.children)
		for ;j<l;j++{
			if cur.children[j].val == string(prefix[i]){
				cur = cur.children[j]
				break
			}

		}
		if j == l{
			return false
		}
	}
	return true

}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	
}
