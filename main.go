package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	initConfig()
	test()
}

func test() {
	giteaBaseURL := viper.Get("giteaBaseURL").(string)
	giteaRepoOwner := viper.Get("giteaRepoOwner").(string)
	giteaRepo := viper.Get("giteaRepo").(string)
	giteaRepoBranch := viper.Get("giteaRepoBranch").(string)
	giteaToken := viper.Get("giteaToken").(string)

	method := "GET"
	url := giteaBaseURL + "/api/v1/repos/" + giteaRepoOwner + "/" + giteaRepo + "/commits?sha=" + giteaRepoBranch
	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "token "+giteaToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(body)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}
