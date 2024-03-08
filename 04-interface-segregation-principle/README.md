# Description

The Interface Segregation Principle in Go states that an interface type should not impose any behavior that is not required.

## Example

Consider the following example:

```
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

type Repository interface {
	CreateBlogger(b *Blogger) error
	ReadBlogger(id int) *Blogger

	CreatePost(p *Post) error
	ReadPost(title string) *Post
	UpdatePost(p *Post) error
}

type Cache struct {
	bloggers []Blogger
	posts    []Post
}

func (c *Cache) CreateBlogger(b *Blogger) error {
	b.Id = len(c.bloggers) + 1
	c.bloggers = append(c.bloggers, *b)
	return nil
}
func (c *Cache) ReadBlogger(id int) *Blogger {
	var result *Blogger
	for b := range c.bloggers {
		if c.bloggers[b].Id == id {
			result = &c.bloggers[b]
			break
		}
	}
	return result
}
func (c *Cache) CreatePost(p *Post) error {
	// ...
	return nil
}
func (c *Cache) ReadPost(title string) *Post {
	// ...
	return nil
}
func (c *Cache) UpdatePost(p *Post) error {
	// ...
	return nil
}

func AddBlogger(repo Repository, b *Blogger) {
	err := repo.CreateBlogger(b)
	if err != nil {
		panic(err)
	}

	fmt.Println("New created blogger id is:", b.Id)
}

func main() {
	cache := Cache{}

	f := Blogger{}
	f.Name = "Foo"
	b := Blogger{}
	b.Name = "Bar"

	AddBlogger(&cache, &f)
	AddBlogger(&cache, &b)
}
```

In the above example, the AddBlogger function uses Repository interface to store a blogger.

## The Reason Why This Is Not a Good Practice

For small projects, this may not be an issue. But later, as the Repository interface grew by adding more behaviors, maintaining it become difficult.

## A Better Approach

The AddBlogger function only relates to blogger, not the post. So, it's better to separate this two into different interfaces.

Better code would be, as follows:

```
package main

...

type BloggerRepository interface {
	CreateBlogger(b *Blogger) error
	ReadBlogger(id int) *Blogger
}

type PostRepository interface {
	CreatePost(p *Post) error
	ReadPost(title string) *Post
	UpdatePost(p *Post) error
}

...

func AddBlogger(repo BloggerRepository, b *Blogger) {
	err := repo.CreateBlogger(b)
	if err != nil {
		panic(err)
	}

	fmt.Println("New created blogger id is:", b.Id)
}
```
