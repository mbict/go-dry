package request

import (
	"bytes"
	"github.com/google/uuid"
	"net/url"
	"strconv"
)

func EncodeInts(v []int) string {
	var buf bytes.Buffer
	for i, val := range v {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(strconv.Itoa(val))
	}
	return buf.String()
}

func EncodeStrings(v []string) string {
	var buf bytes.Buffer
	for i, val := range v {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(url.QueryEscape(val))
	}
	return buf.String()
}

func EncodeUUIDs(v []uuid.UUID) string {
	var buf bytes.Buffer
	for i, val := range v {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(val.String())
	}
	return buf.String()
}
