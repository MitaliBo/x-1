package xconfig

import (
	"flag"
	"os"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
)

// Load config settings for config object(should be a struct pointer).
func Load(config interface{}, setFlags ...func()) {
	var (
		envFile    string
		configFile string
	)
	for _, setFlag := range setFlags {
		setFlag()
	}

	flag.StringVar(&envFile, "e", ".env", "env file")
	flag.StringVar(&configFile, "c", ".config.yml", "config file")
	flag.Parse()
	godotenv.Load(envFile)
	os.Setenv("CONFIGOR_ENV_PREFIX", "-")
	if fileInfo, err := os.Stat(configFile); err == nil && fileInfo.Mode().IsRegular() {
		configor.Load(config, configFile)
	} else {
		configor.Load(config)
	}
}

// DumpSample 生成默认值的配置数据, 可以据此生成sample config file.
func DumpSample(config interface{}) map[string]interface{} {
	// TODO:
	return nil
}

// SyncValues sync values.
type SyncValues struct {
	pairs []pair
}

type pair struct {
	value  interface{}
	target interface{}
}

// AddBool add a bool value->target pair.
func (rv *SyncValues) AddBool(value, target *bool) *bool {
	rv.pairs = append(rv.pairs, pair{value, target})
	return value
}

// AddString add a string value->target pair.
func (rv *SyncValues) AddString(value, target *string) *string {
	rv.pairs = append(rv.pairs, pair{value, target})
	return value
}

// AddInt add a int value->target pair.
func (rv *SyncValues) AddInt(value, target *int) *int {
	rv.pairs = append(rv.pairs, pair{value, target})
	return value
}

// Sync values to targets.
func (rv *SyncValues) Sync() {
	for _, pair := range rv.pairs {
		switch v := (pair.value).(type) {
		case *bool:
			*(pair.target).(*bool) = *v
		case *string:
			*(pair.target).(*string) = *v
		case *int:
			*(pair.target).(*int) = *v
		}
	}
}
