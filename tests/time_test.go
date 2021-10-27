package tests

import (
	"testing"
)

func TestDatetime(t *testing.T) {
	res := 1+1
	t.Log(res)
}

func TestIndex(t *testing.T)  {
	w := TestCase.GetRequest("/",nil)
	t.Log(w.Body.String())
}

func TestTimenow(t *testing.T)  {
	w := TestCase.GetRequest("/time/now",nil)
	t.Log(w.Body.String())
}