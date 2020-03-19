package method

import (
	"strconv"
	"strings"
)

func UnescapeUnicode(str string) (string, error) {
	str2, err := strconv.Unquote(strings.Replace(strconv.Quote(str), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str2, nil
}
