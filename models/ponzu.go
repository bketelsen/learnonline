/*
* CODE GENERATED AUTOMATICALLY WITH github.com/bketelsen/ponzigen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package models

import (
	"github.com/bketelsen/learnonlinecms/content"
	"github.com/bketelsen/ponzi"
	"github.com/pkg/errors"
	"time"
)

var BaseURL string

type CourseListResult struct {
	Data []content.Course `json:"data"`
}
type ModuleListResult struct {
	Data []content.Module `json:"data"`
}
type PostListResult struct {
	Data []content.Post `json:"data"`
}
type WebinarListResult struct {
	Data []content.Webinar `json:"data"`
}

var courseCache *ponzi.Cache
var moduleCache *ponzi.Cache
var postCache *ponzi.Cache
var webinarCache *ponzi.Cache

func initCourseCache() {
	if courseCache == nil {
		courseCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initModuleCache() {
	if moduleCache == nil {
		moduleCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initPostCache() {
	if postCache == nil {
		postCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initWebinarCache() {
	if webinarCache == nil {
		webinarCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}

func GetCourse(id int) (content.Course, error) {
	initCourseCache()
	var sp CourseListResult
	err := courseCache.Get(id, "Course", &sp)
	if err != nil {
		return content.Course{}, err
	}
	if len(sp.Data) == 0 {
		return content.Course{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetModule(id int) (content.Module, error) {
	initModuleCache()
	var sp ModuleListResult
	err := moduleCache.Get(id, "Module", &sp)
	if err != nil {
		return content.Module{}, err
	}
	if len(sp.Data) == 0 {
		return content.Module{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetPost(id int) (content.Post, error) {
	initPostCache()
	var sp PostListResult
	err := postCache.Get(id, "Post", &sp)
	if err != nil {
		return content.Post{}, err
	}
	if len(sp.Data) == 0 {
		return content.Post{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetWebinar(id int) (content.Webinar, error) {
	initWebinarCache()
	var sp WebinarListResult
	err := webinarCache.Get(id, "Webinar", &sp)
	if err != nil {
		return content.Webinar{}, err
	}
	if len(sp.Data) == 0 {
		return content.Webinar{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}

func GetCourseBySlug(slug string) (content.Course, error) {
	initCourseCache()
	var sp CourseListResult
	err := courseCache.GetBySlug(slug, "Course", &sp)
	if err != nil {
		return content.Course{}, err
	}
	if len(sp.Data) == 0 {
		return content.Course{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetModuleBySlug(slug string) (content.Module, error) {
	initModuleCache()
	var sp ModuleListResult
	err := moduleCache.GetBySlug(slug, "Module", &sp)
	if err != nil {
		return content.Module{}, err
	}
	if len(sp.Data) == 0 {
		return content.Module{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetPostBySlug(slug string) (content.Post, error) {
	initPostCache()
	var sp PostListResult
	err := postCache.GetBySlug(slug, "Post", &sp)
	if err != nil {
		return content.Post{}, err
	}
	if len(sp.Data) == 0 {
		return content.Post{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetWebinarBySlug(slug string) (content.Webinar, error) {
	initWebinarCache()
	var sp WebinarListResult
	err := webinarCache.GetBySlug(slug, "Webinar", &sp)
	if err != nil {
		return content.Webinar{}, err
	}
	if len(sp.Data) == 0 {
		return content.Webinar{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}

func GetCourseList() ([]content.Course, error) {
	initCourseCache()
	var sp CourseListResult
	err := courseCache.GetAll("Course", &sp)
	if err != nil {
		return []content.Course{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Course{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetModuleList() ([]content.Module, error) {
	initModuleCache()
	var sp ModuleListResult
	err := moduleCache.GetAll("Module", &sp)
	if err != nil {
		return []content.Module{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Module{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetPostList() ([]content.Post, error) {
	initPostCache()
	var sp PostListResult
	err := postCache.GetAll("Post", &sp)
	if err != nil {
		return []content.Post{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Post{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetWebinarList() ([]content.Webinar, error) {
	initWebinarCache()
	var sp WebinarListResult
	err := webinarCache.GetAll("Webinar", &sp)
	if err != nil {
		return []content.Webinar{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Webinar{}, errors.New("Not Found")
	}
	return sp.Data, err

}
