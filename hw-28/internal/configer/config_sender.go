package configer

type ConfigSender struct {
	RabbitUrl          string `config:"rabbitUrl,required"`
	RabbitQueue        string `config:"rabbitQueue,required"`
	RabbitExchange     string `config:"rabbitExchange,required"`
	RabbitExchangeType string `config:"rabbitExchangeType,required"`
	LogFile            string `config:"logFile,required"`
	LogLevel           string `config:"logLevel,required"`
}
