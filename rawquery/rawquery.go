package rawquery

import (
	"fmt"
	"net/url"
	"strings"
)

type Query struct {
	Field string
	Value interface{}
}

func (q Query) String() string {
	return fmt.Sprintf("%s:%s", q.Field, q.Value)
}

// ParseRawQuery parse raw query into the Query slice.
func ParseRawQuery(rawquery string) []Query {
	slice := strings.Split(rawquery, "&")
	parse := make([]Query, 0, len(slice))

	for _, pair := range slice {
		pairs := strings.Split(pair, "=")
		if len(pairs) == 2 {
			parse = append(parse, Query{
				Field: pairs[0],
				Value: pairs[1],
			})
		}
	}

	return parse
}

// NewRawQueryWithQueries queries are used to generate the original HTTP URL query string.
func NewRawQueryWithQueries(queries []Query) string {
	var s = make([]string, 0)
	for _, query := range queries {
		s = append(s, fmt.Sprintf("%s=%v", query.Field, query.Value))
	}
	return strings.Join(s, "&")
}

// MergeURLRawQuery queries are used to generate the original HTTP URL query string.
func MergeURLRawQuery(rawurl string, rawquery string) (string, error) {
	parse, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	if rawquery == "" {
		return parse.RawQuery, nil
	}

	original := ParseRawQuery(parse.RawQuery)
	external := ParseRawQuery(rawquery)
	for _, item := range external {
		var exists bool
		var index int
		for idx, ori := range original {
			if item.String() == ori.String() {
				exists, index = true, idx
			}
		}
		if exists {
			original[index] = item
		} else {
			original = append(original, item)
		}
	}

	return NewRawQueryWithQueries(original), nil
}
