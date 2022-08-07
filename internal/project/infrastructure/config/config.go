package config

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"projectname/internal/project/domain/configuration"
	"strings"
)

const ServiceName = `DefaultAppConfiguration`

func New() *viper.Viper {
	return viper.NewWithOptions(viper.KeyDelimiter(configuration.LayerSeparator))
}

// SetDefaults устанавливает дефолтные значения конфигурации приложения
func SetDefaults(cfg *viper.Viper) error {
	cfg.SetDefault(configuration.LogLevel, configuration.DefaultLogLevel)
	cfg.SetDefault(configuration.HttpHost, configuration.DefaultHttpHost)
	cfg.SetDefault(configuration.HttpPort, configuration.DefaultHttpPort)
	cfg.SetDefault(configuration.FileLogName, configuration.DefaultFileLogName)
	cfg.SetDefault(configuration.LogMaxAge, configuration.DefaultLogMaxAge)
	cfg.SetDefault(configuration.LogMaxSize, configuration.DefaultLogMaxSize)
	cfg.SetDefault(configuration.LogMaxBackups, configuration.DefaultLogMaxBackups)
	cfg.SetDefault(configuration.DevelopmentMode, configuration.DefaultDevelopmentMode)
	cfg.SetDefault(configuration.DatabaseConn, configuration.DefaultDatabaseConn)
	cfg.SetDefault(configuration.FileConfigName, configuration.DefaultFileConfigName)
	cfg.SetDefault(configuration.DataStoreType, configuration.DefaultDataStoreType)
	cfg.SetDefault(configuration.SyncDeleteUsersTime, configuration.DefaultDeleteUsersTime)
	cfg.Set(configuration.UsersOverdueTimeInSeconds, configuration.DefaultUsersOverdueTimeInSeconds)

	dir, err := determineBaseDir()

	if err != nil {
		return err
	}

	dirBin := filepath.Join(dir, configuration.DefaultDirBin)
	dirVar := filepath.Join(dir, configuration.DefaultDirVar)
	dirLog := filepath.Join(dirVar, configuration.DefaultDirLog)
	dirEtc := filepath.Join(dir, configuration.DefaultDirEtc)
	dirConfig := filepath.Join(dirEtc, configuration.DefaultDirConfig)

	cfg.SetDefault(configuration.DirApp, dir)
	cfg.SetDefault(configuration.DirBin, dirBin)
	cfg.SetDefault(configuration.DirEtc, dirEtc)
	cfg.SetDefault(configuration.DirVar, dirVar)
	cfg.SetDefault(configuration.DirLog, dirLog)
	cfg.SetDefault(configuration.DirConfig, dirConfig)

	return nil
}

// ReadEnv читает конфигурации из ENV
func ReadEnv(cfg *viper.Viper) error {
	if db := os.Getenv(configuration.FullEnvPrefix + configuration.DatabaseConn); db != "" {
		_ = os.Setenv(strings.ToUpper(configuration.FullEnvPrefix+configuration.DatabaseConn), db)
	}

	cfg.SetEnvPrefix(configuration.EnvPrefix)
	cfg.AutomaticEnv()

	return nil
}

// ReadFile читает конфигурации из конфигурационного файла
func ReadFile(cfg *viper.Viper) error {
	dir := cfg.GetString(configuration.DirConfig)
	file := cfg.GetString(configuration.FileConfigName)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(dir, file+`.toml`)

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err = ioutil.WriteFile(path, []byte{}, 0755); err != nil {
				return err
			}
		}
	}

	cfg.AddConfigPath(dir)
	cfg.SetConfigName(file)
	cfg.SetConfigType(`toml`)

	return cfg.ReadInConfig()
}
