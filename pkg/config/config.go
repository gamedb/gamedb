package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const EnvProd = "production"
const EnvLocal = "local"
const EnvConsumer = "consumer"

const prefix = "STEAM_"

var Config BaseConfig

func init() {

	Config.AdminName.Set("ADMIN_NAME")
	Config.AdminEmail.Set("ADMIN_EMAIL")
	Config.AdminSteamID.Set("ADMIN_STEAM_ID")

	Config.RabbitUsername.Set("RABBIT_USER")
	Config.RabbitPassword.Set("RABBIT_PASS")
	Config.RabbitHost.Set("RABBIT_HOST")
	Config.RabbitPort.Set("RABBIT_PORT")
	Config.RabbitManagmentPort.Set("RABBIT_MANAGEMENT_PORT")

	Config.SessionAuthentication.Set("SESSION_AUTHENTICATION")
	Config.SessionEncryption.Set("SESSION_ENCRYPTION")

	Config.InstagramUsername.Set("INSTAGRAM_USERNAME")
	Config.InstagramPassword.Set("INSTAGRAM_PASSWORD")

	Config.MySQLHost.Set("MYSQL_HOST")
	Config.MySQLPort.Set("MYSQL_PORT")
	Config.MySQLUsername.Set("MYSQL_USERNAME")
	Config.MySQLPassword.Set("MYSQL_PASSWORD")
	Config.MySQLDatabase.Set("MYSQL_DATABASE")

	Config.RecaptchaPublic.Set("RECAPTCHA_PUBLIC")
	Config.RecaptchaPrivate.Set("RECAPTCHA_PRIVATE")

	Config.TwitchClientID.Set("TWITCH_CLIENT_ID")
	Config.TwitchClientSecret.Set("TWITCH_CLIENT_SECRET")

	Config.InfluxURL.Set("INFLUX_URL")
	Config.InfluxUsername.Set("INFLUX_USERNAME")
	Config.InfluxPassword.Set("INFLUX_PASSWORD")

	Config.MongoHost.Set("MONGO_HOST")
	Config.MongoPort.Set("MONGO_PORT")
	Config.MongoUsername.Set("MONGO_USERNAME")
	Config.MongoPassword.Set("MONGO_PASSWORD")
	Config.MongoDatabase.Set("MONGO_DATABASE")

	Config.TwitterAccessToken.Set("TWITTER_ACCESS_TOKEN")
	Config.TwitterAccessTokenSecret.Set("TWITTER_ACCESS_TOKEN_SECRET")
	Config.TwitterConsumerKey.Set("TWITTER_CONSUMER_KEY")
	Config.TwitterConsumerSecret.Set("TWITTER_CONSUMER_SECRET")

	Config.PatreonSecret.Set("PATREON_WEBOOK_SECRET")
	Config.PatreonClientID.Set("PATREON_CLIENT_ID")
	Config.PatreonClientSecret.Set("PATREON_CLIENT_SECRET")

	Config.DiscordClientID.Set("DISCORD_CLIENT_ID")
	Config.DiscordSescret.Set("DISCORD_SECRET")

	Config.DiscordChatBotToken.Set("DISCORD_BOT_TOKEN")
	Config.DiscordRelayBotToken.Set("DISCORD_RELAY_TOKEN")
	Config.DiscordChangesBotToken.Set("DISCORD_CHANGES_BOT_TOKEN")

	Config.Path.Set("PATH")
	Config.AssetsPath.Set("ASSETS_PATH")
	Config.TemplatesPath.Set("TEMPLATES_PATH")

	Config.GameDBDomain.Set("DOMAIN")
	Config.Environment.Set("ENV")
	Config.GithubToken.Set("GITHUB_TOKEN")
	Config.GoogleBucket.Set("GOOGLE_BUCKET")
	Config.GoogleProject.Set("GOOGLE_PROJECT")
	Config.SendGridAPIKey.Set("SENDGRID")
	Config.SteamAPIKey.Set("API_KEY")
	Config.WebserverPort.Set("PORT")

	// Defaults
	Config.GameDBShortName.SetDefault("GameDB")
	Config.NewReleaseDays.SetDefault("14")

	switch Config.Environment.Get() {
	case EnvProd:

		Config.MemcacheDSN.SetDefault("memcache:11211")

	case EnvLocal:

		Config.MemcacheDSN.SetDefault("localhost:11211")

	case EnvConsumer:

	default:
		fmt.Println("Unknown environment")
		os.Exit(1)
	}
}

