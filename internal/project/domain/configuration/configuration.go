package configuration

const (
	EnvPrefix      = `tc`
	FullEnvPrefix  = EnvPrefix + "_"
	LayerSeparator = `_`

	// DevelopmentMode отвечает за запуск приложения в development режиме.
	DevelopmentMode        = `dev`
	DefaultDevelopmentMode = false

	// LogLayer Раздел настроек логирования
	LogLayer             = `log`
	LogLevel             = LogLayer + LayerSeparator + `level`
	DefaultLogLevel      = `info`
	LogMaxAge            = LogLayer + LayerSeparator + "maxAge"
	DefaultLogMaxAge     = 0
	LogMaxBackups        = LogLayer + LayerSeparator + "maxBackups"
	DefaultLogMaxBackups = 10
	LogMaxSize           = LogLayer + LayerSeparator + "maxSize"
	DefaultLogMaxSize    = 30

	// DirLayer Раздел настроек директорий
	DirLayer      = `dir`
	DirApp        = DirLayer + LayerSeparator + `app`
	DirBin        = DirLayer + LayerSeparator + `bin`
	DefaultDirBin = `bin`
	DirEtc        = DirLayer + LayerSeparator + `etc`
	DefaultDirEtc = `etc`
	DirVar        = DirLayer + LayerSeparator + `var`
	DefaultDirVar = `var`
	DirLog        = DirLayer + LayerSeparator + `log`
	DirConfig     = DirLayer + LayerSeparator + `config`

	// FileLayer Раздел настроек файлов
	FileLayer             = `file`
	FileLogName           = FileLayer + LayerSeparator + `logName`
	DefaultFileLogName    = `project.log`
	FileConfigName        = FileLayer + LayerSeparator + `configName`
	DefaultFileConfigName = `project`

	// HttpLayer Раздел настроек веб-сервера
	HttpLayer       = `http`
	HttpPort        = HttpLayer + LayerSeparator + `port`
	DefaultHttpPort = `80`
	HttpHost        = HttpLayer + LayerSeparator + `host`
	DefaultHttpHost = `localhost`

	// DatabaseLayer Раздел настроек базы данных
	DatabaseLayer = `database`
	DatabaseConn  = DatabaseLayer + LayerSeparator + `connection`

	// DataStoreLayer Раздел настроек типа хранения данных
	DataStoreLayer       = `data`
	DataStoreType        = DataStoreLayer + LayerSeparator + `storeType`
	DataStoreTypeReg     = `reg`
	DataStoreTypeDb      = `db`
	DataStoreTypeMongo   = `mongo`
	DefaultDataStoreType = DataStoreTypeDb

	// SyncLayer Раздел настроек синхронизации
	SyncLayer              = `sync`
	SyncDeleteUsersTime    = SyncLayer + LayerSeparator + `deleteUsersTime`
	DefaultDeleteUsersTime = 24

	// UsersLayer Раздел настроек пользователей
	UsersLayer                       = `users`
	UsersOverdueTimeInSeconds        = UsersLayer + LayerSeparator + `overdueTimeInSeconds`
	DefaultUsersOverdueTimeInSeconds = 60

	//MongoLayer Раздел настроек соединения с MongoDB
	MongoLayer             = `mongo`
	MongoConnection        = MongoLayer + LayerSeparator + `connection`
	DefaultMongoConnection = `mongodb://localhost:27017`
	MongoDbName            = MongoLayer + LayerSeparator + `db`
	DefaultMongoDbName     = `project`
)
