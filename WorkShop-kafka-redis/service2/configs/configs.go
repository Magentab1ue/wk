package configs

type Config struct {
	App      Fiber
	Postgres PostgresSql
	Kafkas     Kafka
	Redis      Redis
}

type Redis struct {
	Host string
	Port string
}

type Fiber struct {
	Port string
}

type PostgresSql struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	SslMode      string
}

type Kafka struct {
	Hosts []string
	Group string
}
