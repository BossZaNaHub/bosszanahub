package main

import (
	"fmt"
	"github.com/bosszanahub/bzlib/bzgorm"
	"github.com/bosszanahub/bzlib/bzmongo"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type File string

const (
	YAML File = "yaml"
	JSON File = "json"
	ENV File = ".env"
)
const (
	username = "username"
	password = "password"
)

var (
	user = Get(username, "")
	pass = Get(password, "")
)

func Get(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadConfig(file File, directory string) error {
	viper.SetConfigName("config")
	if directory == "" {
		directory = "$pwd"
	}
	viper.AddConfigPath(directory)
	viper.AddConfigPath("$pwd/" + directory)
	viper.AddConfigPath(".")
	fileType := fileType(file)
	viper.SetConfigType(fileType)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	//keys := viper.AllSettings()
	//
	//f(func(key, defaultValue string) string {
	//	return LoadEnvConfigFromMap(keys, key, defaultValue)
	//})

	return nil
}

func fileType(fileType File) string {

	switch fileType {
	case JSON:
		return "json"
	case YAML:
		return "yaml"
	case ENV:
		return ".env"
	default:
		return ""
	}

}

type User struct {
	Name string `json:"name,omitempty"`
	BirthDate time.Time `json:"birth_of_date,omitempty"`
}

func LoadEnvConfigFromMap(keys map[string]interface{}, key, defaultValue string) string {
	key = strings.ToLower(key)
	if _, ok := keys[key]; ok {
		return viper.GetString(key)
	}

	return defaultValue
}

func main() {
	a := LoadConfig(YAML,"bzenv")
	fmt.Println(a)

	client := bzmongo.NewMongo("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb", "shibe", 10)
	err := client.OpenMongoConnection()
	if err != nil {
		panic(err)
	}
	c, err := client.SelectCollection("users", User{Name: "Test"})
	if err != nil {
		panic(err)
	}
	fmt.Println(c)

	_, err = bzgorm.NewClient(bzgorm.Config{Username: "boss", Password: "boss", DB: "boss", Host: "localhost"}).InitMysql()
	if err != nil {
		panic(err)
	}


	defer client.CloseConnection()
}