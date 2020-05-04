package configer

type ConfigApi struct {
	HttpHost    string `config:"http_host,required"`
	HttpPort    int    `config:"http_port,required"`
	GrpcHost    string `config:"grpc_host,required"`
	GrpcPort    int    `config:"grpc_port,required"`
	PostgresDsn string `config:"postgres_dsn,required"`
	LogFile     string `config:"log_file,required"`
	LogLevel    string `config:"log_level,required"`
}
