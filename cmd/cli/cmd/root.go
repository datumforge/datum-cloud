// Package datumcloud is our cobra cli implementation
package datumcloud

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	appName = "datum-cloud"
)

var (
	cfgFile      string
	OutputFormat string
	Logger       *zap.SugaredLogger
	Config       *koanf.Koanf
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   appName,
	Short: "the datum-cloud cli",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfiguration(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	Config = koanf.New(".") // Create a new koanf instance.

	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/."+appName+".yaml)")
	RootCmd.PersistentFlags().String("host", "", "datum cloud api url, if not set the config default will be used")
	RootCmd.PersistentFlags().String("datum-host", "", "datum api url, if not set the config default will be used")

	// Auth flags
	RootCmd.Flags().String("token", "", "datum api token")

	// Logging flags
	RootCmd.PersistentFlags().Bool("debug", false, "enable debug logging")
	RootCmd.PersistentFlags().Bool("pretty", false, "enable pretty (human readable) logging output")

	// Output flags
	RootCmd.PersistentFlags().StringVarP(&OutputFormat, "format", "z", "table", "output format (json, table)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// load the flags to ensure we know the correct config file path
	initConfiguration(RootCmd)

	// load the config file and env vars
	loadConfigFile()

	// reload because flags and env vars take precedence over file
	initConfiguration(RootCmd)

	// setup logging configuration
	setupLogging()
}

// setupLogging configures the logger based on the command flags
func setupLogging() {
	cfg := zap.NewProductionConfig()
	if Config.Bool("pretty") {
		cfg = zap.NewDevelopmentConfig()
	}

	if Config.Bool("debug") {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	l, err := cfg.Build()
	cobra.CheckErr(err)

	Logger = l.Sugar().With("app", appName)
	defer Logger.Sync() //nolint:errcheck
}

// initConfiguration loads the configuration from the command flags of the given cobra command
func initConfiguration(cmd *cobra.Command) {
	loadEnvVars()

	loadFlags(cmd)
}

func loadConfigFile() {
	if cfgFile == "" {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		cfgFile = filepath.Join(home, "."+appName+".yaml")
	}

	// If the config file does not exist, do nothing
	if _, err := os.Stat(cfgFile); errors.Is(err, os.ErrNotExist) {
		return
	}

	err := Config.Load(file.Provider(cfgFile), yaml.Parser())

	cobra.CheckErr(err)
}

func loadEnvVars() {
	err := Config.Load(env.ProviderWithValue("DATUMCLOUD_", ".", func(s string, v string) (string, interface{}) {
		key := strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, "DATUMCLOUD_")), "_", ".")

		if strings.Contains(v, ",") {
			return key, strings.Split(v, ",")
		}

		return key, v
	}), nil)

	cobra.CheckErr(err)
}

func loadFlags(cmd *cobra.Command) {
	err := Config.Load(posflag.Provider(cmd.Flags(), Config.Delim(), Config), nil)

	cobra.CheckErr(err)
}
