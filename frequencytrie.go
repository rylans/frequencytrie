package frequencytrie 

import (
  "strings"
)

type KeySequenceGenerator func(s string) []string

type TrieNode struct {
  children map[string]TrieNode
  character string
  keygen KeySequenceGenerator
}

func (n *TrieNode) keys(str string) []string {
  return n.keygen(str)
}

func (n *TrieNode) print() {
  n.printWithIndentation(" ")
}

func (n *TrieNode) printWithIndentation(str string) {
  for k := range n.children {
    m := n.children[k]
    m.printWithIndentation(str + " ")
  }

}

func (n *TrieNode) loadWord(word string){
  if len(word) > 0 {
    char := string(word[0])
    rest := word[1:]

    if _, exists := n.children[char]; exists {
      v := n.children[char]
      v.loadWord(rest)
    } else {
      next := NewPrefixTree()
      next.character = char
      n.children[char] = next
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
      candidates = v.suggestWithPrefix("", prefix + v.character, candidates) 
    }
  }

  if len(n.children) == 0 {
    candidates = append(candidates, prefix)
  }

  return candidates

}

func NewPrefixTree() TrieNode {
  m1 := make(map[string]TrieNode)
  return TrieNode{children: m1, character: ""}
}

func ForCharacters() TrieNode {
  f := func(s string) []string {
    return strings.Split(strings.ToLower(s), "")
  }
  m := make(map[string]TrieNode)
  return TrieNode{children: m, character: "", keygen: f}
}

func ForWords() TrieNode {
  f := func(s string) []string {
    return strings.Split(strings.ToLower(s), " ")
  }
  m := make(map[string]TrieNode)
  return TrieNode{children: m, character: "", keygen: f}
}
