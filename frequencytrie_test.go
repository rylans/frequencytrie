package frequencytrie

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestWordsTreeIdentity(t *testing.T){
  assert.Equal(t, NewPrefixTree(), NewPrefixTree())
}
