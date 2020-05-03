package configer

type ConfigScheduler struct {
	PostgresDsn        string `config:"postgres_dsn,required"`
	RabbitUrl          string `config:"rabbit_url,required"`
	RabbitQueue        string `config:"rabbit_queue,required"`
	RabbitExchange     string `config:"rabbit_exchange,required"`
	RabbitExchangeType string `config:"rabbit_exchange_type,required"`
	LogFile            string `config:"log_file,required"`
	LogLevel           string `config:"log_level,required"`
}
