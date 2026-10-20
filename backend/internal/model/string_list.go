package model

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringList []string

func (list StringList) Value() (driver.Value, error) {
	encoded, err := list.JSON()
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

func (list *StringList) Scan(value any) error {
	if value == nil {
		*list = StringList{}
		return nil
	}

	var data []byte

	switch storedValue := value.(type) {
	case []byte:
		data = storedValue

	case string:
		data = []byte(storedValue)

	default:
		return fmt.Errorf(
			"scan string list: unsupported database type %T",
			value,
		)
	}

	data = bytes.TrimSpace(data)

	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		*list = StringList{}
		return nil
	}

	var result []string

	if err := json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf(
			"scan string list: invalid JSON %q: %w",
			string(data),
			err,
		)
	}

	*list = StringList(result)

	return nil
}

// JSON returns JSON text suitable for storing in SQLite.
func (list StringList) JSON() (string, error) {
	if list == nil {
		return "[]", nil
	}

	encoded, err := json.Marshal(list)
	if err != nil {
		return "", fmt.Errorf("encode string list: %w", err)
	}

	return string(encoded), nil
}
