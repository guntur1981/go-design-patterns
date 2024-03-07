package main

import "fmt"

type Blogger struct {
	Id   int
	Name string
}

type Post struct {
	Id          int
	Title       string
	Description string
	Blogger     Blogger
}

type BloggerRepository interface {
	CreateBlogger(b *Blogger) error
	ReadBlogger(id int) *Blogger
}

type PostRepository interface {
	CreatePost(p *Post) error
	ReadPost(title string) *Post
	UpdatePost(p *Post) error
}

type CacheRepository struct {
	bloggers []Blogger
	posts    []Post
}

func (c *CacheRepository) CreateBlogger(b *Blogger) error {
	b.Id = len(c.bloggers) + 1
	c.bloggers = append(c.bloggers, *b)
	return nil
}
func (c *CacheRepository) ReadBlogger(id int) *Blogger {
	var result *Blogger
	for b := range c.bloggers {
		if c.bloggers[b].Id == id {
			result = &c.bloggers[b]
			break
		}
	}
	return result
}
func (c *CacheRepository) CreatePost(p *Post) error {
	// ...
	return nil
}
func (c *CacheRepository) ReadPost(title string) *Post {
	// ...
	return nil
}
func (c *CacheRepository) UpdatePost(p *Post) error {
	// ...
	return nil
}

func AddBlogger(repo BloggerRepository, b *Blogger) {
	err := repo.CreateBlogger(b)
	if err != nil {
		panic(err)
	}

	fmt.Println("New created blogger id is:", b.Id)
}

func main() {
	cache := CacheRepository{}

	f := Blogger{}
	f.Name = "Foo"
	b := Blogger{}
	b.Name = "Bar"

	AddBlogger(&cache, &f)
	AddBlogger(&cache, &b)
}
