// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kioku/srv-dba/v1/srv-dba-kanji.proto

package srv_dba

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Kanji with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Kanji) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Kanji with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in KanjiMultiError, or nil if none found.
func (m *Kanji) ValidateAll() error {
	return m.validate(true)
}

func (m *Kanji) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Kanji

	// no validation rules for Primary

	// no validation rules for Level

	if len(errors) > 0 {
		return KanjiMultiError(errors)
	}
	return nil
}

// KanjiMultiError is an error wrapping multiple validation errors returned by
// Kanji.ValidateAll() if the designated constraints aren't met.
type KanjiMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KanjiMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KanjiMultiError) AllErrors() []error { return m }

// KanjiValidationError is the validation error returned by Kanji.Validate if
// the designated constraints aren't met.
type KanjiValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KanjiValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KanjiValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KanjiValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KanjiValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KanjiValidationError) ErrorName() string { return "KanjiValidationError" }

// Error satisfies the builtin error interface
func (e KanjiValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKanji.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KanjiValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KanjiValidationError{}

// Validate checks the field values on GetKanjiV1Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetKanjiV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetKanjiV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetKanjiV1RequestMultiError, or nil if none found.
func (m *GetKanjiV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *GetKanjiV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetKanjiId() <= 0 {
		err := GetKanjiV1RequestValidationError{
			field:  "KanjiId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetKanjiV1RequestMultiError(errors)
	}
	return nil
}

// GetKanjiV1RequestMultiError is an error wrapping multiple validation errors
// returned by GetKanjiV1Request.ValidateAll() if the designated constraints
// aren't met.
type GetKanjiV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetKanjiV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetKanjiV1RequestMultiError) AllErrors() []error { return m }

// GetKanjiV1RequestValidationError is the validation error returned by
// GetKanjiV1Request.Validate if the designated constraints aren't met.
type GetKanjiV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetKanjiV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetKanjiV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetKanjiV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetKanjiV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetKanjiV1RequestValidationError) ErrorName() string {
	return "GetKanjiV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetKanjiV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetKanjiV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetKanjiV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetKanjiV1RequestValidationError{}

// Validate checks the field values on GetKanjiV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetKanjiV1Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetKanjiV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetKanjiV1ResponseMultiError, or nil if none found.
func (m *GetKanjiV1Response) ValidateAll() error {
	return m.validate(true)
}

func (m *GetKanjiV1Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetKanji()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetKanjiV1ResponseValidationError{
					field:  "Kanji",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetKanjiV1ResponseValidationError{
					field:  "Kanji",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKanji()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetKanjiV1ResponseValidationError{
				field:  "Kanji",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetKanjiV1ResponseMultiError(errors)
	}
	return nil
}

// GetKanjiV1ResponseMultiError is an error wrapping multiple validation errors
// returned by GetKanjiV1Response.ValidateAll() if the designated constraints
// aren't met.
type GetKanjiV1ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetKanjiV1ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetKanjiV1ResponseMultiError) AllErrors() []error { return m }

// GetKanjiV1ResponseValidationError is the validation error returned by
// GetKanjiV1Response.Validate if the designated constraints aren't met.
type GetKanjiV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetKanjiV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetKanjiV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetKanjiV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetKanjiV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetKanjiV1ResponseValidationError) ErrorName() string {
	return "GetKanjiV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetKanjiV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetKanjiV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetKanjiV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetKanjiV1ResponseValidationError{}

// Validate checks the field values on ListKanjiV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListKanjiV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListKanjiV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListKanjiV1RequestMultiError, or nil if none found.
func (m *ListKanjiV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *ListKanjiV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetLevel() <= 0 {
		err := ListKanjiV1RequestValidationError{
			field:  "Level",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListKanjiV1RequestMultiError(errors)
	}
	return nil
}

// ListKanjiV1RequestMultiError is an error wrapping multiple validation errors
// returned by ListKanjiV1Request.ValidateAll() if the designated constraints
// aren't met.
type ListKanjiV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListKanjiV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListKanjiV1RequestMultiError) AllErrors() []error { return m }

// ListKanjiV1RequestValidationError is the validation error returned by
// ListKanjiV1Request.Validate if the designated constraints aren't met.
type ListKanjiV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListKanjiV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListKanjiV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListKanjiV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListKanjiV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListKanjiV1RequestValidationError) ErrorName() string {
	return "ListKanjiV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListKanjiV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListKanjiV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListKanjiV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListKanjiV1RequestValidationError{}

// Validate checks the field values on ListKanjiV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListKanjiV1Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListKanjiV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListKanjiV1ResponseMultiError, or nil if none found.
func (m *ListKanjiV1Response) ValidateAll() error {
	return m.validate(true)
}

func (m *ListKanjiV1Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetKanji() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListKanjiV1ResponseValidationError{
						field:  fmt.Sprintf("Kanji[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListKanjiV1ResponseValidationError{
						field:  fmt.Sprintf("Kanji[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListKanjiV1ResponseValidationError{
					field:  fmt.Sprintf("Kanji[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListKanjiV1ResponseMultiError(errors)
	}
	return nil
}

// ListKanjiV1ResponseMultiError is an error wrapping multiple validation
// errors returned by ListKanjiV1Response.ValidateAll() if the designated
// constraints aren't met.
type ListKanjiV1ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListKanjiV1ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListKanjiV1ResponseMultiError) AllErrors() []error { return m }

// ListKanjiV1ResponseValidationError is the validation error returned by
// ListKanjiV1Response.Validate if the designated constraints aren't met.
type ListKanjiV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListKanjiV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListKanjiV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListKanjiV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListKanjiV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListKanjiV1ResponseValidationError) ErrorName() string {
	return "ListKanjiV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListKanjiV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListKanjiV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListKanjiV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListKanjiV1ResponseValidationError{}