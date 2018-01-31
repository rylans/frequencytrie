package frequencytrie

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestWordsTreeIdentity(t *testing.T){
  assert.Equal(t, NewPrefixTree(), NewPrefixTree())
}

func TestCharacterKeyGenerationEmptyString(t *testing.T){
  tree := ForCharacters()

  assert.Equal(t, []string{""}, tree.keys(""), "incorrect key derivation for empty string")
}

func TestCharacterKeyGeneration(t *testing.T){
  tree := ForCharacters()

  assert.Equal(t, []string{"h", "e", "y",""}, tree.keys("Hey"), "incorrect key derivation for 'hey'")
}

func TestCharacterKeyGenerationForHanRunes(t *testing.T){
  tree := ForCharacters()

  assert.Equal(t, []string{"한", "글", "ㅁ", ""}, tree.keys("한글ㅁ"), "incorrect key derivation for '한글ㅁ'")
}

func TestWordKeyGenerationEmptyString(t *testing.T){
  tree := ForWords()

  assert.Equal(t, []string{"", ""}, tree.keys(""), "incorrect key derivation for empty string")
}

func TestWordKeyGeneration(t *testing.T){
  tree := ForWords()

  assert.Equal(t, []string{"why", "hello", "there", ""}, tree.keys("Why hello there"), "incorrect key derivation for sentence")
}

func TestWordKeyGenerationHanRunes(t *testing.T){
  tree := ForWords()

  assert.Equal(t,
    []string{"잘", "아는", "형들", ""},
    tree.keys("잘 아는 형들"),
    "incorrect key derivation for Han sentence")
}

func TestCharacterProbability(t *testing.T){
  tree := ForCharacters()
  tree.Insert("foo")
  tree.Insert("bar")
  tree.Insert("bare")
  tree.Insert("bag")
  tree.Insert("bet")

  assert.Equal(t, 1.0, tree.P("foo", "f"))
  assert.Equal(t, 0.20, tree.P("f", ""))
  assert.Equal(t, 0.8, tree.P("b", ""))
  assert.Equal(t, 0.75, tree.P("ba", "b"))
  assert.Equal(t, 0.25, tree.P("be", "b"))
  assert.Equal(t, 1.0, tree.P("bet", "be"))
}

func TestCharacterProbabilityFromEmptyTree(t *testing.T){
  tree := ForCharacters()

  assert.Equal(t, 0.0, tree.P("", ""))
  assert.Equal(t, 0.0, tree.P("abc", ""))
  assert.Equal(t, 0.0, tree.P("ab", "a"))
}

func TestCharacterProbabilityFromSingularTree(t *testing.T){
  tree := ForCharacters()
  tree.Insert("Hello")

  assert.Equal(t, 0.0, tree.P("x", ""))
  assert.Equal(t, 1.0, tree.P("h", ""))
  assert.Equal(t, 1.0, tree.P("he", ""))
  assert.Equal(t, 1.0, tree.P("hell", ""))
  assert.Equal(t, 1.0, tree.P("hell", "h"))

  tree.Insert("helios")
 
  assert.Equal(t, 0.5, tree.P("hell", "hel"))

  assert.Equal(t, 0.5, tree.P("helios", "h"))
  assert.Equal(t, 0.5, tree.P("helios", "he"))
  assert.Equal(t, 0.5, tree.P("helios", "hel"))
  assert.Equal(t, 0.0, tree.P("helios", "hell"))
  //assert.Equal(t, 1.0, tree.P("helios", "heli"))
}
