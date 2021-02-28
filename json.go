package kissorm

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// This type was created to make it easier to adapt
// input attributes to be convertible to and from JSON
// before sending or receiving it from the database.
type jsonSerializable struct {
	Attr interface{}
}

// Scan Implements the Scanner interface in order to load
// this field from the JSON stored in the database
func (j *jsonSerializable) Scan(value interface{}) error {
	rawJSON, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unexpected type received to Scan: %T", value)
	}
	return json.Unmarshal(rawJSON, j.Attr)
}

// Value Implements the Valuer interface in order to save
// this field as JSON on the database.
func (j jsonSerializable) Value() (driver.Value, error) {
	return json.Marshal(j.Attr)
}