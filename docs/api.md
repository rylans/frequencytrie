

# frequencytrie
`import "github.com/rylans/frequencytrie"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package frequencytrie provides a trie implementation that can be used to calculate the probability of strings from a corpus of text.




## <a name="pkg-index">Index</a>
* [type KeySequenceGenerator](#KeySequenceGenerator)
* [type TransitionChance](#TransitionChance)
  * [func (t TransitionChance) String() string](#TransitionChance.String)
* [type TrieNode](#TrieNode)
  * [func ForCharacters() TrieNode](#ForCharacters)
  * [func ForWords() TrieNode](#ForWords)
  * [func (n *TrieNode) Contains(str string) bool](#TrieNode.Contains)
  * [func (n *TrieNode) FindFirst(str string) (*TrieNode, bool)](#TrieNode.FindFirst)
  * [func (n *TrieNode) Insert(str string)](#TrieNode.Insert)
  * [func (n *TrieNode) Key() string](#TrieNode.Key)
  * [func (n *TrieNode) Len() int](#TrieNode.Len)
  * [func (n *TrieNode) P(str string, given string) float64](#TrieNode.P)
  * [func (n TrieNode) String() string](#TrieNode.String)
  * [func (n *TrieNode) TransitionProbabilities(str string) []TransitionChance](#TrieNode.TransitionProbabilities)


#### <a name="pkg-files">Package files</a>
[frequencytrie.go](/src/github.com/rylans/frequencytrie/frequencytrie.go) 






## <a name="KeySequenceGenerator">type</a> [KeySequenceGenerator](/src/target/frequencytrie.go?s=577:626#L19)
``` go
type KeySequenceGenerator func(s string) []string
```
A KeySequenceGenerator splits the input string into a string slice. The elements of the string slice are to be used as the keys of the trie.










## <a name="TransitionChance">type</a> [TransitionChance](/src/target/frequencytrie.go?s=196:274#L9)
``` go
type TransitionChance struct {
    Probability float64
    // contains filtered or unexported fields
}
```









### <a name="TransitionChance.String">func</a> (TransitionChance) [String](/src/target/frequencytrie.go?s=276:317#L14)
``` go
func (t TransitionChance) String() string
```



## <a name="TrieNode">type</a> [TrieNode](/src/target/frequencytrie.go?s=857:966#L31)
``` go
type TrieNode struct {
    // contains filtered or unexported fields
}
```
A Trie is an N-ary tree. All descendants of a given node have the same prefix







### <a name="ForCharacters">func</a> [ForCharacters](/src/target/frequencytrie.go?s=5240:5269#L213)
``` go
func ForCharacters() TrieNode
```
ForCharacters creates and initializes a new trie with a KeySequenceGenerator that splits the input string into a lowercase sequence of characters.


### <a name="ForWords">func</a> [ForWords](/src/target/frequencytrie.go?s=5609:5633#L222)
``` go
func ForWords() TrieNode
```
ForWords creates and initializes a new trie with a KeySequenceGenerator that splits the input string into a lowercase sequence of words.





### <a name="TrieNode.Contains">func</a> (\*TrieNode) [Contains](/src/target/frequencytrie.go?s=3700:3744#L152)
``` go
func (n *TrieNode) Contains(str string) bool
```



### <a name="TrieNode.FindFirst">func</a> (\*TrieNode) [FindFirst](/src/target/frequencytrie.go?s=3285:3343#L130)
``` go
func (n *TrieNode) FindFirst(str string) (*TrieNode, bool)
```



### <a name="TrieNode.Insert">func</a> (\*TrieNode) [Insert](/src/target/frequencytrie.go?s=4341:4378#L180)
``` go
func (n *TrieNode) Insert(str string)
```
Insert a string value into the tree




### <a name="TrieNode.Key">func</a> (\*TrieNode) [Key](/src/target/frequencytrie.go?s=1275:1306#L50)
``` go
func (n *TrieNode) Key() string
```



### <a name="TrieNode.Len">func</a> (\*TrieNode) [Len](/src/target/frequencytrie.go?s=4085:4113#L168)
``` go
func (n *TrieNode) Len() int
```
Len returns the number of items inserted into the tree




### <a name="TrieNode.P">func</a> (\*TrieNode) [P](/src/target/frequencytrie.go?s=2157:2211#L85)
``` go
func (n *TrieNode) P(str string, given string) float64
```



### <a name="TrieNode.String">func</a> (TrieNode) [String](/src/target/frequencytrie.go?s=968:1001#L37)
``` go
func (n TrieNode) String() string
```



### <a name="TrieNode.TransitionProbabilities">func</a> (\*TrieNode) [TransitionProbabilities](/src/target/frequencytrie.go?s=1337:1410#L54)
``` go
func (n *TrieNode) TransitionProbabilities(str string) []TransitionChance
```


