package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

type Config struct {
	ApiToken  string
	ApiDomain string
}

func SaveToken(token string) {
	write_to_config_file(api_token_path(), token)
}

func SaveDomain(domain string) {
	write_to_config_file(api_domain_path(), domain)
}

func Load() (*Config, error) {
	api_token, token_err := ioutil.ReadFile(api_token_path())

	if token_err != nil {
		return nil, token_err
	}

	api_domain, domain_err := ioutil.ReadFile(api_domain_path())

	if domain_err != nil {
		api_domain = []byte("https://semaphoreci.com")
	}

	config := &Config{
		ApiToken:  strings.TrimSpace(string(api_token)),
		ApiDomain: strings.TrimSpace(string(api_domain)),
	}

	return config, nil
}

func config_folder_path() string {
	user, err := user.Current()

	if err != nil {
		fmt.Println("Fatal error: Can't determine current user")
		os.Exit(1)
	}

	return fmt.Sprintf("%s/.semaphore", user.HomeDir)
}

func api_token_path() string {
	return fmt.Sprintf("%s/api_token", config_folder_path())
}

func api_domain_path() string {
	return fmt.Sprintf("%s/api_domain", config_folder_path())
}

func write_to_config_file(path, content string) {
	err := os.MkdirAll(config_folder_path(), 0777)

	if err != nil {
		fmt.Println("Error: Can't create ~/.semaphore directory.")
		os.Exit(1)
	}

	f, createErr := os.Create(path)

	if createErr != nil {
		fmt.Printf("Error: Can't create %s file.\n", path)
		fmt.Print(createErr)
		os.Exit(1)
	}

	defer f.Close()

	_, writeErr := f.WriteString(content)

	if writeErr != nil {
		fmt.Printf("Error: Can't write to %s file.\n", path)
		os.Exit(1)
	}

	f.Sync()
}
