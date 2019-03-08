package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/theskyinflames/cmdarchetype/config"

	"github.com/davecgh/go-spew/spew"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cfg     *config.Config
	log     *logrus.Logger

	rootCmd = &cobra.Command{
		Use:   "cmdarchetype",
		Short: "cmdarchetype is an archetype to build command line tools based on Cobra+Viper",
		Long:  "cmdarchetype is an archetype to build command line tools based on Cobra+Viper",
		Run:   startCmd,
	}
)

func init() {

	var (
		sourceData      string
		doAsync         bool
		resultReceivers []string
		user            string
		password        string
		dbURL           string
	)

	cfg = &config.Config{}

	cobra.OnInitialize(readConfigFromFile)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Path to config file (with extension)")
	rootCmd.Flags().BoolVarP(&doAsync, "do-async", "a", false, "Do it asynchronously?")
	rootCmd.Flags().StringVarP(&sourceData, "source-data", "s", "", "Source data")
	rootCmd.Flags().StringArrayVarP(&resultReceivers, "result-receivers", "r", nil, "List of receivers for the action result")
	rootCmd.Flags().StringVarP(&user, "user", "u", "", "db user to be connected to db")
	rootCmd.Flags().StringVarP(&password, "password", "p", "", "db password to be connected to db")
	rootCmd.Flags().StringVarP(&dbURL, "db-url", "b", "", "source data db URL")

	viper.BindPFlag("do-async", rootCmd.Flags().Lookup("do-async"))
	viper.BindPFlag("source-data", rootCmd.Flags().Lookup("source-data"))
	viper.BindPFlag("result-receivers", rootCmd.Flags().Lookup("result-receivers"))
	viper.BindPFlag("db-connection-params.user", rootCmd.Flags().Lookup("user"))
	viper.BindPFlag("db-connection-params.password", rootCmd.Flags().Lookup("password"))
	viper.BindPFlag("db-connection-params.db-url", rootCmd.Flags().Lookup("db-url"))
}

func Execute() {

	// Set logging service
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	log = logrus.New()

	// Start the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// If a config file has been specified, it will be used
func readConfigFromFile() {

	if cfgFile != "" {

		// check for file exist
		_, err := os.Stat(cfgFile)
		if err != nil {
			panic(err)
		}

		log.Infof("loading config from [%s] file", cfgFile)

		// get the filepath
		abs, err := filepath.Abs(cfgFile)
		if err != nil {
			panic(err)
		}
		base := filepath.Base(abs)
		path := filepath.Dir(abs)

		viper.SetConfigFile(abs)
		viper.SetConfigName(strings.Split(base, ".")[0])
		viper.AddConfigPath(path)

		// Find and read the config file; Handle errors reading the config file
		if err = viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}
}

func startCmd(cmd *cobra.Command, args []string) {

	// Load configuration
	cfg := &config.Config{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}

	// validate configuration
	err = cfg.Validate()
	if err != nil {
		log.Errorf("some went wrong when validating the given flags: %s\n\n", err.Error())
		cmd.HelpFunc()(cmd, args)
		os.Exit(1)
	}

	// Start the migration
	log.Infof("loaded config: %s\n", spew.Sdump(cfg))
	ts := time.Now()
	log.Infof("starting the command at %s", fmt.Sprint(ts))
	myCommand := NewMyCommand(cfg, log)
	err = myCommand.DoAction()
	if err != nil {
		log.Errorf("some when wrong went trying to do the action: %s\n", err.Error())
		os.Exit(1)
	}
	log.Infof("Action done sucessfully !!!, in %s\n", time.Now().Sub(ts))
	os.Exit(0)
}
