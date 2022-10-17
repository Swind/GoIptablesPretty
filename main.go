package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	var inChain bool = false
	var headers []string
	var rows [][]string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			if inChain {
				headers = append(headers, "extra")
				t := table.NewWriter()
				t.SetOutputMirror(os.Stdout)

				headerRow := make([]interface{}, len(headers))
				for i, v := range headers {
					headerRow[i] = v
				}
				t.AppendHeader(headerRow)

				for _, row := range rows {
					tableRow := make([]interface{}, len(row))
					for i, v := range row {
						tableRow[i] = v
					}
					t.AppendRow(tableRow)
				}
				t.SetStyle(table.StyleColoredBright)
				t.Render()

				headers = nil
				rows = nil
				inChain = false
			}
		} else if strings.HasPrefix(line, "Chain") {
			fmt.Printf("\n%s\n", line)
		} else if strings.HasPrefix(line, "pkts") || strings.HasPrefix(line, "target") {
			headers = append(headers, strings.Fields(line)...)
			inChain = true
		} else if inChain {
			parts := strings.Fields(line)
			begin := parts[:len(headers)]
			extra := strings.Join(parts[len(headers):], " ")
			rows = append(rows, append(begin, extra))
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
