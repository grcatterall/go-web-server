package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Column struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	PrimaryKey bool   `json:"primary_key,omitempty"`
	Unique     bool   `json:"unique,omitempty"`
}

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

type Schema struct {
	Tables []Table `json:"tables"`
}

func main() {
	newSchema, err := loadSchema("schemas/schema.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Found schema.json")

	currentSchema, err := loadSchema("schemas/applied-schema.json")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No current schema found, creating new one")
		} else {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Found applied-schema.json")

	changes := diffSchemas(currentSchema, newSchema)

	if len(changes) <= 0 {
		fmt.Printf("No schema changes detected - Ending migration process")
		return
	}

	changes = append([]string{"-- THIS DATA HAS BEEN AUTO GENERATED \n"}, changes...)

	fmt.Printf("SQL Generated - Writing to file")
	fmt.Println()

	t := time.Now().Format("20060102150405")
	filepath := fmt.Sprintf("migrations/%s.sql", t)

	err = os.WriteFile(filepath, []byte(strings.Join(changes, "\n\n")), 0644)

	if err != nil {
		fmt.Println("Unable to write migration to file")
		return
	}

	fmt.Printf("SQL successfully written to '%s'", filepath)
	fmt.Println()

	schema, err := os.ReadFile("schemas/schema.json")

	fmt.Printf("Storing new schema")
	fmt.Println()

	if err != nil {
		fmt.Println("Unable to write migration to file")
		return
	}

	err = os.WriteFile("schemas/applied-schema.json", []byte(schema), 0644)
	if err != nil {
		fmt.Println("Unable to write to applied schema")
		return
	}

	fmt.Printf("Schema stored")
	fmt.Println()
	fmt.Printf("Migration Successful")
	fmt.Println()
}

func loadSchema(filename string) (*Schema, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var schema Schema
	err = json.Unmarshal(data, &schema)

	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func diffSchemas(oldSchema, newSchema *Schema) []string {
	var changes []string

	if len(oldSchema.Tables) == 0 {
		fmt.Printf("No existing schema found - Generating new schema")
		fmt.Println()
		for _, newTable := range newSchema.Tables {
			createString := createNewTable(newTable)
			changes = append(changes, createString)
		}

		return changes
	}

	changes = append(changes, detectNewAndModifiedTables(oldSchema, newSchema)...)
	changes = append(changes, detectDeletedTables(oldSchema, newSchema)...)

	return changes
}

func detectNewAndModifiedTables(oldSchema, newSchema *Schema) []string {
	var changes []string
	for _, newTable := range newSchema.Tables {
		tableFound := false
		for _, oldTable := range oldSchema.Tables {
			if newTable.Name == oldTable.Name {
				tableFound = true
				changes = append(changes, detectModifiedColumns(oldTable, newTable)...)
				break
			}
		}
		if !tableFound {
			fmt.Printf("New table '%s' found\n", newTable.Name)
			changes = append(changes, createNewTable(newTable))
		}
	}
	return changes
}

func detectModifiedColumns(oldTable, newTable Table) []string {
	var changes []string
	for _, newCol := range newTable.Columns {
		colFound := false
		for _, oldCol := range oldTable.Columns {
			if newCol.Name == oldCol.Name {
				colFound = true
				if newCol.Type != oldCol.Type {
					newType, err := convertType(newCol.Type)
					if err != nil {
						fmt.Println(err)
						break
					}
					changes = append(changes, fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s;", newTable.Name, newCol.Name, newType))
				}
				break
			}
		}
		if !colFound {
			changes = append(changes, fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", newTable.Name, newCol.Name, newCol.Type))
		}
	}
	return changes
}

func detectDeletedTables(oldSchema, newSchema *Schema) []string {
	var changes []string
	for _, oldTable := range oldSchema.Tables {
		tableNameFound := false
		for _, newTable := range newSchema.Tables {
			if oldTable.Name == newTable.Name {
				tableNameFound = true
				break
			}
		}
		if !tableNameFound {
			fmt.Printf("Table %s no longer present, removing\n", oldTable.Name)
			changes = append(changes, fmt.Sprintf("DROP TABLE %s;", oldTable.Name))
		}
	}
	return changes
}


func convertType(colType string) (string, error) {
	switch colType {
	case "uuid":
		return colType, nil
	case "int":
		return colType, nil
	case "float":
		return "float8", nil
	case "string":
		return "text", nil
	default:
		return "", fmt.Errorf("invalid column type - %s", colType)
	}
}

func createNewTable(table Table) string {
	fmt.Printf("Creating new table %s", table.Name)
	fmt.Println()

	var columns []string

	for _, col := range table.Columns {
		colType, err := convertType(col.Type)

		if err != nil {
			fmt.Println(err)
		}

		var colString string = fmt.Sprintf("%s %s", col.Name, colType)

		if col.PrimaryKey || col.Unique {
			appendString := "PRIMARY KEY"
			if !col.PrimaryKey {
				appendString = "UNIQUE"
			}

			colString = fmt.Sprintf("%s %s %s", col.Name, colType, appendString)
		}

		columns = append(columns, colString)
	}

	return fmt.Sprintf("CREATE TABLE \"%s\" (\n\t%s\n);\n", table.Name, strings.Join(columns, ",\n\t"))
}
