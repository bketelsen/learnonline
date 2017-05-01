package models

import (
	"fmt"

	"github.com/bketelsen/learnonlinecms/content"
	"github.com/markbates/pop"
)

// A Webinar represents a single point of purchase for a learning experience
type Webinar struct {
	Webinar  *content.Webinar
}

// URL is the web address of a webinar
func (c Webinar) URL() string {
	return fmt.Sprintf("/training/live/%s", c.Webinar.WebinarSlug)
}

// MarkAsPurchased sets the purchased flag for a webinar
func (c Webinar) MarkAsPurchased(tx *pop.Connection, u *User) error {
	if c.Webinar.Price == 0 {
		c.Webinar.Purchased = true
		return nil
	}
	b, err := tx.Where("course_id = ? and user_id = ?", c.Webinar.ID, u.ID).Exists("purchases")
	fmt.Println("purchased for this user:", b)
	c.Webinar.Purchased = b
	return err
}

// MarkPurchases marks a list of webinars as purchased if the user has bought them.
func (cc Webinars) MarkPurchases(tx *pop.Connection, u *User) error {
	for i, c := range cc {
		err := c.MarkAsPurchased(tx, u)
		if err != nil {
			return err
		}
		cc[i] = c
	}
	return nil
}

// Webinars is a slice of Webinar
type Webinars []Webinar

// GetFullCourseBySlug returns a full course searching by Slug instead of ID
func GetFullWebinarBySlug(slug string) (Webinar, error) {
	var c Webinar
	pl, err := GetWebinarBySlug(slug)
	if err != nil {
		return c, err
	}
	c.Webinar= &pl

	return c, nil
}
// GetWebinars returns all webinars
func GetWebinars() Webinars {
	var pp Webinars
	pl, err := GetWebinarList()
	if err != nil {
		fmt.Println(err)
		return pp
	}
	for _,w := range pl {
		pp = append(pp, Webinar{Webinar:&w})
	}
	return pp
}
