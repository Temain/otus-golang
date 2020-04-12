package configer

type Config struct {
	HttpListen string `config:"httpListen,required"`
	LogFile    string `config:"logFile,required"`
	LogLevel   string `config:"logLevel,required"`
}
