/**
 * @Author: lj
 * @Description:
 * @File:  dialect
 * @Version: 1.0.0
 * @Date: 2022/03/25 15:01
 */

package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
