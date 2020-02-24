package dorkgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dork *GoogleSearch

func TestSite(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Site("example.com").ToString()

	assert.Equal(t, result, "site:\"example.com\"", "they should be equal")
}

func TestIntext(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Intext("text").ToString()

	assert.Equal(t, result, "\"text\"", "they should be equal")
}

func TestInurl(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Inurl("index.php").ToString()

	assert.Equal(t, result, "inurl:\"index.php\"", "they should be equal")
}

func TestFiletype(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Filetype("pdf").ToString()

	assert.Equal(t, result, "filetype:\"pdf\"", "they should be equal")
}

func TestCache(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Cache("www.google.com").ToString()

	assert.Equal(t, result, "cache:\"www.google.com\"", "they should be equal")
}

func TestRelated(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Related("www.google.com").ToString()

	assert.Equal(t, result, "related:\"www.google.com\"", "they should be equal")
}

func TestExt(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Ext("(doc | pdf | xls | txt | xml)").ToString()

	assert.Equal(t, result, "ext:(doc | pdf | xls | txt | xml)", "they should be equal")
}

func TestExclude(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Exclude("html").Exclude("htm").Exclude("php").Exclude("md5sums").ToString()

	assert.Equal(t, result, "-html -htm -php -md5sums", "they should be equal")
}

func TestOr(t *testing.T) {
	dork = &GoogleSearch{}

	result := dork.Site("facebook.com").Or().Site("twitter.com").ToString()

	assert.Equal(t, result, "site:\"facebook.com\" OR site:\"twitter.com\"", "they should be equal")
}