package mail


import (
	"testing"
)


func TestDbSetup(t *testing.T) {

	t.Run("Testing if mail works", func(t *testing.T) {

		err := exp()
		if err != nil {
			t.Fatal(err)
		}

	})

}
