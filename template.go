package main

const tpl = `#### ClamAV
{{- with .Results }}
| Infected      | Result      | Engine      | Updated      |
|:-------------:|:-----------:|:-----------:|:------------:|
| {{.Infected}} | {{.Result}} | {{.Engine}} | {{.Updated}} |
{{ end -}}
`

// func printMarkDownTable(clamav ClamAV, toString bool) string {
// 	table := clitable.New([]string{"Infected", "Result", "Engine", "Updated"})
// 	table.AddRow(map[string]interface{}{
// 		"Infected": clamav.Results.Infected,
// 		"Result":   clamav.Results.Result,
// 		"Engine":   clamav.Results.Engine,
// 		// "Known":    clamav.Results.Known,
// 		"Updated": clamav.Results.Updated,
// 	})
// 	table.Markdown = true
//
// 	if toString {
// 		return table.String("ClamAV")
// 	}
// 	fmt.Println("#### ClamAV")
// 	table.Print()
// 	return ""
// }
