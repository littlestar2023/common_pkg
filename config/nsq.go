package config

type NsqConfig struct {
	LookupAddr         string `mapstructure:"lookup-addr" json:"lookup-addr" yaml:"lookup-addr"`
	LookupAuthUser     string `mapstructure:"lookup-auth-user" json:"lookup-auth-user" yaml:"lookup-auth-user"`
	LookupAuthPassword string `mapstructure:"lookup-auth-password" json:"lookup-auth-password" yaml:"lookup-auth-password"`
	LookupAuthMode     string `mapstructure:"lookup-auth-mode" json:"lookup-auth-mode" yaml:"lookup-auth-mode"`
	NSQDAddr           string `mapstructure:"nsqd-addr" json:"nsqd-addr" yaml:"nsqd-addr"`
	LocalCacheSize     int    `mapstructure:"local-cache-size" json:"local-cache-size" yaml:"local-cache-size"`
}
