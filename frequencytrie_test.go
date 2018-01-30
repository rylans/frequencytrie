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

  assert.Equal(t, []string{}, tree.keys(""), "incorrect key derivation for empty string")
}

func TestCharacterKeyGeneration(t *testing.T){
  tree := ForCharacters()

  assert.Equal(t, []string{"h", "e", "y"}, tree.keys("Hey"), "incorrect key derivation for 'hey'")
}

func TestCharacterKeyGenerationForHanRunes(t *testing.T){
  tree := ForCharacters()

  assert.Equal(t, []string{"한", "글", "ㅁ"}, tree.keys("한글ㅁ"), "incorrect key derivation for '한글ㅁ'")
}

func TestWordKeyGenerationEmptyString(t *testing.T){
  tree := ForWords()

  assert.Equal(t, []string{""}, tree.keys(""), "incorrect key derivation for empty string")
}

func TestWordKeyGeneration(t *testing.T){
  tree := ForWords()

  assert.Equal(t, []string{"why", "hello", "there"}, tree.keys("Why hello there"), "incorrect key derivation for sentence")
}

func TestWordKeyGenerationHanRunes(t *testing.T){
  tree := ForWords()

  assert.Equal(t,
    []string{"잘", "아는", "형들"},
    tree.keys("잘 아는 형들"),
    "incorrect key derivation for Han sentence")
}
