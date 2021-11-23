package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatetime(t *testing.T) {
	res := 1+1
	assert.Equal(t, res,2)

}

func TestIndex(t *testing.T)  {
	w := TestCase.GetRequest("/",nil)
	assert.Equal(t, w.Code,200)
	assert.Equal(t, w.Body.String(),"这是首页")
}

func TestTimenow(t *testing.T)  {
	w := TestCase.GetRequest("/time/now",nil)
	assert.Equal(t, w.Code,200)
}