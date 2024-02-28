package blogposts

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (c StubFailingFS) Open(name string) (fs fs.File, error error) {
	return nil, errors.New("oh no, i always fail")
}

func TestExpectError(t *testing.T) {
	_, err := NewPostsFromFS(StubFailingFS{})

	if err == nil {
		t.Errorf("expected an error, but dint get one")
	}
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
		Description: IDK`
		secondBody = `Title: Post 2
		Description: POP`
	)

	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := Post{Title: "Post 1", Description: "IDK"}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
