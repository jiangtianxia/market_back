package conf

type DefaultConf struct {
	BucketRate     int64
	BucketCapacity int64
	JwtKey         string
	JwtExpire      int
	JwtIssuer      string
	WXAppID        string
	WXAppSecret    string
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

type RedisConf struct {
	RedisHost     string
	RedisPassword string
	RedisPort     int
	RedisDB       int
}
