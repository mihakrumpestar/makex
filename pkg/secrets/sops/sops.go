package sops

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

type Sops struct {
	filePath string
}

func (so *Sops) Export() error {
	envContent, err := exec.Command("sops", "-d", so.filePath).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error in decryptAndExport %s: %s", envContent, err)
	}

	envLines := strings.Split(string(envContent), "\n")
	for _, env := range envLines {
		envParts := strings.SplitN(env, "=", 2)
		if len(envParts) == 2 {
			os.Setenv(envParts[0], envParts[1])
		}
	}

	log.Info().Msgf("Loaded %s\n", so.filePath)

	return nil
}

func (so *Sops) Save(prefix string) error {
	encryptedFile := so.filePath

	decryptedDir := os.Getenv("DECRYPTED_DIR")
	if len(decryptedDir) == 0 {
		return fmt.Errorf("Variable 'DECRYPTED_DIR' must be set for Sops.Save()")
	}

	decryptedFile := filepath.Join(decryptedDir, filepath.Dir(encryptedFile))
	decryptedFile = strings.ReplaceAll(decryptedFile, prefix, "")
	if err := os.MkdirAll(filepath.Dir(decryptedFile), 0755); err != nil {
		return err
	}

	content, err := exec.Command("sops", "-d", encryptedFile).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error in decryptAndSave for file '%s' %s: %s", encryptedFile, content, err)
	}

	if err := os.WriteFile(decryptedFile, content, 0755); err != nil {
		return err
	}

	log.Info().Msgf("Decrypted %s to %s\n", encryptedFile, decryptedFile)
	return nil
}
