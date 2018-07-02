package config

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets92eb9aa9996a0b36b47b257d4eef1fba816fa60a = "name=easygin\ndebug=false\ntimezone=UTC\nlocale=en\nencryptKey=JJ*:NHN\":aaf9(ads0-(*;JFadFFasdf322GGe*"
var _Assets4e8207f328c5153b6d8461bdedb53cdbf588c7a6 = "[mysql]\nhost=127.0.0.1\nusername=root\ndbname=myinfo\npassword=ln6265431\nmaxIdle=10\n\n[redis]\nhost=127.0.0.1\nport=6379\npassword=\ndb=0\npoolSize=5\nmaxRetry=3"
var _Assetsd87e5f42318663b68b2d30a1eb12c159551bbe5d = "name=easygin-develop\ndebug=true\ntimezone=UTC\nlocale=en\nencryptKey=JJ*:NHN\":aaf9(ads0-(*;JFadFFasdf322GGe*"
var _Assets1e69113a9a49f438cdaaf2d7d6a6bb538f8b72d2 = "name=easygin\ndebug=false\ntimezone=UTC\nlocale=en\nencryptKey=JJ*:NHN\":aaf9(ads0-(*;JFadFFasdf322GGe*"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"config"}, "/config": []string{}, "/config/development": []string{"database.ini", "app.ini"}, "/config/sandbox": []string{"app.ini"}, "/config/production": []string{"app.ini"}}, map[string]*assets.File{
	"/config/production/app.ini": &assets.File{
		Path:     "/config/production/app.ini",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530512035, 1530512035859866914),
		Data:     []byte(_Assets92eb9aa9996a0b36b47b257d4eef1fba816fa60a),
	}, "/config": &assets.File{
		Path:     "/config",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530516563, 1530516563942702986),
		Data:     nil,
	}, "/config/development": &assets.File{
		Path:     "/config/development",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530528404, 1530528404976001594),
		Data:     nil,
	}, "/config/development/database.ini": &assets.File{
		Path:     "/config/development/database.ini",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530528404, 1530528404974579695),
		Data:     []byte(_Assets4e8207f328c5153b6d8461bdedb53cdbf588c7a6),
	}, "/config/development/app.ini": &assets.File{
		Path:     "/config/development/app.ini",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530516541, 1530516541850721172),
		Data:     []byte(_Assetsd87e5f42318663b68b2d30a1eb12c159551bbe5d),
	}, "/config/sandbox": &assets.File{
		Path:     "/config/sandbox",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530512040, 1530512040217261297),
		Data:     nil,
	}, "/config/sandbox/app.ini": &assets.File{
		Path:     "/config/sandbox/app.ini",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530512040, 1530512040215727648),
		Data:     []byte(_Assets1e69113a9a49f438cdaaf2d7d6a6bb538f8b72d2),
	}, "/config/production": &assets.File{
		Path:     "/config/production",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530512035, 1530512035861505382),
		Data:     nil,
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530528643, 1530528643985617881),
		Data:     nil,
	}}, "")
