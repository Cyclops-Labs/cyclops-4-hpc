// For furture references related to this:
// https://coussejhub.io/2016/02/16/Handling-JSONB-in-Go-Structs/
// https://github.com/jinzhu/gorm/issues/1155

package datamodels

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONdb type
//
// This type is meant to be used when the double interaction between GORM and
// JSON (un)marshalling is needed
type JSONdb map[string]interface{}

// Value method for the interface.
//
// To satisfy this interface, we must implement the Value method, which
// must transform our type to a database driver compatible type.
// In our case, we’ll marshall the map to JSONB data (= []byte):
func (j JSONdb) Value() (driver.Value, error) {

	m, e := json.Marshal(j)

	return m, e

}

// Scan method for the interface.
//
// This method must take the raw data that comes from the database
// and transform it to our new type. In our case, the database will
// return JSONB ([]byte) that we must transform to our type
// (the reverse of what we did with driver.Valuer):
func (p *JSONdb) Scan(src interface{}) error {

	source, ok := src.([]byte)

	if !ok {

		return errors.New("type assertion .([]byte) failed")

	}

	var i interface{}

	e := json.Unmarshal(source, &i)

	if e != nil {

		return e

	}

	*p, ok = i.(map[string]interface{})

	if !ok {

		return errors.New("type assertion .(map[string]interface{}) failed")

	}

	return nil

}
