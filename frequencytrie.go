// Package frequencytrie provides a trie implementation that can be used to calculate the probability of strings from a corpus of text.
package frequencytrie 

import (
  "strings"
  "strconv"
)

type TransitionChance struct {
  fromKey, toKey string
  Probability float64
}

func (t TransitionChance) String() string {
  return "{'" + t.fromKey + "' -> '" + t.toKey + "' " + strconv.FormatFloat(t.Probability, 'f', -1, 64) + "}"
}

// A KeySequenceGenerator splits the input string into a string slice. The elements of the string slice are to be used as the keys of the trie.
type KeySequenceGenerator func(s string) []string

type countedKey struct {
  key string
  count int
}

func (k *countedKey) String() string {
  return k.key + "  " + strconv.Itoa(k.count)
}

type TrieNode struct {
  children map[string]TrieNode
  character *countedKey
  keygen KeySequenceGenerator
}

func (n TrieNode) String() string {
  return "TrieNode{" + n.character.String() + "}"
}

func (n *TrieNode) keys(str string) []string {
  return append(n.keygen(str), "")
}

func (n *TrieNode) Key() string {
  return n.character.key
}

func (n *TrieNode) TransitionProbabilities(str string) []TransitionChance {
  transitions := make([]TransitionChance, 0)

  upperNode := n
  var lowerNode *TrieNode
  for _, k := range n.keys(str) {
    if next, exists := upperNode.nextFor(k); exists {
      lowerNode = next
      p := float64(lowerNode.character.count) / float64(upperNode.character.count)
      if lowerNode.character.count == 0 {
	p = 1
      }

      fromkey := upperNode.character.key
      tokey := lowerNode.character.key
      transitions = append(transitions, TransitionChance{
	fromKey: fromkey,
	toKey: tokey,
	Probability: p});
      upperNode = lowerNode
    } else {
      transitions = append(transitions, TransitionChance{
	fromKey: upperNode.character.key,
	toKey: k,
	Probability: 0.0});
      break
    }
  }
  return transitions
}

func (n *TrieNode) P(str string, given string) float64 {
  return n.probability(n.keys(str), n.keys(given))
}

func (n *TrieNode) probability(sequence []string, givenSequence []string) float64 {
  head, tails := sequence[0], sequence[1:]
  givenHead, givenTails := givenSequence[0], givenSequence[1:]

  if head == "" && givenHead == "" {
    return 1.0
  }
  if head == givenHead {
    if child, exists := n.nextFor(head); exists {
      return child.probability(tails, givenTails)
    } 
  } else if givenHead != "" {
    return 0.0
  } else {
    if child, exists := n.nextFor(head); exists {
      p := float64(child.character.count) / float64(n.character.count)
      return child.internalP(tails, p)
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

func (n *TrieNode) FindFirst(str string) (*TrieNode, bool) {
  return n.find(n.keys(str))
}

func (n *TrieNode) find(keys []string) (*TrieNode, bool) {
  if keys[0] == "" {
    return nil, false
  }
  if n.containsKeySequence(keys) {
    return n, true
  }

  for k := range n.children {
    m := n.children[k]
    if node, exists := m.find(keys); exists {
      return node, true
    }
  }
  return nil, false

}

func (n *TrieNode) Contains(str string) bool {
  return n.containsKeySequence(n.keys(str))
}

func (n *TrieNode) containsKeySequence(keys []string) bool {
  if keys[0] == "" {
    return true
  }
  if next, exists := n.nextFor(keys[0]); exists {
    return next.containsKeySequence(keys[1:])
  } else {
    return false
  }
}

func (n *TrieNode) Len() int {
  return n.character.count
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
    n.character.count++

    if v, exists := n.nextFor(head); exists {
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

func newEmptyCountedKey() *countedKey {
  return &countedKey{key: "", count: 0}
}

func NewPrefixTree() TrieNode {
  return TrieNode{children: make(map[string]TrieNode), character: newEmptyCountedKey()}
}

// ForCharacters creates and initializes a new trie with a KeySequenceGenerator that splits the input string into a lowercase sequence of characters.
func ForCharacters() TrieNode {
  f := func(s string) []string {
    return strings.Split(strings.ToLower(s), "")
  }
  m := make(map[string]TrieNode)
  return TrieNode{children: m, character: newEmptyCountedKey(), keygen: f}
}

// ForWords creates and initializes a new trie with a KeySequenceGenerator that splits the input string into a lowercase sequence of words.
func ForWords() TrieNode {
  f := func(s string) []string {
    return strings.Split(strings.ToLower(s), " ")
  }
  m := make(map[string]TrieNode)
  return TrieNode{children: m, character: newEmptyCountedKey(), keygen: f}
}
