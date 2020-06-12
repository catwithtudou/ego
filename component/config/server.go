package config

import "github.com/spf13/viper"

func init() {
	viper.SetDefault(systemNameKey, "app")
	viper.SetDefault(httpPortKey, 8080)
	viper.SetDefault(maxResponseLogSizeKey, 2000)
	viper.SetDefault(logPathKey, "/tmp/app.log")
	viper.SetDefault(traceHeaderKey, "gateway-trace")
	viper.SetDefault(httpRequestTimeoutKey, 3)
}

// server  服务配置
type server struct {
	// 服务名称
	Name string

	// 服务端口号
	Port int

	// 响应日志最大长度
	MaxResponseLogSize int

	// 日志路径
	LogPath string

	// jwt的key
	JwtSignKey []byte

	// trace header key
	TraceHeader string

	// http request timeout
	HttpRequestTimeOut int

	// 是否开启debug,开启后会显示debug信息
	Debug bool
}

const (
	systemNameKey         = "server.systemName"
	httpPortKey           = "server.httpPort"
	maxResponseLogSizeKey = "server.maxResponseLogSize"
	logPathKey            = "server.logPath"
	traceHeaderKey        = "server.traceHeader"
	httpRequestTimeoutKey = "server.httpRequestTimeout"
	jwtSignKey            = "server.jwtSign"
	debugKey              = "server.debug"
)