package envdir

import (
	"os"
	"os/exec"
)

// RunCmd выполняет команду с переданными переменными окружения.
func RunCmd(cmd []string, env map[string]string) int {
	name := cmd[0]
	args := cmd[1:]

	var envFlat []string
	for key, value := range env {
		envFlat = append(envFlat, key+"="+value)
	}

	proc := exec.Command(name, args...)
	proc.Env = append(os.Environ(), envFlat...)
	proc.Stdout = os.Stdout
	proc.Stdin = os.Stdin
	proc.Stderr = os.Stderr
	if err := proc.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}

	return 0
}
