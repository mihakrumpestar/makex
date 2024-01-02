package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func RecursiveGlob(dir string, ext string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ext) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// Get the current working directory
func BaseDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get the current working directory %s", err)
	}

	return currentDir, nil
}

func FindFilesFromBaseDir(subfolder string, prefixes []string, baseFileName string, extensions []string, recursive bool) ([]string, error) {
	baseDir, err := BaseDir()
	if err != nil {
		return nil, err
	}

	// The slice to store matching file paths
	var files []string

	// Define the function to process each file
	var walkFn filepath.WalkFunc = func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err // Propagate the error
		}

		// Check for recursion flag and if it's a directory
		if !recursive && file.IsDir() && path != filepath.Join(baseDir, subfolder) {
			return nil
		}

		// Process only if it's a file
		if file.IsDir() {
			return nil
		}

		filename := file.Name()
		matched := true

		// Check for file name if specified
		if baseFileName != "" && !strings.Contains(filename, baseFileName) {
			matched = false
		}

		// Check for prefixes if any are specified
		if matched && len(prefixes) > 0 {
			matched = false
			for _, prefix := range prefixes {
				if strings.HasPrefix(filename, prefix) {
					matched = true
					break
				}
			}
		}

		// Check for extensions if any are specified
		if matched && len(extensions) > 0 {
			matched = false
			for _, extension := range extensions {
				if strings.HasSuffix(filename, extension) {
					matched = true
					break
				}
			}
		}

		// If the file matches the criteria, add it to the list
		if matched {
			files = append(files, path)
		}

		return nil
	}

	// Walk the directory tree
	if recursive {
		err = filepath.Walk(filepath.Join(baseDir, subfolder), walkFn)
	} else {
		dir := filepath.Join(baseDir, subfolder)
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			fileInfo, err := entry.Info()
			if err != nil {
				return nil, err
			}
			err = walkFn(filepath.Join(dir, entry.Name()), fileInfo, nil)
			if err != nil {
				return nil, err
			}
		}
	}

	log.Info().Strs("configFiles", files).Msg("")

	return files, err
}
