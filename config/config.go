package config

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsb276aa3ee4dc386a0202e613b8367c15e4a04ecf = "[logging]\nrequestLog = request\nlevel = debug\nchannel = file\n[logging.file]\nmode = daily\nmodule = true\nbasePath = /tmp/log\ndays = 7"
var _Assets1e69113a9a49f438cdaaf2d7d6a6bb538f8b72d2 = "name=easygin\ndebug=false\ntimezone=UTC\nlocale=en\nencryptKey=JJ*:NHN\":aaf9(ads0-(*;JFadFFasdf322GGe*"
var _Assets1b98c2870417114de9060008c289a10c5358ab73 = ""
var _Assets4e8207f328c5153b6d8461bdedb53cdbf588c7a6 = "[mysql]\nhost = 127.0.0.1\nusername = root\ndbname = myinfo\npassword = ln6265431\nmaxIdle = 10\n\n[redis]\nhost = 127.0.0.1\nport = 6379\npassword =\ndb = 0\npoolSize = 5\nmaxRetry = 3"
var _Assetsd87e5f42318663b68b2d30a1eb12c159551bbe5d = "name=easygin-develop\ndebug=true\ntimezone=UTC\nlocale=en\nencryptKey=JJ*:NHN\":aaf9(ads0-(*;JFadFFasdf322GGe*"
var _Assets92eb9aa9996a0b36b47b257d4eef1fba816fa60a = "name=easygin\ndebug=false\ntimezone=UTC\nlocale=en\nencryptKey=JJ*:NHN\":aaf9(ads0-(*;JFadFFasdf322GGe*"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"config"}, "/config": []string{"config.go"}, "/config/development": []string{"logging.ini", "database.ini", "app.ini"}, "/config/sandbox": []string{"app.ini"}, "/config/production": []string{"app.ini"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530902413, 1530902413518252850),
		Data:     nil,
	}, "/config/config.go": &assets.File{
		Path:     "/config/config.go",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530723256, 1530723256692674110),
		Data:     []byte(_Assets1b98c2870417114de9060008c289a10c5358ab73),
	}, "/config/development/database.ini": &assets.File{
		Path:     "/config/development/database.ini",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530532428, 1530532428082344637),
		Data:     []byte(_Assets4e8207f328c5153b6d8461bdedb53cdbf588c7a6),
	}, "/config/development/app.ini": &assets.File{
		Path:     "/config/development/app.ini",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530516541, 1530516541850721172),
		Data:     []byte(_Assetsd87e5f42318663b68b2d30a1eb12c159551bbe5d),
	}, "/config/production": &assets.File{
		Path:     "/config/production",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530512035, 1530512035861505382),
		Data:     nil,
	}, "/config/production/app.ini": &assets.File{
		Path:     "/config/production/app.ini",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530512035, 1530512035859866914),
		Data:     []byte(_Assets92eb9aa9996a0b36b47b257d4eef1fba816fa60a),
	}, "/config": &assets.File{
		Path:     "/config",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530723130, 1530723130352723475),
		Data:     nil,
	}, "/config/development": &assets.File{
		Path:     "/config/development",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1530899664, 1530899664539752382),
		Data:     nil,
	}, "/config/development/logging.ini": &assets.File{
		Path:     "/config/development/logging.ini",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1530899664, 1530899664539258285),
		Data:     []byte(_Assetsb276aa3ee4dc386a0202e613b8367c15e4a04ecf),
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
	}}, "")
