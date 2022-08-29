package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// "implement" constraints for URL and *URL
var _ fmt.Stringer = URL{}

var _ json.Unmarshaler = (*URL)(nil)
var _ json.Marshaler = (*URL)(nil)

var _ sql.Scanner = (*URL)(nil)
var _ driver.Valuer = (*URL)(nil)

// URL http url
type URL struct {
	*url.URL
}

func (u *URL) Value() (driver.Value, error) {
	if u == nil {
		return "", nil
	}

	return u.String(), nil
}

func (u *URL) Scan(src any) error {
	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("unsupported data type '%T' for URLs.Scan", src)
	}

	return u.UnmarshalJSON([]byte(str))
}

// MarshalJSON returns the string value for URL
func (u *URL) MarshalJSON() ([]byte, error) {
	return []byte(`"` + u.String() + `"`), nil
}

// UnmarshalJSON decodes the string value into *url.URL
func (u *URL) UnmarshalJSON(bytes []byte) (err error) {
	rawURL := string(bytes)

	if strings.HasPrefix(rawURL, `"`) && strings.HasSuffix(rawURL, `"`) {
		rawURL = rawURL[1 : len(rawURL)-1]
	}

	u.URL, err = url.Parse(rawURL)
	return
}

// String returns the raw string value for URL
func (u URL) String() string {
	if u.URL == nil {
		return ""
	}
	return u.URL.String()
}

// "implement" constraints for URLs
var _ sql.Scanner = (*URLs)(nil)
var _ driver.Valuer = (URLs)(nil)

type URLs []URL

func (u URLs) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *URLs) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unsupported data type '%T' for URLs.Scan", src)
	}

	return json.Unmarshal(bytes, u)
}
