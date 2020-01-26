package envdir

import (
	"io/ioutil"
	"testing"
)

func TestReadDir(t *testing.T) {
	err := prepareFile("data/CITY", "Krasnodar")
	if err != nil {
		t.Fatal("error on prepare env file")
	}
	err = prepareFile("data/USER_NAME", "Artem")
	if err != nil {
		t.Fatal("error on prepare env file")
	}

	env, err := ReadDir("./data")
	if err != nil {
		t.Fatal("error on prepare env file")
	}

	if env["CITY"] == "" {
		t.Fatal("bad result, expected CITY=Krasnodar")
	}
	if env["USER_NAME"] == "" {
		t.Fatal("bad result, expected USER_NAME=Artem")
	}
}

func TestReadDirWithEmptyFile(t *testing.T) {
	dir := "data/"
	envName := "USER_AGE"
	err := prepareFile(dir+envName, "")
	if err != nil {
		t.Fatal("error on prepare env file")
	}
	err = prepareFile(dir+envName, "")

	env, err := ReadDir(dir)
	if err != nil {
		t.Fatal("error on prepare env file")
	}

	if env[envName] != "" {
		t.Fatalf("bad result, expected %v=", envName)
	}
}

func prepareFile(name string, data string) error {
	fileBytes := []byte(data)
	err := ioutil.WriteFile(name, fileBytes, 0644)
	return err
}
