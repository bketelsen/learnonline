package actions

import (
	"fmt"

	"github.com/bketelsen/learnonline/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
)

// WebinarsIndex default implementation.
func WebinarsIndex(c buffalo.Context) error {
	// webinars automatically loaded in the
	// middleware
	c.Set("title", "Available Classes")
	cl, err := models.GetWebinarList()
	if err != nil {
		fmt.Println("Error getting models frm cms")
	}
	c.Set("webinarlist", cl)
	return c.Render(200, r.HTML(sitePath(c, "webinars.html")))
}

// WebinarsShow default implementation.
func WebinarsShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	webinar, err := models.GetFullWebinarBySlug(c.Param("slug"))
	if err != nil {
		return err
	}

	if c.Value("current_user") != nil {
		webinar.MarkAsPurchased(tx, c.Value("current_user").(*models.User))
	}
	c.Set("webinar", webinar)
	c.Set("title", webinar.Webinar.Title)
	return c.Render(200, r.HTML(sitePath(c, "webinar.html")))
}

