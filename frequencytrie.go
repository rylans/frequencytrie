package frequencytrie 

import (
  "strings"
  "strconv"
  "fmt"
)

type KeySequenceGenerator func(s string) []string

type CountedKey struct {
  key string
  count int
}

func (k *CountedKey) String() string {
  return k.key + "  " + strconv.Itoa(k.count)
}

type TrieNode struct {
  children map[string]TrieNode
  character *CountedKey
  keygen KeySequenceGenerator
}

func (n TrieNode) String() string {
  return "TrieNode{" + n.character.String() + "}"
}

func (n *TrieNode) keys(str string) []string {
  return append(n.keygen(str), "")
}

func (n *TrieNode) print() {
  n.printWithIndentation(" ")
}

func (n *TrieNode) printWithIndentation(str string) {
  for k := range n.children {
    m := n.children[k]
    fmt.Println(str, n.character)
    m.printWithIndentation(str + " ")
  }

}

func (n *TrieNode) P(str string, given string) float64 {
  return n.probability(n.keys(str), n.keys(given), n.character.count)
}

func (n *TrieNode) probability(sequence []string, givenSequence []string, parentCount int) float64 {
  if len(sequence) > 0 && len(givenSequence) > 0 {
    head, tails := sequence[0], sequence[1:]
    givenHead, givenTails := givenSequence[0], givenSequence[1:]

    thisCount := n.character.count
    if head == givenHead {
      if child, exists := n.children[head]; exists {
	return child.probability(tails, givenTails, thisCount)
      } else {
	return 0.0
      }
    } else if givenHead != "" {
      return 0.0
    } else {
      // compute properly when sequence has more characters than one here
      queryCount := 0
      if child, exists := n.children[head]; exists {
	queryCount = child.character.count
	if thisCount > 0 {
	  p := float64(queryCount) / float64(thisCount)
	  return child.internalP(tails, p)
	} 
      }
    }
  }
  return 0.0
}

func (n *TrieNode) internalP(seq []string, prob float64) float64 {
  if len(seq) == 0 {
    return prob
  }
  head, tail := seq[0], seq[1:]
  cnt := n.character.count

  if m, exists := n.nextFor(head); exists {
    subcnt := m.character.count
    if subcnt == 0 {
      return prob
    }
    thisP := float64(subcnt)/float64(cnt)
    return m.internalP(tail, prob * thisP)
  }
  return prob

}

func (n *TrieNode) nextFor(key string) (*TrieNode, bool) {
  if next, exists := n.children[key]; exists {
    return &next, true
  }
  return nil, false
}

func (n *TrieNode) Insert(str string) {
  keySequence := n.keys(str)
  n.loadWord(keySequence)
}

func (n *TrieNode) loadWord(keySequence []string){
  if len(keySequence) > 0 {
    head := keySequence[0]
    rest := keySequence[1:]
    n.character.count = n.character.count + 1

    if v, exists := n.children[head]; exists {
      v.loadWord(rest)
    } else {
      next := NewPrefixTree()
      next.keygen = n.keygen
      next.character = newEmptyCountedKey()
      next.character.key = head
      n.children[head] = next
      next.loadWord(rest)
    }
  }
}

func (n *TrieNode) Suggest(wordPrefix string) []string {
  return n.suggestWithPrefix(wordPrefix, "", make([]string, 0))
}

func (n* TrieNode) suggestWithPrefix(remainingChars string, prefix string, candidates []string) []string {
  if len(remainingChars) > 0 {
    branchingChar := string(remainingChars[0])
    rest := remainingChars[1:]

    cnode := n.children[branchingChar]
    return cnode.suggestWithPrefix(rest, prefix + branchingChar, candidates)
  }

  for k := range n.children {
    if v, exists := n.children[k]; exists {
      candidates = v.suggestWithPrefix("", prefix + v.character.key, candidates) 
    }
  }

  if len(n.children) == 0 {
    candidates = append(candidates, prefix)
  }

  return candidates

}

func newEmptyCountedKey() *CountedKey {
  return &CountedKey{key: "", count: 0}
}

func NewPrefixTree() TrieNode {
  m1 := make(map[string]TrieNode)
  return TrieNode{children: m1, character: newEmptyCountedKey()}
}

func ForCharacters() TrieNode {
  f := func(s string) []string {
    return strings.Split(strings.ToLower(s), "")
  }
  m := make(map[string]TrieNode)
  return TrieNode{children: m, character: newEmptyCountedKey(), keygen: f}
}

func ForWords() TrieNode {
  f := func(s string) []string {
    return strings.Split(strings.ToLower(s), " ")
  }
  m := make(map[string]TrieNode)
  return TrieNode{children: m, character: newEmptyCountedKey(), keygen: f}
}

