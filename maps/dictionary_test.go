package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("right key", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is just a test",
		}
		value, err := Search(dictionary, "test")
		assert.Equal(t, err, nil)

		want := "this is just a test"
		assert.Equal(t, want, value)
	})

	t.Run("error key", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is just a test",
		}
		value, err := Search(dictionary, "error_key")
		assert.EqualError(t, err, ErrNotFound.Error())

		want := ""
		assert.Equal(t, want, value)
	})

	t.Run("test create and update", func(t *testing.T) {
		//create
		dictionary := Dictionary{}
		err := dictionary.Create("test", "this is just a test")
		assert.Equal(t, nil, err)

		value, err := Search(dictionary, "test")
		assert.Equal(t, err, nil)
		want := "this is just a test"
		assert.Equal(t, want, value)

		//repeat create
		err = dictionary.Create("test", "this is new value")
		assert.EqualError(t, err, ErrWordExists.Error())

		//update
		err = dictionary.Update("test", "this is updated value")
		value, err = Search(dictionary, "test")
		assert.Equal(t, err, nil)
		want = "this is updated value"
		assert.Equal(t, want, value)

		//error key to update value
		err = dictionary.Update("test2", "this is updated value")
		assert.EqualError(t, err, ErrWordDoesNotExist.Error())



	})

}
