package configer

type ConfigApi struct {
	HttpListen  string `config:"http_listen,required"`
	GrpcListen  string `config:"grpc_listen,required"`
	PostgresDsn string `config:"postgres_dsn,required"`
	LogFile     string `config:"log_file,required"`
	LogLevel    string `config:"log_level,required"`
}
