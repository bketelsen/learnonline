package grifts

import (
	"fmt"

	"github.com/bketelsen/learnonline/models"
	. "github.com/markbates/grift/grift"
	"github.com/markbates/pop"
)

var _ = Desc("seed:courses", "Deletes all the courses and purchases in the database and seeds new courses")
var _ = Add("seed:courses", func(c *Context) error {
	return models.DB.Transaction(func(tx *pop.Connection) error {
		for _, x := range []string{"courses", "purchases"} {
			err := tx.RawQuery(fmt.Sprintf("delete from %s", x)).Exec()
			if err != nil {
				return err
			}
		}

		// Seed Distributed Systems
		c := &models.Course{
			Slug:        "micro",
			Title:       "Microservices with Go Micro",
			Description: "Microservices with Go Micro",
			Price:       15000,
			Status:      "public",
		}
		verrs, err := tx.ValidateAndCreate(c)
		if verrs.HasAny() {
			return verrs
		}

		m := &models.Module{
			Title: "Microservices Overview",
			Slug: "micro-overview",
			Description: "In this module we will present a definition of microservices and explain some of the problems microservices attempt to solve. We will review common vocabulary and lay the groundwork for the rest of the class by defining the problem space.",
		}

		verrs, err = tx.ValidateAndCreate(m)
		if verrs.HasAny() {
			return verrs
		}
		for i, slug := range []string{"micro-overview"} {
			m := &models.Module{}
			err = tx.Where("slug = ?", slug).First(m)
			if err != nil {
				return err
			}
			err = tx.Create(&models.CourseModule{CourseID: c.ID, ModuleID: m.ID, Position: i})
			if err != nil {
				return err
			}
		}
		return err
	})
})

var _ = Desc("seed:purchases", "Deletes all the courses and purchases in the database and seeds new courses")
var _ = Add("seed:purchases", func(c *Context) error {
	return models.DB.Transaction(func(tx *pop.Connection) error {
		for _, x := range []string{"purchases"} {
			err := tx.RawQuery(fmt.Sprintf("delete from %s", x)).Exec()
			if err != nil {
				return err
			}
		}

		return nil
	})
})
