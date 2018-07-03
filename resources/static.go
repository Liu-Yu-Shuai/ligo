package resources

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _StaticAssertsa75c2d05e7abb6eb0285ea5709cfc49c4fe9bd12 = "alert(\"test\");"

// StaticAsserts returns go-assets FileSystem
var StaticAsserts = assets.NewFileSystem(map[string][]string{"/": []string{"static"}, "/static": []string{}, "/static/js": []string{"index.js"}}, map[string]*assets.File{
	"/static/js/index.js": &assets.File{
		Path:     "/static/js/index.js",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530590079, 1530590079533590077),
		Data:     []byte(_StaticAssertsa75c2d05e7abb6eb0285ea5709cfc49c4fe9bd12),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530590134, 1530590134840501366),
		Data:     nil,
	}, "/static": &assets.File{
		Path:     "/static",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530589941, 1530589941772633848),
		Data:     nil,
	}, "/static/js": &assets.File{
		Path:     "/static/js",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530590079, 1530590079534371357),
		Data:     nil,
	}}, "")
