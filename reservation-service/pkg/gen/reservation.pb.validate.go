package pb

import (
	"bytes"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/anypb"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

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
var _reservation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

func (m *CreateReservationRequest) Validate() error {
	return m.validate(false)
}
func (m *CreateReservationRequest) ValidateAll() error {
	return m.validate(true)
}
func (m *CreateReservationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}
	var errors []error
	if err := m._validateUuid(m.GetReservationId()); err != nil {
		err = CreateReservationRequestValidationError{field: "ReservationId", reason: "value must be a valid UUID", cause: err}
		if !all {
			return err
		}
		errors = append(errors, err)
	}
	if err := m._validateUuid(m.GetProductId()); err != nil {
		err = CreateReservationRequestValidationError{field: "ProductId", reason: "value must be a valid UUID", cause: err}
		if !all {
			return err
		}
		errors = append(errors, err)
	}
	if err := m._validateUuid(m.GetCustomerId()); err != nil {
		err = CreateReservationRequestValidationError{field: "CustomerId", reason: "value must be a valid UUID", cause: err}
		if !all {
			return err
		}
		errors = append(errors, err)
	}
	if m.GetQuantity() <= 0 {
		err := CreateReservationRequestValidationError{field: "Quantity", reason: "value must be greater than 0"}
		if !all {
			return err
		}
		errors = append(errors, err)
	}
	if val := m.GetExpirationMinutes(); val < 1 || val > 1440 {
		err := CreateReservationRequestValidationError{field: "ExpirationMinutes", reason: "value must be inside range [1, 1440]"}
		if !all {
			return err
		}
		errors = append(errors, err)
	}
	if len(errors) > 0 {
		return CreateReservationRequestMultiError(errors)
	}
	return nil
}
func (m *CreateReservationRequest) _validateUuid(uuid string) error {
	if matched := _reservation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}
	return nil
}

type CreateReservationRequestMultiError []error

func (m CreateReservationRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}
func (m CreateReservationRequestMultiError) AllErrors() []error {
	return m
}

type CreateReservationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

func (e CreateReservationRequestValidationError) Field() string {
	return e.field
}
func (e CreateReservationRequestValidationError) Reason() string {
	return e.reason
}
func (e CreateReservationRequestValidationError) Cause() error {
	return e.cause
}
func (e CreateReservationRequestValidationError) Key() bool {
	return e.key
}
func (e CreateReservationRequestValidationError) ErrorName() string {
	return "CreateReservationRequestValidationError"
}
func (e CreateReservationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}
	key := ""
	if e.key {
		key = "key for "
	}
	return fmt.Sprintf("invalid %sCreateReservationRequest.%s: %s%s", key, e.field, e.reason, cause)
}

var _ error = CreateReservationRequestValidationError{}
var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReservationRequestValidationError{}

func (m *ReservationRequest) Validate() error {
	return m.validate(false)
}
func (m *ReservationRequest) ValidateAll() error {
	return m.validate(true)
}
func (m *ReservationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}
	var errors []error
	if err := m._validateUuid(m.GetReservationId()); err != nil {
		err = ReservationRequestValidationError{field: "ReservationId", reason: "value must be a valid UUID", cause: err}
		if !all {
			return err
		}
		errors = append(errors, err)
	}
	if len(errors) > 0 {
		return ReservationRequestMultiError(errors)
	}
	return nil
}
func (m *ReservationRequest) _validateUuid(uuid string) error {
	if matched := _reservation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}
	return nil
}

type ReservationRequestMultiError []error

func (m ReservationRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}
func (m ReservationRequestMultiError) AllErrors() []error {
	return m
}

type ReservationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

func (e ReservationRequestValidationError) Field() string {
	return e.field
}
func (e ReservationRequestValidationError) Reason() string {
	return e.reason
}
func (e ReservationRequestValidationError) Cause() error {
	return e.cause
}
func (e ReservationRequestValidationError) Key() bool {
	return e.key
}
func (e ReservationRequestValidationError) ErrorName() string {
	return "ReservationRequestValidationError"
}
func (e ReservationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}
	key := ""
	if e.key {
		key = "key for "
	}
	return fmt.Sprintf("invalid %sReservationRequest.%s: %s%s", key, e.field, e.reason, cause)
}

var _ error = ReservationRequestValidationError{}
var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReservationRequestValidationError{}

func (m *Reservation) Validate() error {
	return m.validate(false)
}
func (m *Reservation) ValidateAll() error {
	return m.validate(true)
}
func (m *Reservation) validate(all bool) error {
	if m == nil {
		return nil
	}
	var errors []error
	if len(errors) > 0 {
		return ReservationMultiError(errors)
	}
	return nil
}

type ReservationMultiError []error

func (m ReservationMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}
func (m ReservationMultiError) AllErrors() []error {
	return m
}

type ReservationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

func (e ReservationValidationError) Field() string {
	return e.field
}
func (e ReservationValidationError) Reason() string {
	return e.reason
}
func (e ReservationValidationError) Cause() error {
	return e.cause
}
func (e ReservationValidationError) Key() bool {
	return e.key
}
func (e ReservationValidationError) ErrorName() string {
	return "ReservationValidationError"
}
func (e ReservationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}
	key := ""
	if e.key {
		key = "key for "
	}
	return fmt.Sprintf("invalid %sReservation.%s: %s%s", key, e.field, e.reason, cause)
}

var _ error = ReservationValidationError{}
var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReservationValidationError{}
