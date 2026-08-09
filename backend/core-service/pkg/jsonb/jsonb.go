package jsonb

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Document []byte

func FromBytes(raw []byte) Document {
	if len(raw) == 0 {
		return Document("null")
	}

	return Document(raw)
}

func Marshal(value any) (Document, error) {
	raw, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("jsonb: marshal failed: %w", err)
	}

	return Document(raw), nil
}

func (d Document) Unmarshal(target any) error {
	if len(d) == 0 {
		return nil
	}

	return json.Unmarshal(d, target)
}

func (d Document) String() string {
	if len(d) == 0 {
		return "null"
	}

	return string(d)
}

func (d Document) IsEmpty() bool {
	return len(d) == 0 || bytes.Equal(d, []byte("null"))
}

func (d *Document) Scan(src any) error {
	switch typed := src.(type) {
	case nil:
		*d = nil

		return nil
	case []byte:
		*d = append((*d)[:0], typed...)

		return nil
	case string:
		*d = Document(typed)

		return nil
	default:
		return fmt.Errorf("jsonb: cannot scan %T into Document", src)
	}
}

func (d Document) Value() (driver.Value, error) {
	if len(d) == 0 {
		return nil, nil
	}

	if !json.Valid(d) {
		return nil, fmt.Errorf("jsonb: refusing to store invalid JSON")
	}

	return string(d), nil
}

func (d Document) MarshalJSON() ([]byte, error) {
	if len(d) == 0 {
		return []byte("null"), nil
	}

	return d, nil
}

func (d *Document) UnmarshalJSON(data []byte) error {
	*d = append((*d)[:0], data...)

	return nil
}

func (Document) GormDataType() string {
	return "jsonb"
}
