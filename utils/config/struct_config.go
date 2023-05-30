package config

type config struct {
	LogLevel        uint32   `yaml:"logLevel"`
	FrontendAddress string   `yaml:"frontAddress"`
	BackendAddress  string   `yaml:"backAddress"`
	ServerHeader    string   `yaml:"serverHeader"`
	Cors            []string `yaml:"cors"`
	MySqlDsn        string   `yaml:"mySqlDsn"`
	MySqlMigrate    bool     `yaml:"mySqlMigrate"`
	ProductionURL   string   `yaml:"productionURL"`
	Path            string   `yaml:"path"`
}
