package trieDeStrings

import (
	"fmt"
	//"reflect"
	//"strings"
	"os"
	"bufio"
	"regexp/syntax"
	//"unsafe"
)


type ZNode struct {
	runeVal		rune
	isEntry		bool
	parent		*ZNode
	children	map[rune]*ZNode
	depth		uint8
}

type ZTrie struct {
	rootNode	*ZNode
	titleCount	uint
	nodeCount	uint
}



func NewTrie() *ZTrie{
	return &ZTrie{
		rootNode: &ZNode{children: make(map[rune]*ZNode), depth: 0},
		titleCount: 0,
		nodeCount:	0,
	}
}



// add string to trie
func (tr *ZTrie) AddString(newString string) *ZNode {
	tr.titleCount++
	nodeWalker := tr.rootNode
	runes := []rune(newString)

	for i := range runes {
		aRune := runes[i]
		if cNode, exists := nodeWalker.children[aRune]; exists {
			nodeWalker = cNode
		} else {
			tr.nodeCount++
			nodeWalker = nodeWalker.NewChild(aRune, false)
		}
	}
	tr.nodeCount++
	nodeWalker = nodeWalker.NewChild(0x0, true)
	return nodeWalker
}



// add new child node
func (pNode *ZNode) NewChild(aRune rune, titleBoo bool) *ZNode {
	newNode := &ZNode{
		runeVal:	aRune,
		isEntry:	titleBoo,
		parent:		pNode,
		children:	make(map[rune]*ZNode),
		depth:		pNode.depth + 1,
	}
	pNode.children[aRune] = newNode
	return newNode
}



// Returns root node
func (tr *ZTrie) GetCount() uint {
	return tr.titleCount
}
// Returns root node
func (tr *ZTrie) getRoot() *ZNode {
	return tr.rootNode
}

// Returns parent node
func (n ZNode) getParent() *ZNode {
	return n.parent
}
// Returns Children nodes
func (n ZNode) getChildren() map[rune]*ZNode {
	return n.children
}
// Return Rune Value
func (n ZNode) getRuneVal() rune {
	return n.runeVal
}




// search
func (tr ZTrie) PrefixSearch(qString string, rCount uint64) *[]string {
	node := findNode(tr.rootNode, []rune(qString))
	if node == nil {
		return &[]string{}
		//no match
	}
	return getMatches(node, rCount)
}

// returns end-node of most specific path
// a node to examine & next set of runes to check
func findNode(currentNode *ZNode, qRunes []rune) *ZNode {


	if currentNode == nil {
		return nil
		//dead end
	}
	if len(qRunes) == 0 {

		return currentNode
		//end of search string
	}

	cNode, exists := currentNode.getChildren()[qRunes[0]]
	if !exists {
		return nil
		//no match in any children
	}

	var newRuneSet []rune
	if len(qRunes) > 1 {
		newRuneSet = qRunes[1:]
		//progress forward 1 character
	} else {
		newRuneSet = qRunes[0:0]
		//trigger match with empty slice
	}

	return findNode(cNode, newRuneSet)
}

// 
func getMatches(subRtNode *ZNode, rCount uint64) *[]string {

	nodes := []*ZNode{subRtNode}
	var matches []string
	var idx int
	var nodeWalker *ZNode

	for cnt := len(nodes); cnt != 0 && uint64(len(matches))<rCount; cnt = len(nodes) {
		idx = cnt - 1
		nodeWalker = nodes[idx]
		nodes = nodes[:idx]
		for _, aChildNode := range nodeWalker.children {
			nodes = append(nodes, aChildNode)
		}
		if nodeWalker.isEntry {
			aTitle := ""
			for p := nodeWalker.parent; p.depth != 0; p = p.parent {
				aTitle = string(p.runeVal) + aTitle
			}
			matches = append(matches, aTitle)
		}
	}
	return &matches
}










func (tr *ZTrie) LoadTrie(filePath *os.File) {


	r := bufio.NewScanner(filePath)
	for r.Scan() {
		line := r.Text()

		if line != "" {
			atx := []rune(line)[0]
			if syntax.IsWordChar(atx) && len(line)<8 {  //66=B  98=b
				fmt.Println(line)
				tr.AddString(line)
			}
		}

	}


	if err := r.Err(); err != nil {
		fmt.Println("Error Here: ")
		fmt.Println(err)
	}
	fmt.Println("Done Making Trie Structure")
	fmt.Print(tr.titleCount," strings added ")
	fmt.Printf("using %d nodes\n", tr.nodeCount )
	//fmt.Println(  *tr.rootNode  )
	//fmt.Println(  *((*tr.rootNode).children[90])  )
	//fmt.Println(  (*((*tr.rootNode).children[90]).children[101])  )
	//fmt.Println(  unsafe.Sizeof( *(*((*tr.rootNode).children[90]).children[101]).children[0]    ))


}

