package envdir

import (
	"io/ioutil"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir := "data/"
	envName1 := "CITY"
	envValue1 := "Krasnodar"
	err := prepareFile(dir+envName1, envValue1)
	if err != nil {
		t.Fatal("error on prepare env file")
	}
	envName2 := "USER_NAME"
	envValue2 := "Artem"
	err = prepareFile(dir+envName2, envValue2)
	if err != nil {
		t.Fatal("error on prepare env file")
	}

	env, err := ReadDir(dir)
	if err != nil {
		t.Fatal("error on prepare env file")
	}

	if env[envName1] == "" {
		t.Fatalf("bad result, expected %v=%v", envName1, envValue1)
	}
	if env[envName2] == "" {
		t.Fatalf("bad result, expected %v=%v", envName2, envValue2)
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
