/* 
* Generated by
* 
*      _____ _          __  __      _     _
*     / ____| |        / _|/ _|    | |   | |
*    | (___ | | ____ _| |_| |_ ___ | | __| | ___ _ __
*     \___ \| |/ / _` |  _|  _/ _ \| |/ _` |/ _ \ '__|
*     ____) |   < (_| | | | || (_) | | (_| |  __/ |
*    |_____/|_|\_\__,_|_| |_| \___/|_|\__,_|\___|_|
*
* The code generator that works in many programming languages
*
*			https://www.skaffolder.com
*
*
* You can generate the code from the command-line
*       https://npmjs.com/package/skaffolder-cli
*
*       npm install -g skaffodler-cli
*
*   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *
*
* To remove this comment please upgrade your plan here: 
*      https://app.skaffolder.com/#!/upgrade
*
* Or get up to 70% discount sharing your unique link:
*       https://app.skaffolder.com/#!/register?friend=5d525f5a7b158e50f28eb4af
*
* You will get 10% discount for each one of your friends
* 
*/
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/viper"
)

type Constants struct {
	PORT   string
	PUBLIC string
	SECRET string
	Mongo  struct {
		URL    string
		DBName string
	}
}

type Config struct {
	Constants
	Database *mgo.Database
}

// NewConfig is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	config.Constants = constants
	if err != nil {
		return &config, err
	}
	
	// Connection to MangoDB database
	dbSession, err := mgo.Dial(config.Constants.Mongo.URL)
	if err != nil {
		log.Panicln("Cannot connect to MangoDB", err)
		return &config, err
	}

	config.Database = dbSession.DB(config.Constants.Mongo.DBName)

	// Create Admin user
	count, err := config.Database.C("User").Find(nil).Count()
	if count == 0 {
		println("User 'admin' created")
		type User struct {
			ID       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
			Password string        `json:"password"`
			Username string        `json:"username"`
			Roles    []string      `json:"roles"`
		}

		user := User{
			Username: "admin",
			Password: "62f264d7ad826f02a8af714c0a54b197935b717656b80461686d450f7b3abde4c553541515de2052b9af70f710f0cd8a1a2d3f4d60aa72608d71a63a9a93c0f5",
			Roles:    []string{"ADMIN"},
		}

		config.Database.C("User").Insert(user)
	}

	return &config, err
}

func initViper() (Constants, error) {
	
	if os.Getenv("ENVIRONMENT") == "DOCKER" {
		viper.SetConfigName("MyUber_docker.config") // Configuration fileName without the .TOML or .YAML extension
	} else {
		viper.SetConfigName("MyUber.config") // Configuration fileName without the .TOML or .YAML extension
	}

	viper.AddConfigPath(".")                // Search the root directory for the configuration file
	err := viper.ReadInConfig()             // Find and read the config file
	if err != nil {                         // Handle errors reading the config file
		return Constants{}, err
	}
	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("public", "public")
	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}