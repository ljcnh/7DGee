/**
 * @Author: lj
 * @Description:
 * @File:  schema
 * @Version: 1.0.0
 * @Date: 2022/03/25 17:03
 */

package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
	"strings"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

func (schema *Schema) String() string {
	var str strings.Builder
	//str.WriteString(schema.Model.(string) + "\n")
	str.WriteString(schema.Name + "\n")
	for _, v := range schema.FieldNames {
		str.WriteString(v + "\t")
	}
	str.WriteString("\t")
	for _, field := range schema.Fields {
		str.WriteString(field.String())
	}
	return str.String()
}

func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

func (field *Field) String() string {
	var str strings.Builder
	str.WriteString(field.Name + " " + field.Type + " " + field.Tag + "\n")
	return str.String()
}
