package config

type (
	Producer struct {
		SchemaPaths      []string `env:"SCHEMA_PATHS"`
		BootstrapServers string   `env:"BOOTSTRAP_SERVERS"`
		Topic            string   `env:"TOPIC"`
	}

	Consumer struct {
		SchemaPaths      []string `env:"SCHEMA_PATHS"`
		BootstrapServers string   `env:"BOOTSTRAP_SERVERS"`
		Topic            string   `env:"TOPIC"`
		ConsumerGroupID  string   `env:"CONSUMER_GROUP_ID"`
	}
)
