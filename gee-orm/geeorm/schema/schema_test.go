/**
 * @Author: lj
 * @Description:
 * @File:  schema_test
 * @Version: 1.0.0
 * @Date: 2022/03/25 20:50
 */

package schema

import (
	"fmt"
	"geeorm/dialect"
	"testing"
)

// schema_test.go
type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{Name: "name", Age: 34}, TestDial)
	fmt.Println(schema)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
