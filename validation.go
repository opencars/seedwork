package seedwork

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	Required     = "required"
	Invalid      = "invalid"
	IsNotInreger = "is_not_integer"
)

func Validate(v validation.Validatable, prefix string) error {
	err := v.Validate()
	if err == nil {
		return nil
	}

	messages := ErrorMessages(prefix, err)
	return NewValidationError(messages)
}

func ErrorMessages(prefix string, err error) map[string][]string {
	errs, ok := err.(validation.Errors)
	if !ok {
		return nil
	}

	messages := make(map[string][]string)
	for k, v := range errs {
		if _, ok := v.(validation.Errors); ok {
			key := fmt.Sprintf("%s.%s", prefix, ToSnakeCase(k))
			errMap := ErrorMessages(key, v)
			for k, v := range errMap {
				messages[ToSnakeCase(k)] = v
			}
		} else {
			key := fmt.Sprintf("%s.%s", prefix, ToSnakeCase(k))
			messages[key] = append(messages[key], v.Error())
		}
	}

	return messages
}

func ToSnakeCase(str string) string {
	return ToScreamingDelimited(str, '_')
}

func ToScreamingDelimited(s string, delimiter uint8) string {
	s = strings.TrimSpace(s)

	n := strings.Builder{}
	n.Grow(len(s) + 2) // nominal 2 bytes of extra space for inserted delimiters

	if len(s) != 0 {
		currIsCap := s[0] >= 'A' && s[0] <= 'Z'

		if currIsCap {
			n.WriteByte(s[0] - 'A' + 'a')
		} else {
			n.WriteByte(s[0])
		}
	}

	for i := 1; i < len(s); i++ {
		prevIsCap := s[i] >= 'A' && s[i] <= 'Z'
		prevIsLow := s[i-1] >= 'a' && s[i-1] <= 'z'

		currIsCap := s[i] >= 'A' && s[i] <= 'Z'

		hasNext := len(s) > i+1
		var nextIsLow bool
		var next byte
		if hasNext {
			next = s[i+1]
			nextIsLow = s[i+1] >= 'a' && s[i+1] <= 'z'
		}

		if prevIsLow && currIsCap {
			n.WriteByte(delimiter)
		} else if currIsCap && prevIsCap && hasNext && nextIsLow && next != 's' {
			n.WriteByte(delimiter)
		}

		if currIsCap {
			n.WriteByte(s[i] - 'A' + 'a')
		} else {
			n.WriteByte(s[i])
		}
	}

	return n.String()
}
