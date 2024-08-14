package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

type OutputFormat string

const (
	OutputFormatJSON  OutputFormat = "json"
	OutputFormatYAML  OutputFormat = "yaml"
	OutputFormatTable OutputFormat = "table"
)

func Print(data interface{}, format OutputFormat) error {
	switch format {
	case OutputFormatJSON:
		return printJSON(data)
	case OutputFormatYAML:
		return printYAML(data)
	case OutputFormatTable:
		return printTable(data)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}

func printJSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func printYAML(data interface{}) error {
	return yaml.NewEncoder(os.Stdout).Encode(data)
}

func printTable(data interface{}) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(getHeaders(data))
	table.AppendBulk(getRows(data))
	table.Render()
	return nil
}

func getHeaders(data interface{}) []string {
	// Implement this function to return the appropriate headers based on the data type
	return []string{"ID", "Name", "Status"}
}

func getRows(data interface{}) [][]string {
	// Implement this function to return the appropriate rows based on the data type
	return [][]string{
		{"1", "Example", "Active"},
		{"2", "Test", "Inactive"},
	}
}