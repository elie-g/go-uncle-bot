package main

import (
    "github.com/DrunkenPoney/go-uncle-bot/bot"
    . "github.com/DrunkenPoney/go-uncle-bot/utils"
    _ "github.com/joho/godotenv/autoload"
    "github.com/mitchellh/go-homedir"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "log"
    "os"
)

const (
    CONFIG_NAME    = ".uncle-bot"
    DEFAULT_PREFIX = "!"
)

var cmd = &cobra.Command{}
var cfgFile string

func main() {
    execFile, err := os.Executable()
    CheckErr(err)
    
    cmd.Use = execFile
    cmd.Short = "Uncle Bot is the uncle of all the Discord's bots"
    cmd.Run = func(cmd *cobra.Command, args []string) {
        bot.Initialize()
    }
    
    CheckErr(cmd.Execute())
}

func init() {
    cobra.OnInitialize(initConfig)
    cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is .uncle-bot.yaml)")
    cmd.PersistentFlags().StringP("token", "t", "", "Token to use for the bot")
    cmd.PersistentFlags().StringP("prefix", "p", "", "The commands prefix (default: `!`)")
    cmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output (default: false)")
    
    CheckErr(viper.BindEnv("botToken", "BOT_TOKEN"))
    CheckErr(viper.BindEnv("botPrefix", "BOT_PREFIX"))
    CheckErr(viper.BindEnv("verbose", "BOT_DEBUG"))
    
    CheckErr(viper.BindPFlag("botToken", cmd.PersistentFlags().Lookup("token")))
    CheckErr(viper.BindPFlag("botPrefix", cmd.PersistentFlags().Lookup("prefix")))
    CheckErr(viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose")))
    
    viper.RegisterAlias("verbose", "debug")
    viper.SetDefault("botPrefix", DEFAULT_PREFIX)
    viper.SetDefault("verbose", false)
    viper.AllowEmptyEnv(true)
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        cwd, err := os.Getwd()
        CheckErr(err)
        viper.AddConfigPath(cwd)
        
        home, err := homedir.Dir()
        CheckErr(err)
        viper.AddConfigPath(home)
        
        viper.SetConfigName(CONFIG_NAME)
    }
    
    _ = viper.ReadInConfig()
    
    if viper.GetString("botToken") == "" {
        log.Fatalln("bot token not provided")
    }
}
