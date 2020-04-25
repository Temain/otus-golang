package configer

type Config struct {
	HttpListen  string `config:"httpListen,required"`
	GrpcListen  string `config:"grpcListen,required"`
	PostgresDsn string `config:"postgresDsn,required"`
	RabbitUrl   string `config:"rabbitUrl,required"`
	RabbitQueue string `config:"rabbitQueue,required"`
	LogFile     string `config:"logFile,required"`
	LogLevel    string `config:"logLevel,required"`
}
