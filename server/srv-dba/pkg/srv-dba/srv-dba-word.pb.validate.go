// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kioku/srv-dba/v1/srv-dba-word.proto

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

// Validate checks the field values on SentenceTranslation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SentenceTranslation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SentenceTranslation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SentenceTranslationMultiError, or nil if none found.
func (m *SentenceTranslation) ValidateAll() error {
	return m.validate(true)
}

func (m *SentenceTranslation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Language

	// no validation rules for Translation

	if len(errors) > 0 {
		return SentenceTranslationMultiError(errors)
	}
	return nil
}

// SentenceTranslationMultiError is an error wrapping multiple validation
// errors returned by SentenceTranslation.ValidateAll() if the designated
// constraints aren't met.
type SentenceTranslationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SentenceTranslationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SentenceTranslationMultiError) AllErrors() []error { return m }

// SentenceTranslationValidationError is the validation error returned by
// SentenceTranslation.Validate if the designated constraints aren't met.
type SentenceTranslationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SentenceTranslationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SentenceTranslationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SentenceTranslationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SentenceTranslationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SentenceTranslationValidationError) ErrorName() string {
	return "SentenceTranslationValidationError"
}

// Error satisfies the builtin error interface
func (e SentenceTranslationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSentenceTranslation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SentenceTranslationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SentenceTranslationValidationError{}

// Validate checks the field values on Sentence with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Sentence) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Sentence with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SentenceMultiError, or nil
// if none found.
func (m *Sentence) ValidateAll() error {
	return m.validate(true)
}

func (m *Sentence) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Origin

	for idx, item := range m.GetTranslations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SentenceValidationError{
						field:  fmt.Sprintf("Translations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SentenceValidationError{
						field:  fmt.Sprintf("Translations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SentenceValidationError{
					field:  fmt.Sprintf("Translations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SentenceMultiError(errors)
	}
	return nil
}

// SentenceMultiError is an error wrapping multiple validation errors returned
// by Sentence.ValidateAll() if the designated constraints aren't met.
type SentenceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SentenceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SentenceMultiError) AllErrors() []error { return m }

// SentenceValidationError is the validation error returned by
// Sentence.Validate if the designated constraints aren't met.
type SentenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SentenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SentenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SentenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SentenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SentenceValidationError) ErrorName() string { return "SentenceValidationError" }

// Error satisfies the builtin error interface
func (e SentenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSentence.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SentenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SentenceValidationError{}

// Validate checks the field values on Word with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Word) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Word with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in WordMultiError, or nil if none found.
func (m *Word) ValidateAll() error {
	return m.validate(true)
}

func (m *Word) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Word

	// no validation rules for Primary

	// no validation rules for Level

	for idx, item := range m.GetComposition() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WordValidationError{
						field:  fmt.Sprintf("Composition[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WordValidationError{
						field:  fmt.Sprintf("Composition[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WordValidationError{
					field:  fmt.Sprintf("Composition[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSentences() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WordValidationError{
						field:  fmt.Sprintf("Sentences[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WordValidationError{
						field:  fmt.Sprintf("Sentences[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WordValidationError{
					field:  fmt.Sprintf("Sentences[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return WordMultiError(errors)
	}
	return nil
}

// WordMultiError is an error wrapping multiple validation errors returned by
// Word.ValidateAll() if the designated constraints aren't met.
type WordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WordMultiError) AllErrors() []error { return m }

// WordValidationError is the validation error returned by Word.Validate if the
// designated constraints aren't met.
type WordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WordValidationError) ErrorName() string { return "WordValidationError" }

// Error satisfies the builtin error interface
func (e WordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWord.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WordValidationError{}

// Validate checks the field values on GetWordByIdV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetWordByIdV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetWordByIdV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetWordByIdV1RequestMultiError, or nil if none found.
func (m *GetWordByIdV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *GetWordByIdV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetWordId() <= 0 {
		err := GetWordByIdV1RequestValidationError{
			field:  "WordId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetWordByIdV1RequestMultiError(errors)
	}
	return nil
}

// GetWordByIdV1RequestMultiError is an error wrapping multiple validation
// errors returned by GetWordByIdV1Request.ValidateAll() if the designated
// constraints aren't met.
type GetWordByIdV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetWordByIdV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetWordByIdV1RequestMultiError) AllErrors() []error { return m }

// GetWordByIdV1RequestValidationError is the validation error returned by
// GetWordByIdV1Request.Validate if the designated constraints aren't met.
type GetWordByIdV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetWordByIdV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetWordByIdV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetWordByIdV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetWordByIdV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetWordByIdV1RequestValidationError) ErrorName() string {
	return "GetWordByIdV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetWordByIdV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWordByIdV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetWordByIdV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetWordByIdV1RequestValidationError{}

// Validate checks the field values on GetWordByIdV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetWordByIdV1Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetWordByIdV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetWordByIdV1ResponseMultiError, or nil if none found.
func (m *GetWordByIdV1Response) ValidateAll() error {
	return m.validate(true)
}

func (m *GetWordByIdV1Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWord()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetWordByIdV1ResponseValidationError{
					field:  "Word",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetWordByIdV1ResponseValidationError{
					field:  "Word",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWord()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetWordByIdV1ResponseValidationError{
				field:  "Word",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetWordByIdV1ResponseMultiError(errors)
	}
	return nil
}

// GetWordByIdV1ResponseMultiError is an error wrapping multiple validation
// errors returned by GetWordByIdV1Response.ValidateAll() if the designated
// constraints aren't met.
type GetWordByIdV1ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetWordByIdV1ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetWordByIdV1ResponseMultiError) AllErrors() []error { return m }

// GetWordByIdV1ResponseValidationError is the validation error returned by
// GetWordByIdV1Response.Validate if the designated constraints aren't met.
type GetWordByIdV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetWordByIdV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetWordByIdV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetWordByIdV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetWordByIdV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetWordByIdV1ResponseValidationError) ErrorName() string {
	return "GetWordByIdV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetWordByIdV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWordByIdV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetWordByIdV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetWordByIdV1ResponseValidationError{}

// Validate checks the field values on ListWordsByLevelV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListWordsByLevelV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWordsByLevelV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListWordsByLevelV1RequestMultiError, or nil if none found.
func (m *ListWordsByLevelV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWordsByLevelV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetLevel() <= 0 {
		err := ListWordsByLevelV1RequestValidationError{
			field:  "Level",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() >= 1000 {
		err := ListWordsByLevelV1RequestValidationError{
			field:  "Limit",
			reason: "value must be less than 1000",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Offset

	// no validation rules for Min

	if len(errors) > 0 {
		return ListWordsByLevelV1RequestMultiError(errors)
	}
	return nil
}

// ListWordsByLevelV1RequestMultiError is an error wrapping multiple validation
// errors returned by ListWordsByLevelV1Request.ValidateAll() if the
// designated constraints aren't met.
type ListWordsByLevelV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWordsByLevelV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWordsByLevelV1RequestMultiError) AllErrors() []error { return m }

// ListWordsByLevelV1RequestValidationError is the validation error returned by
// ListWordsByLevelV1Request.Validate if the designated constraints aren't met.
type ListWordsByLevelV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWordsByLevelV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWordsByLevelV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWordsByLevelV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWordsByLevelV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWordsByLevelV1RequestValidationError) ErrorName() string {
	return "ListWordsByLevelV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListWordsByLevelV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWordsByLevelV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWordsByLevelV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWordsByLevelV1RequestValidationError{}

// Validate checks the field values on ListWordsByKanjiV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListWordsByKanjiV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWordsByKanjiV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListWordsByKanjiV1RequestMultiError, or nil if none found.
func (m *ListWordsByKanjiV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWordsByKanjiV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetKanjiId() <= 0 {
		err := ListWordsByKanjiV1RequestValidationError{
			field:  "KanjiId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() >= 1000 {
		err := ListWordsByKanjiV1RequestValidationError{
			field:  "Limit",
			reason: "value must be less than 1000",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Offset

	// no validation rules for Min

	if len(errors) > 0 {
		return ListWordsByKanjiV1RequestMultiError(errors)
	}
	return nil
}

// ListWordsByKanjiV1RequestMultiError is an error wrapping multiple validation
// errors returned by ListWordsByKanjiV1Request.ValidateAll() if the
// designated constraints aren't met.
type ListWordsByKanjiV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWordsByKanjiV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWordsByKanjiV1RequestMultiError) AllErrors() []error { return m }

// ListWordsByKanjiV1RequestValidationError is the validation error returned by
// ListWordsByKanjiV1Request.Validate if the designated constraints aren't met.
type ListWordsByKanjiV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWordsByKanjiV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWordsByKanjiV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWordsByKanjiV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWordsByKanjiV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWordsByKanjiV1RequestValidationError) ErrorName() string {
	return "ListWordsByKanjiV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListWordsByKanjiV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWordsByKanjiV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWordsByKanjiV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWordsByKanjiV1RequestValidationError{}

// Validate checks the field values on ListWordsByKanjiAndLevelV1Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *ListWordsByKanjiAndLevelV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWordsByKanjiAndLevelV1Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// ListWordsByKanjiAndLevelV1RequestMultiError, or nil if none found.
func (m *ListWordsByKanjiAndLevelV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWordsByKanjiAndLevelV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Level

	// no validation rules for KanjiId

	if m.GetLimit() >= 1000 {
		err := ListWordsByKanjiAndLevelV1RequestValidationError{
			field:  "Limit",
			reason: "value must be less than 1000",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Offset

	// no validation rules for Min

	if len(errors) > 0 {
		return ListWordsByKanjiAndLevelV1RequestMultiError(errors)
	}
	return nil
}

// ListWordsByKanjiAndLevelV1RequestMultiError is an error wrapping multiple
// validation errors returned by
// ListWordsByKanjiAndLevelV1Request.ValidateAll() if the designated
// constraints aren't met.
type ListWordsByKanjiAndLevelV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWordsByKanjiAndLevelV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWordsByKanjiAndLevelV1RequestMultiError) AllErrors() []error { return m }

// ListWordsByKanjiAndLevelV1RequestValidationError is the validation error
// returned by ListWordsByKanjiAndLevelV1Request.Validate if the designated
// constraints aren't met.
type ListWordsByKanjiAndLevelV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWordsByKanjiAndLevelV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWordsByKanjiAndLevelV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWordsByKanjiAndLevelV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWordsByKanjiAndLevelV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWordsByKanjiAndLevelV1RequestValidationError) ErrorName() string {
	return "ListWordsByKanjiAndLevelV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListWordsByKanjiAndLevelV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWordsByKanjiAndLevelV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWordsByKanjiAndLevelV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWordsByKanjiAndLevelV1RequestValidationError{}

// Validate checks the field values on ListWordsByIdsV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListWordsByIdsV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWordsByIdsV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListWordsByIdsV1RequestMultiError, or nil if none found.
func (m *ListWordsByIdsV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWordsByIdsV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetWordId() {
		_, _ = idx, item

		if item <= 0 {
			err := ListWordsByIdsV1RequestValidationError{
				field:  fmt.Sprintf("WordId[%v]", idx),
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for Min

	if len(errors) > 0 {
		return ListWordsByIdsV1RequestMultiError(errors)
	}
	return nil
}

// ListWordsByIdsV1RequestMultiError is an error wrapping multiple validation
// errors returned by ListWordsByIdsV1Request.ValidateAll() if the designated
// constraints aren't met.
type ListWordsByIdsV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWordsByIdsV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWordsByIdsV1RequestMultiError) AllErrors() []error { return m }

// ListWordsByIdsV1RequestValidationError is the validation error returned by
// ListWordsByIdsV1Request.Validate if the designated constraints aren't met.
type ListWordsByIdsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWordsByIdsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWordsByIdsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWordsByIdsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWordsByIdsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWordsByIdsV1RequestValidationError) ErrorName() string {
	return "ListWordsByIdsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListWordsByIdsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWordsByIdsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWordsByIdsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWordsByIdsV1RequestValidationError{}

// Validate checks the field values on ListWordsV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListWordsV1Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWordsV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListWordsV1ResponseMultiError, or nil if none found.
func (m *ListWordsV1Response) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWordsV1Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetWords() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListWordsV1ResponseValidationError{
						field:  fmt.Sprintf("Words[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListWordsV1ResponseValidationError{
						field:  fmt.Sprintf("Words[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListWordsV1ResponseValidationError{
					field:  fmt.Sprintf("Words[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListWordsV1ResponseMultiError(errors)
	}
	return nil
}

// ListWordsV1ResponseMultiError is an error wrapping multiple validation
// errors returned by ListWordsV1Response.ValidateAll() if the designated
// constraints aren't met.
type ListWordsV1ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWordsV1ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWordsV1ResponseMultiError) AllErrors() []error { return m }

// ListWordsV1ResponseValidationError is the validation error returned by
// ListWordsV1Response.Validate if the designated constraints aren't met.
type ListWordsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWordsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWordsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWordsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWordsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWordsV1ResponseValidationError) ErrorName() string {
	return "ListWordsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListWordsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWordsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWordsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWordsV1ResponseValidationError{}
