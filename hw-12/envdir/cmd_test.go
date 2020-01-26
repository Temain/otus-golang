package envdir

import (
	"testing"
)

func TestRunCmd(t *testing.T) {
	cmd := []string{"env"}
	env := map[string]string{
		"USER_NAME": "User",
	}
	code := RunCmd(cmd, env)
	if code != 0 {
		t.Fatal("bad exit code, expected 0")
	}
}

func TestRunWrongCmd(t *testing.T) {
	cmd := []string{"env", "-t"}
	env := map[string]string{
		"USER_NAME": "User",
	}
	code := RunCmd(cmd, env)
	if code == 0 {
		t.Fatal("bad exit code, expected not 0")
	}
}
