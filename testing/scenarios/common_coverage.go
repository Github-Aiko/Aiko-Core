//go:build coverage
// +build coverage

package scenarios

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/Github-Aiko/Aiko-Core/common/uuid"
)

func BuildAiko() error {
	genTestBinaryPath()
	if _, err := os.Stat(testBinaryPath); err == nil {
		return nil
	}

	cmd := exec.Command("go", "test", "-tags", "coverage coveragemain", "-coverpkg", "github.com/Github-Aiko/Aiko-Core/...", "-c", "-o", testBinaryPath, GetSourcePath())
	return cmd.Run()
}

func RunAikoProtobuf(config []byte) *exec.Cmd {
	genTestBinaryPath()

	covDir := os.Getenv("Aiko_COV")
	os.MkdirAll(covDir, os.ModeDir)
	randomID := uuid.New()
	profile := randomID.String() + ".out"
	proc := exec.Command(testBinaryPath, "-config=stdin:", "-format=pb", "-test.run", "TestRunMainForCoverage", "-test.coverprofile", profile, "-test.outputdir", covDir)
	proc.Stdin = bytes.NewBuffer(config)
	proc.Stderr = os.Stderr
	proc.Stdout = os.Stdout

	return proc
}