type BaseConfig struct {
	AdminName    ConfigItem
	AdminEmail   ConfigItem
	AdminSteamID ConfigItem

	DiscordClientID        ConfigItem
	DiscordSescret         ConfigItem
	DiscordRelayBotToken   ConfigItem
	DiscordChatBotToken    ConfigItem
	DiscordChangesBotToken ConfigItem

	InfluxURL      ConfigItem
	InfluxPassword ConfigItem
	InfluxUsername ConfigItem

	InstagramPassword ConfigItem
	InstagramUsername ConfigItem

	MongoHost     ConfigItem
	MongoPort     ConfigItem
	MongoUsername ConfigItem
	MongoPassword ConfigItem
	MongoDatabase ConfigItem

	MySQLHost     ConfigItem
	MySQLPort     ConfigItem
	MySQLUsername ConfigItem
	MySQLPassword ConfigItem
	MySQLDatabase ConfigItem

	RabbitUsername      ConfigItem
	RabbitPassword      ConfigItem
	RabbitHost          ConfigItem
	RabbitPort          ConfigItem
	RabbitManagmentPort ConfigItem

	RecaptchaPrivate ConfigItem
	RecaptchaPublic  ConfigItem

	SessionAuthentication ConfigItem
	SessionEncryption     ConfigItem

	TwitchClientID     ConfigItem
	TwitchClientSecret ConfigItem

	TwitterAccessToken       ConfigItem
	TwitterAccessTokenSecret ConfigItem
	TwitterConsumerKey       ConfigItem
	TwitterConsumerSecret    ConfigItem

	PatreonSecret       ConfigItem
	PatreonClientID     ConfigItem
	PatreonClientSecret ConfigItem

	Path          ConfigItem
	AssetsPath    ConfigItem
	TemplatesPath ConfigItem

	Environment     ConfigItem
	GameDBDomain    ConfigItem
	GameDBShortName ConfigItem
	GithubToken     ConfigItem
	GoogleBucket    ConfigItem
	GoogleProject   ConfigItem
	MemcacheDSN     ConfigItem
	SendGridAPIKey  ConfigItem
	SteamAPIKey     ConfigItem
	WebserverPort   ConfigItem
	CommitHash      ConfigItem
	NewReleaseDays  ConfigItem
}

// ConfigItem
type ConfigItem struct {
	value        string
	defaultValue string
}

func (ci *ConfigItem) Set(environment string) {
	env, b := os.LookupEnv(prefix + environment)
	if !b {
		fmt.Println("MISSING ENV: " + environment)
	}
	ci.value = env
}

func (ci *ConfigItem) SetDefault(defaultValue string) {
	ci.defaultValue = defaultValue
}

func (ci ConfigItem) Get() string {
	if ci.value != "" {
		return ci.value
	}
	return ci.defaultValue
}

func (ci ConfigItem) GetBool() bool {
	b, err := strconv.ParseBool(ci.Get())
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func (ci ConfigItem) GetInt() int {
	i, err := strconv.Atoi(ci.Get())
	if err != nil {
		fmt.Println(err)
	}
	return i
}

//
func RabbitDSN() string {
	return "amqp://" + Config.RabbitUsername.Get() + ":" + Config.RabbitPassword.Get() + "@" + Config.RabbitHost.Get() + ":" + Config.RabbitPort.Get()
}

func MySQLDNS() string {
	return Config.MySQLUsername.Get() + ":" + Config.MySQLPassword.Get() + "@tcp(" + Config.MySQLHost.Get() + ":" + Config.MySQLPort.Get() + ")/" + Config.MySQLDatabase.Get()
}

func MongoDSN() string {
	return "mongodb://" + Config.MongoHost.Get() + ":" + Config.MongoPort.Get()
}

// func RabbitAPI(values url.Values) string {
// 	return "http://" + Config.RabbitHost.Get() + ":" + Config.RabbitManagmentPort.Get() + "/api/overview?" + values.Encode()
// }

func ListenOn() string {
	return "0.0.0.0:" + Config.WebserverPort.Get()
}

func HasMemcache() bool {
	return Config.MemcacheDSN.Get() != ""
}

func IsLocal() bool {
	return Config.Environment.Get() == EnvLocal
}

func IsProd() bool {
	return Config.Environment.Get() == EnvProd
}

// func IsConsumer() bool {
// 	return Config.Environment.Get() == EnvConsumer
// }

func GetSteamKeyTag() string {

	key := Config.SteamAPIKey.Get()
	if len(key) > 5 {
		key = key[0:5]
	}

	return strings.ToUpper(key)
}
