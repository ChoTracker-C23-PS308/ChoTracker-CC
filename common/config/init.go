package config

import "github.com/caarlos0/env/v6"

type (
	config struct {
		Port        int    `env:"PORT,unset" envDefault:"4001"`
		DatabaseURL string `env:"DATABASE_URL,unset"`
		//GMapAPIKey      string `env:"GMAP_API_KEY,unset"`

		Firebase
	}

	Firebase struct {
		CredentialType  string `env:"FIREBASE_CREDENTIAL_TYPE,unset"`
		CredentialValue string `env:"FIREBASE_CREDENTIAL_VALUE,unset"`
		//BucketName      string `env:"FIREBASE_STORAGE_BUCKET,unset"`
	}
)

func LoadConfig() *config {
	cfg := &config{}
	if err := env.Parse(cfg, env.Options{
		RequiredIfNoDef: true,
	}); err != nil {
		panic(err)
	}
	return cfg
}
