package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	qRef = Queue{
		NewBlock("week 4", NewAlbum("nice"), NewAlbum("nice")),
		NewBlock("week 3", NewAlbum("placeholder")),
		NewBlock("week 2", NewAlbum("fizz"), NewAlbum("buzz"), NewAlbum("fizbuzz")),
		NewBlock("week 1", NewAlbum("foo"), NewAlbum("bar"), NewAlbum("foobar")),
	}
)

func TestAddBlock(t *testing.T) {
	tmp := qRef
	qLocal := &tmp

	b := NewBlock("week 5", NewAlbum("darude - sandstorm"))

	q := &Queue{
		b,
		NewBlock("week 4", NewAlbum("nice"), NewAlbum("nice")),
		NewBlock("week 3", NewAlbum("placeholder")),
		NewBlock("week 2", NewAlbum("fizz"), NewAlbum("buzz"), NewAlbum("fizbuzz")),
	}

	qLocal.Add(b)

	if !cmp.Equal(qLocal, q) {
		t.Error("ms | Error with AddBlock: didn't add properly")

		fmt.Println("Reference:")
		qRef.ShowCurrent()

		fmt.Println("\nNew:")
		q.ShowCurrent()
	}

	q2 := &Queue{}
	q2.Add(b)
	q3 := &Queue{b}

	if !cmp.Equal(q2, q3) {
		t.Error("ms | Error with AddBlock: couldn't add properly to empty queue")
	}
}
