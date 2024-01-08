package solutions

import "fmt"

/* Question **********************************
? 208. Implement Trie (Prefix Tree)
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of t data structure, such as autocomplete and spellchecker.

Implement the Trie class:
Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.


Example 1:
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True


Constraints:
1 <= word.length, prefix.length <= 2000
word and prefix consist only of lowercase English letters.
At most 3 * 104 calls in total will be made to insert, search, and startsWith.
*********************************************/

func Output208() any {
	trie := NewTrie()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
	return fmt.Sprintln("Success")
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func NewTrie() Trie {
	return Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
}

func (t *Trie) Insert(word string) {
	current := t
	for _, c := range word {
		idx := c - 'a'
		if current.children[idx] == nil {
			newNode := NewTrie()
			current.children[idx] = &newNode
		}
		current = current.children[idx]
	}
	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	current := t
	for _, c := range word {
		idx := c - 'a'
		if current.children[idx] == nil {
			return false
		}
		current = current.children[idx]
	}
	return current.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t
	for _, c := range prefix {
		idx := c - 'a'
		if current.children[idx] == nil {
			return false
		}
		current = current.children[idx]
	}
	return true
}
