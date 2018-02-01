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

  assert.Equal(t, 1.0, tree.P("", ""))
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
  assert.Equal(t, 1.0, tree.P("helios", "heli"))
}

func TestCharacterProbabilityForSeveralWords(t *testing.T){
  tree := ForCharacters()
  tree.Insert("aardvark")
  tree.Insert("buttercup")
  tree.Insert("chai")
  tree.Insert("doodle")
  tree.Insert("doorway")

  assert.Equal(t, 0.2, tree.P("chai", ""))
  assert.Equal(t, 1.0, tree.P("chai", "c"))
  assert.Equal(t, 1.0, tree.P("chai", "cha"))
  assert.Equal(t, 1.0, tree.P("chai", "chai"))

  assert.Equal(t, 0.2, tree.P("doodle", ""))
  assert.Equal(t, 0.5, tree.P("doodle", "d"))
  assert.Equal(t, 0.5, tree.P("doodle", "doo"))
  assert.Equal(t, 0.5, tree.P("doodle", "doo"))
  assert.Equal(t, 1.0, tree.P("doodle", "dood"))

  assert.Equal(t, 0.0, tree.P("aabc", "aabc"))
  assert.Equal(t, 1.0, tree.P("aar", "aa"))
}


func TestCharacterTreeContains(t *testing.T){
  tree := ForCharacters()

  assert.Equal(t, true, tree.Contains(""))
  assert.Equal(t, false, tree.Contains("a"))

  tree.Insert("anagrams")

  assert.Equal(t, true, tree.Contains(""))
  assert.Equal(t, true, tree.Contains("a"))
  assert.Equal(t, true, tree.Contains("anagrams"))
  assert.Equal(t, false, tree.Contains("anagrams "))
  assert.Equal(t, false, tree.Contains("nagrams"))
}

func TestCharacterTreeContainsWithSpaces(t *testing.T){
  tree := ForCharacters()

  tree.Insert("a word or two")

  assert.Equal(t, true, tree.Contains(""))
  assert.Equal(t, true, tree.Contains("a word or two"))
  assert.Equal(t, false, tree.Contains("a word or 2"))
}

func TestWordTreeContainsWithSpaces(t *testing.T){
  tree := ForWords()

  tree.Insert("a word or two")

  assert.Equal(t, true, tree.Contains(""))
  assert.Equal(t, true, tree.Contains("a word or two"))
  assert.Equal(t, false, tree.Contains("a word or 2"))
}

func TestTreeLength(t *testing.T){
  tree := ForCharacters()

  assert.Equal(t, 0, tree.Len())

  tree.Insert("foo")

  assert.Equal(t, 1, tree.Len())

  tree.Insert("bar")
  tree.Insert("foo")

  assert.Equal(t, 3, tree.Len())
}

func TestCharacterTreeFindFirst(t *testing.T){
  tree := ForCharacters()

  _, found := tree.FindFirst("mal")
  assert.Equal(t, false, found)

  tree.Insert("normal")

  n, found := tree.FindFirst("mal")
  assert.Equal(t, true, found)
  assert.Equal(t, 1, n.Len())
  assert.Equal(t, "r", n.Key())

  tree.Insert("mal")

  n, found = tree.FindFirst("mal")
  assert.Equal(t, true, found)
  assert.Equal(t, 2, n.Len())
  assert.Equal(t, "", n.Key())

  n, found = tree.FindFirst("ma")
  assert.Equal(t, true, found)
  assert.Equal(t, "", n.Key())
}

func TestCharacterTransitionProbabilities(t *testing.T){
  tree := ForCharacters()

  tree.Insert("apple")
  tree.Insert("avocado")
  tree.Insert("banana")
  tree.Insert("bandana")

  appleTransitions := tree.TransitionProbabilities("apple")
  assert.Equal(t, 0.5, appleTransitions[0].probability) // "" to "a"
  assert.Equal(t, 0.5, appleTransitions[1].probability) // "a" to "p"
  assert.Equal(t, 1.0, appleTransitions[2].probability) // "p" to "p"
  assert.Equal(t, 1.0, appleTransitions[3].probability) // "p" to "l"
  assert.Equal(t, 1.0, appleTransitions[4].probability) // "l" to "e"
  assert.Equal(t, 1.0, appleTransitions[5].probability) // "e" to ""
  assert.Equal(t, 6, len(appleTransitions))

  bandanaTransitions := tree.TransitionProbabilities("bandana")
  assert.Equal(t, 0.5, bandanaTransitions[0].probability) // "" to "b"
  assert.Equal(t, 1.0, bandanaTransitions[1].probability) // "b" to "a"
  assert.Equal(t, 1.0, bandanaTransitions[2].probability) // "a" to "n"
  assert.Equal(t, 0.5, bandanaTransitions[3].probability) // "n" to "d"
  assert.Equal(t, 8, len(bandanaTransitions))
}

func TestCharacterTransitionProbabilitiesNoSkip(t *testing.T){
  tree := ForCharacters()

  tree.Insert("axial")

  axvialTransitions := tree.TransitionProbabilities("axvial")
  assert.Equal(t, 1.0, axvialTransitions[0].probability)
  assert.Equal(t, 1.0, axvialTransitions[1].probability)
  assert.Equal(t, 0.0, axvialTransitions[2].probability)
  assert.Equal(t, 3, len(axvialTransitions))
}
