package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/volatiletech/sqlboiler/strmangle"
)

var sqlTypeMap = map[string]string{
	"object":  "text",
	"integer": "int",
	"array":   "text",
	"boolean": "bool",
	"number":  "float",
	"string":  "varchar(255)",
}

var sqlBindMap = map[string]string{
	"object":  "json",
	"integer": "int",
	"array":   "json",
	"boolean": "bool",
	"number":  "float",
	"string":  "string",
}

type API struct {
	Schemas []*Schema
	Types   map[string]*JSONSchema
}

type ColumnConfig struct {
	Path   string
	Type   string
	Bind   string
	Column string
	Name   string
}

type Schema struct {
	FileName    string
	ID          string                 `yaml:"id"`
	Plural      string                 `yaml:"plural"`
	Type        string                 `yaml:"type"`
	Title       string                 `yaml:"title"`
	Description string                 `yaml:"description"`
	Parents     map[string]*Reference  `yaml:"parents"`
	References  map[string]*Reference  `yaml:"references"`
	Prefix      string                 `yaml:"prefix"`
	JSONSchema  *JSONSchema            `yaml:"schema"`
	Definitions map[string]*JSONSchema `yaml:"definitions"`
	Extends     []string               `yaml:"extends"`
	Columns     []*ColumnConfig
}

type JSONSchema struct {
	Title          string                 `yaml:"title"`
	Description    string                 `yaml:"description"`
	SQL            string                 `yaml:"sql"`
	Default        interface{}            `yaml:"default"`
	Operation      string                 `yaml:"operation"`
	Presence       string                 `yaml:"presence"`
	Type           string                 `yaml:"type"`
	Permission     []string               `yaml:"permission"`
	Properties     map[string]*JSONSchema `yaml:"properties"`
	Enum           []string               `yaml:"enum"`
	Minimum        interface{}            `yaml:"minimum"`
	Maximum        interface{}            `yaml:"maximum"`
	Ref            string                 `yaml:"$ref"`
	CollectionType string                 `yaml:"collectionType"`
	Column         string                 `yaml:"column"`
	Item           *JSONSchema            `yaml:"item"`
	GoName         string
	GoType         string
}

func (s *JSONSchema) String() string {
	data, _ := json.Marshal(s)
	return string(data)
}

type Reference struct {
	Description string `json:"description"`
	Operation   string `json:"operation"`
	Presence    string `json:"presence"`
}

func (s *JSONSchema) GetRef() (string, string) {
	if s.Ref == "" {
		return "", ""
	}
	refs := strings.Split(s.Ref, "#")
	types := strings.Split(s.Ref, "/")
	return refs[0], types[2]
}

func (s *JSONSchema) Copy() *JSONSchema {

	copied := &JSONSchema{
		Title:      s.Title,
		SQL:        s.SQL,
		Default:    s.Default,
		Enum:       s.Enum,
		Minimum:    s.Minimum,
		Maximum:    s.Maximum,
		Ref:        s.Ref,
		Column:     s.Column,
		Permission: s.Permission,
		Operation:  s.Operation,
		Type:       s.Type,
		Presence:   s.Presence,
		Properties: map[string]*JSONSchema{},
	}
	for name, property := range s.Properties {
		copied.Properties[name] = property.Copy()
	}
	if s.Item != nil {
		copied.Item = s.Item.Copy()
	}
	return copied
}

func (s *JSONSchema) Update(s2 *JSONSchema) {
	if s2 == nil {
		return
	}
	if s.Title == "" {
		s.Title = s2.Title
	}
	if s.Description == "" {
		s.Description = s2.Description
	}
	if s.SQL == "" {
		s.SQL = s2.SQL
	}
	if s.Default == nil {
		s.Default = s2.Default
	}
	if s.Operation == "" {
		s.Operation = s2.Operation
	}
	if s.Presence == "" {
		s.Presence = s2.Presence
	}

	if s.Properties == nil {
		s.Properties = map[string]*JSONSchema{}
	}
	for name, property := range s2.Properties {
		if _, ok := s.Properties[name]; !ok {
			s.Properties[name] = property.Copy()
		}
	}
	if s.Type == "" {
		s.Type = s2.Type
	}

	if s.Enum == nil {
		s.Enum = s2.Enum
	}
	if s.Minimum == nil {
		s.Minimum = s2.Minimum
	}
	if s.Maximum == nil {
		s.Maximum = s2.Maximum
	}
}

