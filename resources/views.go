package resources

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets4dca0144e4c8f5ca33ec111881758340a5539f68 = "<!doctype html>\n<html>\n\n<head>\n    <title>{{ .title }}</title>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <meta charset=\"UTF-8\">\n\n</head>\n\n<body class=\"container\">\n\n<h1>hello world layout!</h1>\n</body>\n</html>"
var _Assets4587821e4798bf19a900a1038deedfa61738f8bf = ""
var _Assetsa7fbb3de6c068e82d0f96351369e88869e8ef1b1 = ""
var _Assets8d99f7461968ba15fd7779f56d344eabee04d0d5 = "<!doctype html>\n<html>\n\n<head>\n    <title>{{ .title }}</title>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <meta charset=\"UTF-8\">\n\n</head>\n\n<body class=\"container\">\n\n<h1>hello world!</h1>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/views/layouts": []string{"app.html"}, "/": []string{"views"}, "/views": []string{"index.html"}, "/views/user": []string{"user.html"}, "/views/user/password": []string{"change.html"}}, map[string]*assets.File{
	"/views": &assets.File{
		Path:     "/views",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530206706, 1530206706545174147),
		Data:     nil,
	}, "/views/layouts": &assets.File{
		Path:     "/views/layouts",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530206217, 1530206217862031727),
		Data:     nil,
	}, "/views/layouts/app.html": &assets.File{
		Path:     "/views/layouts/app.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530206217, 1530206217860423407),
		Data:     []byte(_Assets4dca0144e4c8f5ca33ec111881758340a5539f68),
	}, "/views/user/password": &assets.File{
		Path:     "/views/user/password",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530206719, 1530206719127243572),
		Data:     nil,
	}, "/views/user/password/change.html": &assets.File{
		Path:     "/views/user/password/change.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530206719, 1530206719127225180),
		Data:     []byte(_Assets4587821e4798bf19a900a1038deedfa61738f8bf),
	}, "/views/user/user.html": &assets.File{
		Path:     "/views/user/user.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530206747, 1530206747700855807),
		Data:     []byte(_Assetsa7fbb3de6c068e82d0f96351369e88869e8ef1b1),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530241995, 1530241995860904590),
		Data:     nil,
	}, "/views/index.html": &assets.File{
		Path:     "/views/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530200595, 1530200595213989334),
		Data:     []byte(_Assets8d99f7461968ba15fd7779f56d344eabee04d0d5),
	}, "/views/user": &assets.File{
		Path:     "/views/user",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530206747, 1530206747700903242),
		Data:     nil,
	}}, "")
