package conf

type DefaultConf struct {
	BucketRate     int64
	BucketCapacity int64
}

type ServerConf struct {
	HTTPPort        uint
	Logstash        bool
	LogstashConn    string
	LogEnableFile   bool
	LogPath         string
	LogRotationTime int64
	LogMaxAge       int64
	AllowOrigins    []string
	OpenCORS        bool
	IsReleaseMode   bool
}
