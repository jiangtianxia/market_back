package main

import (
	"market_back/conf"
	mlog "market_back/logger"
	"market_back/router"
	"market_back/store"
	"market_back/utils"

	"github.com/spf13/pflag"
)

var (
	DefaultConf     conf.DefaultConf
	SrvCnf          conf.ServerConf
	MarketRedisConf conf.RedisConf
)

func init() {
	// 基础配置
	{
		pflag.Int64Var(&DefaultConf.BucketRate, "bucket_rate", 1000, "bucket rate")
		pflag.Int64Var(&DefaultConf.BucketCapacity, "bucket_capacity", 5000, "bucket capacity")

		pflag.StringVar(&DefaultConf.JwtKey, "jwt_key", "h2wnknlsd", "jwt key")
		pflag.IntVar(&DefaultConf.JwtExpire, "jwt_expire", 1200, "jwt expire")
		pflag.StringVar(&DefaultConf.JwtIssuer, "jwt_issuer", "market", "jwt issuer")

		pflag.StringVar(&DefaultConf.WXAppID, "wx_appid", "wx7242543c6f6dcc38", "wx appid")
		pflag.StringVar(&DefaultConf.WXAppSecret, "wx_app_secret", "", "wx app secret")
	}

	// 服务配置
	{
		pflag.UintVar(&SrvCnf.HTTPPort, "http_port", 8888, "server http port")
		pflag.StringVar(&SrvCnf.LogPath, "log_path", "./log/market.log", "log path")
		pflag.BoolVar(&SrvCnf.OpenCORS, "cors", true, "open cors")
		pflag.StringSliceVar(&SrvCnf.AllowOrigins, "origins", []string{""}, "allow origins")
		pflag.BoolVar(&SrvCnf.IsReleaseMode, "release", false, "is release mode")
	}

	// redis
	{
		pflag.StringVar(&MarketRedisConf.RedisHost, "market_redis_host", "159.75.164.227", "market redis host")
		pflag.StringVar(&MarketRedisConf.RedisPassword, "market_redis_pass", "579021", "market redis password")
		pflag.IntVar(&MarketRedisConf.RedisPort, "market_redis_port", 6379, "market redis port")
		pflag.IntVar(&MarketRedisConf.RedisDB, "market_redis_db", 0, "market redis db")
	}

	pflag.Parse()
}

func main() {
	// 初始化logger
	mlog.InitLogger(SrvCnf.LogPath)

	// 初始化令牌桶
	utils.InitCurrentLimit(DefaultConf.BucketRate, DefaultConf.BucketCapacity)

	// 初始化redis
	store.InitRedis(&MarketRedisConf)
	defer store.RedisClose()

	// 初始化其他相关配置
	utils.InitJwt(DefaultConf.JwtKey, DefaultConf.JwtExpire, DefaultConf.JwtIssuer)
	utils.InitWXAPI(DefaultConf.WXAppID, DefaultConf.WXAppSecret)

	// 启动服务
	router.Start(&SrvCnf)
}
