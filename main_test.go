package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//t 是你在测试框架中的 hook（钩子），所以当你想让测试失败时可以执行 t.Fail() 之类的操作
func TestHello(t *testing.T) {
	//子测试
	//有时，对一个「事情」进行分组测试，然后再对不同场景进行子测试非常有效
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Luis", "")
		want := "hello Luis"
		assert.Equal(t, want, got)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		//测试传空字符串时，默认返回hello world
		got := Hello("", "")
		want := "hello world"
		assert.Equal(t, want, got)
	})

	//西班牙语言时，更换predix
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Luis", spanish)
		want := "hola Luis"
		assert.Equal(t, want, got)

	})

	//法语时
	t.Run("in French", func(t *testing.T) {
		got := Hello("Luis", french)
		want := "bonjour Luis"
		assert.Equal(t, want, got)

	})
}

//通过运行 godoc -http :8000，可以在本地启动文档。如果你访问 localhost:8000/pkg，将看到系统上安装的所有包。
//浏览 http://localhost:8000/pkg/testing/ 是非常值得的，去看一下有什么对你有价值的内容。
