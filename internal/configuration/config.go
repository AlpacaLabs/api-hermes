package configuration

import (
	"encoding/json"

	"github.com/rs/xid"

	configuration "github.com/AlpacaLabs/go-config"

	flag "github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	flagForGrpcPort = "grpc_port"
	flagForHTTPPort = "http_port"

	flagEmailEnabled = "email_enabled"
	flagSMSEnabled   = "sms_enabled"

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

	EmailEnabled bool
	SMSEnabled   bool

	TwilioAccountSID  string
	TwilioAuthToken   string
	TwilioPhoneNumber string

	// KafkaConfig provides configuration for connecting to Apache Kafka.
	KafkaConfig configuration.KafkaConfig
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
		AppName:  "api-hermes",
		AppID:    xid.New().String(),
		GrpcPort: 8081,
		HTTPPort: 8083,
	}

	c.KafkaConfig = configuration.LoadKafkaConfig()

	flag.Int(flagForGrpcPort, c.GrpcPort, "gRPC port")
	flag.Int(flagForHTTPPort, c.HTTPPort, "HTTP port")

	flag.Bool(flagEmailEnabled, c.EmailEnabled, "whether email sending is enabled")
	flag.Bool(flagSMSEnabled, c.SMSEnabled, "whether sms sending is enabled")

	flag.String(flagForTwilioAccountSID, c.TwilioAccountSID, "Twilio Account SID")
	flag.String(flagForTwilioAuthToken, c.TwilioAuthToken, "Twilio Auth Token")
	flag.String(flagForTwilioPhoneNumber, c.TwilioPhoneNumber, "Twilio Phone Number")

	flag.Parse()

	viper.BindPFlag(flagForGrpcPort, flag.Lookup(flagForGrpcPort))
	viper.BindPFlag(flagForHTTPPort, flag.Lookup(flagForHTTPPort))

	viper.BindPFlag(flagEmailEnabled, flag.Lookup(flagEmailEnabled))
	viper.BindPFlag(flagSMSEnabled, flag.Lookup(flagSMSEnabled))

	viper.BindPFlag(flagForTwilioAccountSID, flag.Lookup(flagForTwilioAccountSID))
	viper.BindPFlag(flagForTwilioAuthToken, flag.Lookup(flagForTwilioAuthToken))
	viper.BindPFlag(flagForTwilioPhoneNumber, flag.Lookup(flagForTwilioPhoneNumber))

	viper.AutomaticEnv()

	c.GrpcPort = viper.GetInt(flagForGrpcPort)
	c.HTTPPort = viper.GetInt(flagForHTTPPort)

	c.EmailEnabled = viper.GetBool(flagEmailEnabled)
	c.SMSEnabled = viper.GetBool(flagSMSEnabled)

	c.TwilioAccountSID = viper.GetString(flagForTwilioAccountSID)
	c.TwilioAuthToken = viper.GetString(flagForTwilioAuthToken)
	c.TwilioPhoneNumber = viper.GetString(flagForTwilioPhoneNumber)

	return c
}
