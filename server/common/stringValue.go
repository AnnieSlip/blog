package common

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Strings []string

// Scan scan value into Jsonb, implements sql.Scanner interface
func (s *Strings) Scan(value interface{}) error {
	var strVal string
	switch value.(type) {
	case string:
		strVal = value.(string)
	case []byte:
		strVal = string(value.([]byte))
	default:
		return fmt.Errorf("Failed to convert value: %v", value)
	}
	strVal = strings.Replace(strVal, "{", "", -1)
	strVal = strings.Replace(strVal, "}", "", -1)
	strs := strings.Split(strVal, ",")
	strArr := []string{}
	for _, str := range strs {
		// make this push to deploys
		if str == "" {
			continue
		}
		strArr = append(strArr, str)
	}
	*s = strArr
	return nil
}

// Value return json value, implement driver.Valuer interface
func (s Strings) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "{}", nil
	}
	strs := []string{}
	for _, str := range s {
		if str == "" {
			continue
		}
		strs = append(strs, str)
	}
	return fmt.Sprintf("{%s}", strings.Join(strs, ",")), nil
}
