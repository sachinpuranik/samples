package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var (
	textQry = `
		{{- if .PrintIt -}}
		\n Hello {{.Name}} , printIt is true 
		\n len is {{len .Numbers}} 
		{{- end -}}
		{{- if .FindNumberInList 2 -}}
		\n WOW found number 1
		{{- end -}}
		`
	processedQry = template.Must(
		template.New("sql").Parse(textQry))
)

type paramObject struct {
	PrintIt bool
	Name    string
	Numbers []int
}

func (p *paramObject) FindNumberInList(n int) bool {
	for _, num := range p.Numbers {
		if num == n {
			return true
		}
	}
	return false
}

func main() {
	param := &paramObject{}
	param.PrintIt = true
	param.Numbers = append(param.Numbers, 1)
	param.Numbers = append(param.Numbers, 3)
	param.Numbers = append(param.Numbers, 5)
	param.Numbers = append(param.Numbers, 7)
	param.Name = "Sachin"

	var b bytes.Buffer
	err := processedQry.Execute(&b, param)
	if err != nil {
		panic(err)
	}
	query := b.String()

	res := strings.Split(query, "\n")
	for _, val := range res {
		fmt.Println(val)
	}
}
