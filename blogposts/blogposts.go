package blogposts

import (
	"io/fs"
)

func NewPostsFromFS(fileSystem fs.FS) (posts []Post, err error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	for _, f := range dir {

		post, err := getPost(fileSystem, f.Name())

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, file string) (post Post, err error) {
	postFile, err := fileSystem.Open(file)

	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()
	return newPost(postFile)
}
