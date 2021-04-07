package main

import (
	"./github"
	"html/template"
	"log"
	"os"
	"time"
)

// {{}}为一个action，获取IssuesSearchResult中的参数
// {{range .Items}}和{{end}}组成循环
const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------- 
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

// 运行方法：./text repo:golang/go is:open json decoder
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}