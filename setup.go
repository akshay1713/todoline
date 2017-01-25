package main

import (
	"fmt"
	//"github.com/akshaysingh1713/gotodoist"
	"github.com/pelletier/go-toml"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func makeConfigDirectory(configDirectory string) {

}

func setupConfig(auth_token string) error {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(usr.HomeDir)
	config_dir_path := filepath.Join(usr.HomeDir, ".config", "todoline")
	os.MkdirAll(config_dir_path, os.ModePerm)
	config_path := filepath.Join(config_dir_path, "todoline.toml")
	var file *os.File
	var file_error error
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		file, file_error = os.Create(config_path)
	} else {
		file, file_error = os.Open(config_path)
	}
	if file_error != nil {
		return file_error
	}
	file.WriteString("auth_token=\"" + auth_token + "\"")
	return nil
}

func getConfig() *toml.TomlTree {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	config_dir_path := filepath.Join(usr.HomeDir, ".config", "todoline")
	config_path := filepath.Join(config_dir_path, "todoline.toml")
	config, err := toml.LoadFile(config_path)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
	return config
}
