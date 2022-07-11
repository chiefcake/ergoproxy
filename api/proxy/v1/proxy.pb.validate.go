// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/proxy/v1/proxy.proto

package proxyv1

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

// Validate checks the field values on RedirectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RedirectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RedirectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RedirectRequestMultiError, or nil if none found.
func (m *RedirectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RedirectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _RedirectRequest_Method_InLookup[m.GetMethod()]; !ok {
		err := RedirectRequestValidationError{
			field:  "Method",
			reason: "value must be in list [POST GET PUT PATCH DELETE]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		err = RedirectRequestValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := RedirectRequestValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Headers

	// no validation rules for Body

	if len(errors) > 0 {
		return RedirectRequestMultiError(errors)
	}

	return nil
}

// RedirectRequestMultiError is an error wrapping multiple validation errors
// returned by RedirectRequest.ValidateAll() if the designated constraints
// aren't met.
type RedirectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RedirectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RedirectRequestMultiError) AllErrors() []error { return m }

// RedirectRequestValidationError is the validation error returned by
// RedirectRequest.Validate if the designated constraints aren't met.
type RedirectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedirectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedirectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedirectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedirectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedirectRequestValidationError) ErrorName() string { return "RedirectRequestValidationError" }

// Error satisfies the builtin error interface
func (e RedirectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedirectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedirectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedirectRequestValidationError{}

var _RedirectRequest_Method_InLookup = map[string]struct{}{
	"POST":   {},
	"GET":    {},
	"PUT":    {},
	"PATCH":  {},
	"DELETE": {},
}

// Validate checks the field values on RedirectResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RedirectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RedirectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RedirectResponseMultiError, or nil if none found.
func (m *RedirectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RedirectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Status

	// no validation rules for Headers

	// no validation rules for Length

	// no validation rules for Body

	if len(errors) > 0 {
		return RedirectResponseMultiError(errors)
	}

	return nil
}

// RedirectResponseMultiError is an error wrapping multiple validation errors
// returned by RedirectResponse.ValidateAll() if the designated constraints
// aren't met.
type RedirectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RedirectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RedirectResponseMultiError) AllErrors() []error { return m }

// RedirectResponseValidationError is the validation error returned by
// RedirectResponse.Validate if the designated constraints aren't met.
type RedirectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedirectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedirectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedirectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedirectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedirectResponseValidationError) ErrorName() string { return "RedirectResponseValidationError" }

// Error satisfies the builtin error interface
func (e RedirectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedirectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedirectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedirectResponseValidationError{}
