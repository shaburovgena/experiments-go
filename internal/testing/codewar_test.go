package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestCamelString(t *testing.T) {
	s := "moskow_Piter_eburg"
	assert.Equal(t, returnCamel(s), "moskowPiterEburg")
}

func returnCamel(s string) string {
	return fmt.Sprint(regexp.MustCompile("[_-](.)").ReplaceAllStringFunc(s, func(w string) string {
		return strings.ToUpper(w[1:])
	}))
}

func TestTriangle(t *testing.T) {

	assert.Equal(t, false, IsTriangle(5, 1, 2))
	assert.Equal(t, false, IsTriangle(1, 2, 5))
	assert.Equal(t, false, IsTriangle(2, 5, 1))
	assert.Equal(t, true, IsTriangle(4, 2, 3))
	assert.Equal(t, true, IsTriangle(5, 1, 5))
	assert.Equal(t, true, IsTriangle(2, 2, 2))
	assert.Equal(t, false, IsTriangle(-1, 2, 3))
}

func IsTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && a+c > b
}
func TestCamelString1(t *testing.T) {
	s := "moskow piter eburg"
	assert.Equal(t, "MoskowPiterEburg", returnCamelSpace(s))
	assert.Equal(t, "TestCase", returnCamelSpace("test case"))
	assert.Equal(t, "SayHello", returnCamelSpace("say hello "))
	assert.Equal(t, "CamelCaseWord", returnCamelSpace(" camel case word"))
	assert.Equal(t, "", returnCamelSpace(""))
}

func returnCamelSpace(s string) string {

}
