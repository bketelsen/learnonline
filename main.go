package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bketelsen/learnonline/actions"
	"github.com/bketelsen/learnonline/models"
	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

func init() {
	LoadConfig()

}

var lookupPaths = []string{"", "./config", "/config", "../", "../config", "../..", "../../config"}
var ConfigName = "sites.yml"
var Debug = true

func main() {
	port := envy.Get("PORT", "3000")

	baseURL := envy.Get("CMS_URL", "http://127.0.0.1:8080")
	models.BaseURL = baseURL
	log.Printf("Sites: %#v\n", actions.Sites)
	log.Printf("Starting learnonline on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}

func LookupPaths() []string {
	return lookupPaths
}
func LoadConfig() {
	path, err := findConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	sites, err := loadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	actions.Sites = sites

}
func findConfigPath() (string, error) {
	for _, p := range LookupPaths() {
		path, _ := filepath.Abs(filepath.Join(p, ConfigName))
		if _, err := os.Stat(path); err == nil {
			return path, err
		}
	}
	return "", errors.New("[learnonline]: Tried to load configuration file, but couldn't find it.")
}

func loadConfig(path string) (map[string]*models.Site, error) {
	if Debug {
		fmt.Printf("[learnonline]: Loading config file from %s\n", path)
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't read file %s", path)
	}

	sites := map[string]*models.Site{}
	err = yaml.Unmarshal(b, &sites)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't unmarshal config to yaml")
	}
	return sites, nil
}
