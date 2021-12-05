package main

import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.Totalcount}} issues</h1>
`))
