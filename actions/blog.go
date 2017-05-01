package actions

import (
	"fmt"

	"github.com/bketelsen/learnonline/models"
	"github.com/gobuffalo/buffalo"
)

// BlogIndex default implementation.
func BlogIndex(c buffalo.Context) error {
	// courses automatically loaded in the
	// middleware
	c.Set("title", "Brian Blogs")
	cl, err := models.GetPostList()
	if err != nil {
		fmt.Println("Error getting models frm cms")
	}
	c.Set("posts", cl)
	return c.Render(200, r.HTML(sitePath(c, "posts.html")))
}

// BlogShow default implementation.
func BlogShow(c buffalo.Context) error {
	post, err := models.GetPostBySlug(c.Param("slug"))
	if err != nil {
		return err
	}

	c.Set("post", post)
	c.Set("title", post.Title)
	return c.Render(200, r.HTML(sitePath(c, "post.html")))
}
