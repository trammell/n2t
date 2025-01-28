package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// If `src` is a directory, return all the .vm files in the directory,
// otherwise return a single .vm filename.
func getSourceFiles(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return []string{}, err
	}

	// if it's a directory, match out the .vm files
	if info.IsDir() {
		log.Printf(`source file "%s" is a directory`, path)
		f, err := os.Open(path)
		if err != nil {
			return []string{}, err
		}
		files, err := f.Readdir(0)
		if err != nil {
			return []string{}, err
		}
		var vmfiles []string
		for _, v := range files {
			if v.Mode().IsRegular() && strings.HasSuffix(v.Name(), ".vm") {
				vmfiles = append(vmfiles, v.Name())
			}
		}
		if len(vmfiles) == 0 {
			err := fmt.Errorf(`no .vm files found in directory "%s"`, path)
			return []string{}, err
		}
		return vmfiles, nil
	}

	// `path` is a regular file, but is it a VM file?
	if strings.HasSuffix(path, ".vm") {
		return []string{path}, nil
	}

	// give up
	err = fmt.Errorf(`"%s" is the wrong type, should be .vm`, path)
	return []string{}, err
}

// Construct the name of the destination .asm file. If the source path is a
// directory, create the .asm file in the directory, e.g.
// `/path/to/project/project.asm`.
func getDestFile(path string) (string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	// it's a directory
	if stat.IsDir() {
		log.Printf(`"%s" is a directory`, path)
		dest := filepath.Join(path, filepath.Base(path)+".asm")
		return dest, nil
	}
	// it's a file
	log.Printf(`"%s" is a regular file`, path)
	if strings.HasSuffix(path, ".vm") {
		dest := regexp.MustCompile(`.vm$`).ReplaceAllString(path, ".asm")
		return dest, nil
	}
	return "", fmt.Errorf(`"%s" is the wrong type, should be .vm`, path)
}
