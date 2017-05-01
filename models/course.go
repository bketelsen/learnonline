package models

import (
	"fmt"
	"sort"

	"github.com/bketelsen/learnonlinecms/content"
	"github.com/markbates/pop"
)

// A Course represents a single point of purchase for a learning experience
type Course struct {
	Course  *content.Course
	Modules []content.Module
}

// URL is the web address of a course
func (c Course) URL() string {
	return fmt.Sprintf("/training/online/%s", c.Course.CourseSlug)
}

// MarkAsPurchased sets the purchased flag for a course
func (c Course) MarkAsPurchased(tx *pop.Connection, u *User) error {
	if c.Course.Price == 0 {
		c.Course.Purchased = true
		return nil
	}
	b, err := tx.Where("course_id = ? and user_id = ?", c.Course.ID, u.ID).Exists("purchases")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("purchased for this user:", b)
	c.Course.Purchased = b
	return err
}

// MarkPurchases marks a list of courses as purchased if the user has bought them.
func (cc Courses) MarkPurchases(tx *pop.Connection, u *User) error {
	for i, c := range cc {
		err := c.MarkAsPurchased(tx, u)
		if err != nil {
			return err
		}
		cc[i] = c
	}
	return nil
}

// Courses is a slice of Course
type Courses []Course

// GetFullCourse returns the course and related modules.
func GetFullCourse(id int) (Course, error) {
	var c Course
	pl, err := GetCourse(id)
	if err != nil {
		return c, err
	}
	c.Course = &pl
	for _, s := range pl.Modules {
		id, err := getID(s)
		if err != nil {
			fmt.Println(err, id, s)
			continue
		}
		sp, err := GetModule(id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.Modules = append(c.Modules, sp)
	}

	return c, nil
}

// GetFullCourseBySlug returns a full course searching by Slug instead of ID
func GetFullCourseBySlug(slug string) (Course, error) {
	var c Course
	pl, err := GetCourseBySlug(slug)
	if err != nil {
		return c, err
	}
	c.Course = &pl
	for _, s := range pl.Modules {
		id, err := getID(s)
		if err != nil {
			fmt.Println(err, id, s)
			continue
		}
		sp, err := GetModule(id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.Modules = append(c.Modules, sp)
	}

	return c, nil
}

// GetCourses returns all courses
func GetCourses() Courses {
	var pp Courses
	pl, err := GetCourseList()
	if err != nil {
		fmt.Println(err)
		return pp
	}

	for _, p := range pl {
		var pr Course
		pr.Course = &p
		for _, s := range p.Modules {
			id, err := getID(s)
			if err != nil {
				fmt.Println(err, id, s)
				continue
			}
			sp, err := GetModule(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			pr.Modules = append(pr.Modules, sp)
		}
		sort.Slice(pr.Modules, func(i, j int) bool { return pr.Modules[i].Order < pr.Modules[j].Order })

		pp = append(pp, pr)
	}

	return pp
}
