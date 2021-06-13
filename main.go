package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os/user"
	"strings"
)

var activeUser, err = user.Current()
var userHome = activeUser.HomeDir
var awsConf = userHome + "/.aws/config"
var confFile = userHome + "/.awswitch"

func main() {

	color.Set(color.FgCyan)
	fmt.Println("Active Profile :", getActiveProfile())
	color.Unset()

	profiles := GetProfiles()
	prompt := promptui.Select{
		Label:    "Choose the AWS Profile to switch to",
		Items:    profiles,
		HideHelp: true,
	}

	_, result, prmt_err := prompt.Run()

	if prmt_err != nil {
		fmt.Printf("Prompt failed %v\n", prmt_err)
		return
	}

	var choosenProfile = result

	if result != "default" {
		choosenProfile = strings.Split(result, " ")[1]
	}
	WriteConfFile(choosenProfile)
	println("Switched AWS profile to", choosenProfile)

}

func getActiveProfile() string {
	// read the whole file at once
	activeProfile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(string(activeProfile))
}

func GetProfiles() []string {
	cfg, aws_err := ini.Load(awsConf)

	if aws_err != nil {
		fmt.Printf("Error loading AWS CLI configuration %v\n", aws_err)
	}

	raw_profiles := cfg.SectionStrings()
	i := 0 // Index to delete

	copy(raw_profiles[i:], raw_profiles[i+1:])        // Shift a[i+1:] left one index.
	raw_profiles[len(raw_profiles)-1] = ""            // Erase last element (write zero value).
	raw_profiles = raw_profiles[:len(raw_profiles)-1] // Truncate slice.

	return raw_profiles
}

func WriteConfFile(profile string) {
	// write the whole body at once
	err = ioutil.WriteFile(confFile, []byte(profile), 0644)
	if err != nil {
		panic(err)
	}
}
