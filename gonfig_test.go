package gonfig

import (
	"embed"
	"testing"
)

//go:embed *.json
var confJsonDir embed.FS

//go:embed *.yaml
var confYamlDir embed.FS

type AppConf struct {
	AppName string `yaml:"app-name" json:"app-name"`
	DB      DBConf `yaml:"db" json:"db"`
}

type DBConf struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

func TestReadJsonConfig(t *testing.T) {
	var appConf AppConf
	err := Unmarshal(confJsonDir, &appConf, UnmarshalWith(Json))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("%v", appConf)
}

func TestReadYamlConfig(t *testing.T) {
	var appConf AppConf
	err := Unmarshal(confYamlDir, &appConf, UnmarshalWith(Yaml))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("%v", appConf)
}

func TestProfileUseEnv(t *testing.T) {
	var appConf AppConf
	err := Unmarshal(confYamlDir, &appConf, ProfileUseEnv("env", "dev"))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("%v", appConf)
}

func TestActiveFunc(t *testing.T) {
	var appConf AppConf
	err := Unmarshal(confYamlDir, &appConf, FilePrefix("config"))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("%v", appConf)
}

func TestProfileFunc(t *testing.T) {
	var appConf AppConf
	err := Unmarshal(confYamlDir, &appConf, ProfileFunc(func() string {
		return "dev"
	}))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("%v", appConf)
}
