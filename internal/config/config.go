package config

import (
	"encoding/json"

	"github.com/google/uuid"
	flag "github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	flagForGrpcPort       = "grpc_port"
	flagForGrpcPortHealth = "grpc_port_health"
	flagForHTTPPort       = "http_port"

	flagForTwilioAccountSID  = "twilio_account_sid"
	flagForTwilioAuthToken   = "twilio_auth_token"
	flagForTwilioPhoneNumber = "twilio_phone_number"
)

type Config struct {
	// AppName is a low cardinality identifier for this service.
	AppName string

	// AppID is a unique identifier for the instance (pod) running this app.
	AppID string

	// GrpcPort controls what port our gRPC server runs on.
	GrpcPort int

	// HTTPPort controls what port our HTTP server runs on.
	HTTPPort int

	// HealthPort controls what port our gRPC health endpoints run on.
	HealthPort int

	TwilioAccountSID  string
	TwilioAuthToken   string
	TwilioPhoneNumber string
}

func (c Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Could not marshal config to string: %v", err)
	}
	return string(b)
}

func LoadConfig() Config {
	c := Config{
		AppName:    "citadel",
		AppID:      uuid.New().String(),
		GrpcPort:   8081,
		HealthPort: 8082,
		HTTPPort:   8083,
	}

	flag.Int(flagForGrpcPort, c.GrpcPort, "gRPC port")
	flag.Int(flagForGrpcPortHealth, c.HealthPort, "gRPC health port")
	flag.Int(flagForHTTPPort, c.HTTPPort, "gRPC HTTP port")

	flag.String(flagForTwilioAccountSID, c.TwilioAccountSID, "Twilio Account SID")
	flag.String(flagForTwilioAuthToken, c.TwilioAuthToken, "Twilio Auth Token")
	flag.String(flagForTwilioPhoneNumber, c.TwilioPhoneNumber, "Twilio Phone Number")

	flag.Parse()

	viper.BindPFlag(flagForGrpcPort, flag.Lookup(flagForGrpcPort))
	viper.BindPFlag(flagForGrpcPortHealth, flag.Lookup(flagForGrpcPortHealth))
	viper.BindPFlag(flagForHTTPPort, flag.Lookup(flagForHTTPPort))

	viper.BindPFlag(flagForTwilioAccountSID, flag.Lookup(flagForTwilioAccountSID))
	viper.BindPFlag(flagForTwilioAuthToken, flag.Lookup(flagForTwilioAuthToken))
	viper.BindPFlag(flagForTwilioPhoneNumber, flag.Lookup(flagForTwilioPhoneNumber))

	viper.AutomaticEnv()

	c.GrpcPort = viper.GetInt(flagForGrpcPort)
	c.HealthPort = viper.GetInt(flagForGrpcPortHealth)
	c.HTTPPort = viper.GetInt(flagForHTTPPort)

	c.TwilioAccountSID = viper.GetString(flagForTwilioAccountSID)
	c.TwilioAuthToken = viper.GetString(flagForTwilioAuthToken)
	c.TwilioPhoneNumber = viper.GetString(flagForTwilioPhoneNumber)

	return c
}