func (s *JSONSchema) Walk(name string, do func(name string, s2 *JSONSchema) error) error {
	if s == nil {
		return nil
	}
	err := do(name, s)
	if err != nil {
		return err
	}
	if s.Properties == nil {
		return nil
	}
	for name, property := range s.Properties {
		err = property.Walk(name, do)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *JSONSchema) ResolveSQL(column string, goPath string, schema *Schema, used map[string]bool) error {
	if s == nil {
		return nil
	}
	if s.Column != "" {
		return nil
	}
	if len(s.Properties) == 0 || s.CollectionType != "" || s.Type == "array" {
		s.Column = column
		if s.SQL == "" {
			s.SQL = sqlTypeMap[s.Type]
		}
		bind := ""
		if s.GoType != "" {
			bind = sqlBindMap[s.Type]
		}
		schema.Columns = append(schema.Columns, &ColumnConfig{
			Path:   goPath,
			Bind:   bind,
			Type:   s.SQL,
			Column: column,
			Name:   strings.Replace(goPath, ".", "", -1),
		})
		return nil
	}
	for name, property := range s.Properties {
		columnName := name
		if _, ok := used[name]; column != "" && ok {
			columnName = column + "_" + name
		}
		used[columnName] = true
		err := property.ResolveSQL(columnName, goPath+"."+property.GoName, schema, used)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *JSONSchema) ResolveGoName(name string) error {
	if s == nil {
		return nil
	}
	s.GoName = strmangle.TitleCase(name)
	_, goType := s.GetRef()
	if goType == "" {
		switch s.Type {
		case "integer":
			goType = "int"
		case "number":
			goType = "float64"
		case "string":
			goType = "string"
		case "boolean":
			goType = "bool"
		case "object":
			if s.Properties == nil {
				goType = "map[string]interface{}"
			}
		case "array":
			s.Item.ResolveGoName(name)
			if s.Item == nil {
				goType = "[]interface{}"
			} else if s.Item.Type == "integer" || s.Item.Type == "number" || s.Item.Type == "boolean" || s.Item.Type == "string" {
				goType = "[]" + s.Item.GoType
			} else {
				goType = "[]*" + s.Item.GoType
			}
		}
	}
	// if s.CollectionType == "list" {
	// 	goType = "[]string"
	// }
	// if s.CollectionType == "map" {
	// 	goType = "map[string]string"
	// }
	s.GoType = goType
	for name, property := range s.Properties {
		err := property.ResolveGoName(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *API) schemaByFileName(fileName string) *Schema {
	for _, s := range api.Schemas {
		if s.FileName == fileName {
			return s
		}
	}
	return nil
}

func (api *API) schemaByID(id string) *Schema {
	for _, s := range api.Schemas {
		if s.ID == id {
			return s
		}
	}
	return nil
}

func (api *API) loadType(schema *JSONSchema) (*JSONSchema, error) {
	schemaFile, typeName := schema.GetRef()

	if definition, ok := api.Types[typeName]; ok {
		return definition, nil
	}
	definitions := api.schemaByFileName(schemaFile)
	if definitions == nil {
		return nil, fmt.Errorf("Can't find file for %s", schema.Ref)
	}
	definition, ok := definitions.Definitions[typeName]
	if !ok {
		return nil, fmt.Errorf("%s isn't defined", schema.Ref)
	}
	definition.Walk("", api.resolveRef)
	definition.ResolveGoName("")
	api.Types[typeName] = definition
	return definition, nil
}

func (api *API) resolveRef(name string, schema *JSONSchema) error {
	if schema == nil {
		return nil
	}
	if schema.Type == "array" {
		api.resolveRef("", schema.Item)
	}
	if schema.Ref == "" {
		return nil
	}
	definition, err := api.loadType(schema)
	if err != nil {
		return err
	}
	schema.Update(definition)
	return nil
}

func (api *API) resolveAllRef() error {
	for _, s := range api.Schemas {
		err := s.JSONSchema.Walk("", api.resolveRef)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *API) resolveAllSQL() error {
	for _, s := range api.Schemas {
		err := s.JSONSchema.ResolveSQL("", "", s, map[string]bool{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *API) resolveAllGoName() error {
	for _, s := range api.Schemas {
		s.JSONSchema.ResolveGoName(s.ID)
	}
	return nil
}

func (api *API) resolveExtend() error {
	for _, s := range api.Schemas {
		for _, baseSchemaID := range s.Extends {
			baseSchema := api.schemaByID(baseSchemaID)
			if baseSchema == nil {
				continue
			}
			s.JSONSchema.Update(baseSchema.JSONSchema)
		}
	}
	return nil
}

func MakeAPI(dir string) (*API, error) {
	api := &API{
		Schemas: []*Schema{},
		Types:   map[string]*JSONSchema{},
	}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		var schema Schema
		err = LoadFile(path, &schema)
		if err != nil {
			return nil
		}
		schema.FileName = strings.Replace(filepath.Base(path), ".yml", ".json", 1)
		if &schema == nil {
			return nil
		}
		api.Schemas = append(api.Schemas, &schema)
		return nil
	})
	err = api.resolveAllRef()
	if err != nil {
		return nil, err
	}
	err = api.resolveExtend()
	if err != nil {
		return nil, err
	}
	err = api.resolveAllGoName()
	if err != nil {
		return nil, err
	}
	err = api.resolveAllSQL()
	if err != nil {
		return nil, err
	}
	return api, err
}
