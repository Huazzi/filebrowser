// Code generated by go-enum
// DO NOT EDIT!

package http

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

const (
	// PreviewSizeThumb is a PreviewSize of type Thumb
	PreviewSizeThumb PreviewSize = iota
	// PreviewSizeBig is a PreviewSize of type Big
	PreviewSizeBig
)

const _PreviewSizeName = "thumbbig"

var _PreviewSizeNames = []string{
	_PreviewSizeName[0:5],
	_PreviewSizeName[5:8],
}

// PreviewSizeNames returns a list of possible string values of PreviewSize.
func PreviewSizeNames() []string {
	tmp := make([]string, len(_PreviewSizeNames))
	copy(tmp, _PreviewSizeNames)
	return tmp
}

var _PreviewSizeMap = map[PreviewSize]string{
	0: _PreviewSizeName[0:5],
	1: _PreviewSizeName[5:8],
}

// String implements the Stringer interface.
func (x PreviewSize) String() string {
	if str, ok := _PreviewSizeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PreviewSize(%d)", x)
}

var _PreviewSizeValue = map[string]PreviewSize{
	_PreviewSizeName[0:5]: 0,
	_PreviewSizeName[5:8]: 1,
}

// ParsePreviewSize attempts to convert a string to a PreviewSize
func ParsePreviewSize(name string) (PreviewSize, error) {
	if x, ok := _PreviewSizeValue[name]; ok {
		return x, nil
	}
	return PreviewSize(0), fmt.Errorf("%s is not a valid PreviewSize, try [%s]", name, strings.Join(_PreviewSizeNames, ", "))
}

// MarshalText implements the text marshaller method
func (x PreviewSize) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *PreviewSize) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePreviewSize(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Scan implements the Scanner interface.
func (x *PreviewSize) Scan(value interface{}) error {
	var name string

	switch v := value.(type) {
	case string:
		name = v
	case []byte:
		name = string(v)
	case nil:
		*x = PreviewSize(0)
		return nil
	}

	tmp, err := ParsePreviewSize(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Value implements the driver Valuer interface.
func (x PreviewSize) Value() (driver.Value, error) {
	return x.String(), nil
}
