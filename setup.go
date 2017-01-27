package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

func makeConfigDirectory(configDirectory string) {

}

func setupConfig(auth_token string) error {
	file, file_err := getConfigFilePointer(os.O_TRUNC)
	if file_err != nil {
		return file_err
	}
	file.WriteString("auth_token=\"" + auth_token + "\"\n")
	file.Close()
	return nil
}

func saveInboxId(inbox_id int64) {
	file, file_err := getConfigFilePointer(os.O_APPEND | os.O_WRONLY)
	if file_err != nil {
		fmt.Println("Error while opening config file:")
		fmt.Println(file_err)
		return
	}
	file.WriteString("inbox_id=\"" + strconv.Itoa(int(inbox_id)) + "\"\n")
	file.Close()
}

func getConfigFilePointer(file_mode int) (*os.File, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	config_dir_path := filepath.Join(usr.HomeDir, ".config", "todoline")
	os.MkdirAll(config_dir_path, os.ModePerm)
	config_path := filepath.Join(config_dir_path, "todoline.toml")
	var file *os.File
	var file_err error
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		file, file_err = os.Create(config_path)
	} else {
		file, file_err = os.OpenFile(config_path, file_mode, os.ModePerm)
	}
	//defer file.Close()
	return file, file_err
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
