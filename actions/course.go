package actions

import (
	"errors"
	"fmt"

	"github.com/bketelsen/learnonline/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
)

// CoursesIndex default implementation.
func CoursesIndex(c buffalo.Context) error {
	// courses automatically loaded in the
	// middleware
	c.Set("title", "Available Classes")
	cl, err := models.GetCourseList()
	if err != nil {
		fmt.Println("Error getting models frm cms")
	}
	c.Set("courselist", cl)
	return c.Render(200, r.HTML(sitePath(c, "courses.html")))
}

// CoursesShow default implementation.
func CoursesShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	course, err := models.GetFullCourseBySlug(c.Param("slug"))
	if err != nil {
		return err
	}

	fmt.Println("Got course")
	if c.Value("current_user") != nil {
		course.MarkAsPurchased(tx, c.Value("current_user").(*models.User))
	}
	fmt.Println("Marked Purchased")
	c.Set("cour", course)
	c.Set("title", course.Course.Title)
	return c.Render(200, r.HTML(sitePath(c, "course.html")))
}

// CoursesShow default implementation.
func ClassroomShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	course, err := models.GetFullCourseBySlug(c.Param("slug"))
	if err != nil {
		return err
	}
	if c.Value("current_user") != nil {
		course.MarkAsPurchased(tx, c.Value("current_user").(*models.User))
	}
	c.Set("cour", course)
	c.Set("title", course.Course.Title)
	if !course.Course.Purchased {
		fmt.Println("redirecting because course isn't purchased")
		return c.Redirect(302, "/courses/"+c.Param("slug"))
	}

	c.Set("modules", course.Modules)

	return c.Render(200, r.HTML(sitePath(c, "classroom.html")))
}

// CoursesShow default implementation.
func ClassroomModuleShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	course, err := models.GetFullCourseBySlug(c.Param("slug"))
	if err != nil {
		return err
	}
	if c.Value("current_user") != nil {
		course.MarkAsPurchased(tx, c.Value("current_user").(*models.User))
	}
	c.Set("cour", course)

	if !course.Course.Purchased {
		return c.Redirect(302, "/courses/"+c.Param("slug"))
	}
	var found bool
	for _, m := range course.Modules {
		if m.ModuleSlug == c.Param("module") {
			c.Set("module", m)
			c.Set("title", m.Title)
			found = true
		}
	}
	if !found {
		return c.Error(400, errors.New("Module Not Found"))
	}

	return c.Render(200, r.HTML(sitePath(c, "module.html")))
}
