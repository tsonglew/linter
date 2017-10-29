package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func check(code, linter, filePath, outPath string) (result []string, err error) {
	err = ioutil.WriteFile(filePath, []byte(code), 0666)
	if err != nil {
		logrus.Errorf("write file error: %v", err)
	}

	cmd = exec.Command(linter, filePath)
	logrus.Infof("%v", cmd)
	outfile, err := os.Create(outPath)
	if err != nil {
		logrus.Errorf("create out file error: %v", err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile
	if err := cmd.Start(); err != nil {
		logrus.Errorf("running command error: %v", err)
	}
	cmd.Wait()
	result, err = readLines(strings.Join([]string{tmpPre, "out.txt"}, ""))
	return
}
