package main

import (
	"market_back/common/conf"
	mlog "market_back/logger"
	"market_back/router"
	"market_back/utils"

	"github.com/spf13/pflag"
)

var (
	DefaultConf conf.DefaultConf
	SrvCnf      conf.ServerConf
)

func init() {
	// 基础配置
	{
		pflag.Int64Var(&DefaultConf.BucketRate, "bucket_rate", 1000, "bucket rate")
		pflag.Int64Var(&DefaultConf.BucketCapacity, "bucket_capacity", 5000, "bucket capacity")
	}

	// 服务配置
	{
		pflag.UintVar(&SrvCnf.HTTPPort, "http_port", 8888, "server http port")
		pflag.StringVar(&SrvCnf.LogPath, "log_path", "./log/market.log", "log path")
		pflag.BoolVar(&SrvCnf.OpenCORS, "cors", true, "open cors")
		pflag.StringSliceVar(&SrvCnf.AllowOrigins, "origins", []string{""}, "allow origins")
		pflag.BoolVar(&SrvCnf.IsReleaseMode, "release", false, "is release mode")
	}

	pflag.Parse()
}

func main() {
	// 初始化logger
	mlog.InitLogger(SrvCnf.LogPath)

	// 初始化令牌桶
	utils.InitCurrentLimit(DefaultConf.BucketRate, DefaultConf.BucketCapacity)

	// 初始化服务
	router.Start(&SrvCnf)
}
