package main

import (
	"html/template"
	"net/http"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

var cmd *exec.Cmd
var tmpPre = "./tmp/"
var staticPre = "./static/"
var outPath = "./tmp/out.txt"
var result []string
var filePath string
var linter string
var err error

func lintHandler(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles(strings.Join([]string{staticPre, "index.html"}, ""))

	if request.Method == "GET" {
		logrus.Infof("GET /lint/")
		data := struct {
			Code     string
			Comments []string
		}{
			"Input Your Code Here...",
			[]string{"Comments will be displayed here..."},
		}
		if err = t.Execute(writer, data); err != nil {
			logrus.Errorf("Getting Index Error: %v", err)
		}
	} else if request.Method == "POST" {
		logrus.Infof("POST /lint/ %v", request.Form)
		request.ParseForm()
		lang := strings.Join(request.Form["lang"], "")
		code := strings.Join(request.Form["code"], "")

		switch lang {
		case "Python":
			filePath = strings.Join([]string{tmpPre, "test.py"}, "")
			linter = "pylint"
		case "Cpp":
			filePath = strings.Join([]string{tmpPre, "test.cpp"}, "")
			linter = "cpplint"
		default:
			if err := t.Execute(writer, []string{"Choose in \"Python\" and \"CPP\""}); err != nil {
				logrus.Errorf("Error while showing error: %v", err)
			}
			return
		}
		if result, err = check(code, linter, filePath, outPath); err != nil {
			logrus.Errorf("linting error: %v", err)
			return
		}
		data := struct {
			Code     string
			Comments []string
		}{
			code,
			result,
		}
		if err = t.Execute(writer, data); err != nil {
			logrus.Errorf("Rendering Error: %v", err)
		}
	}
}
