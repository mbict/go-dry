package request

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/url"
	"strconv"
	"strings"
)

func ParseInts(v string) ([]int, error) {
	v = strings.TrimSpace(v)
	if len(v) == 0 {
		return nil, nil
	}

	var err error
	rs := strings.Split(v, ",")
	ints := make([]int, len(rs))
	for i, val := range rs {
		ints[i], err = strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("has a invalid integer value `%s`", val)
		}
	}
	return ints, nil
}

func ParseStrings(v string) ([]string, error) {
	v = strings.TrimSpace(v)
	if len(v) == 0 {
		return nil, nil
	}

	strs := strings.Split(v, ",")
	for i := len(strs) - 1; i >= 0; i-- {
		str, err := url.QueryUnescape(strs[i])
		if err != nil {
			return nil, fmt.Errorf("has a invalid url encoded value `%s`", strs[i])
		}
		str = strings.TrimSpace(str)

		//remove empty elements
		if len(str) == 0 {
			strs = append(strs[:i], strs[i+1:]...)
			continue
		}
		strs[i] = str
	}
	if len(strs) == 0 {
		return nil, nil
	}
	return strs, nil
}

func ParseUUIDs(v string) ([]uuid.UUID, error) {
	v = strings.TrimSpace(v)
	if len(v) == 0 {
		return nil, nil
	}

	var err error
	rs := strings.Split(v, ",")
	uuids := make([]uuid.UUID, len(rs))
	for i, val := range rs {
		uuids[i], err = uuid.FromString(val)
		if err != nil {
			return nil, fmt.Errorf("has a invalid uuid value `%s`", val)
		}
	}
	return uuids, nil
}
