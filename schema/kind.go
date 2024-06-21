package schema

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"
)

// Kind represents the basic type of a field in an object.
// Each kind defines the types of go values which should be accepted
// by listeners and generated by decoders when providing entity updates.
type Kind int

const (
	// InvalidKind indicates that an invalid type.
	InvalidKind Kind = iota

	// StringKind is a string type and values of this type must be of the go type string
	// containing valid UTF-8.
	StringKind

	// BytesKind is a bytes type and values of this type must be of the go type []byte.
	BytesKind

	// Int8Kind is an int8 type and values of this type must be of the go type int8.
	Int8Kind

	// Uint8Kind is a uint8 type and values of this type must be of the go type uint8.
	Uint8Kind

	// Int16Kind is an int16 type and values of this type must be of the go type int16.
	Int16Kind

	// Uint16Kind is a uint16 type and values of this type must be of the go type uint16.
	Uint16Kind

	// Int32Kind is an int32 type and values of this type must be of the go type int32.
	Int32Kind

	// Uint32Kind is a uint32 type and values of this type must be of the go type uint32.
	Uint32Kind

	// Int64Kind is an int64 type and values of this type must be of the go type int64.
	Int64Kind

	// Uint64Kind is a uint64 type and values of this type must be of the go type uint64.
	Uint64Kind

	// IntegerStringKind represents an arbitrary precision integer number. Values of this type must
	// be of the go type string and formatted as base10 integers, specifically matching to
	// the IntegerFormat regex.
	IntegerStringKind

	// DecimalStringKind represents an arbitrary precision decimal or integer number. Values of this type
	// must be of the go type string and match the DecimalFormat regex.
	DecimalStringKind

	// BoolKind is a boolean type and values of this type must be of the go type bool.
	BoolKind

	// TimeKind is a time type and values of this type must be of the go type time.Time.
	TimeKind

	// DurationKind is a duration type and values of this type must be of the go type time.Duration.
	DurationKind

	// Float32Kind is a float32 type and values of this type must be of the go type float32.
	Float32Kind

	// Float64Kind is a float64 type and values of this type must be of the go type float64.
	Float64Kind

	// Bech32AddressKind is a bech32 address type and values of this type must be of the go type []byte.
	// Fields of this type are expected to set the AddressPrefix field in the field definition to the
	// bech32 address prefix so that indexers can properly convert them to strings.
	Bech32AddressKind

	// EnumKind is an enum type and values of this type must be of the go type string.
	// Fields of this type are expected to set the EnumDefinition field in the field definition to the enum
	// definition.
	EnumKind

	// JSONKind is a JSON type and values of this type should be of go type json.RawMessage and represent
	// valid JSON.
	JSONKind
)

// MAX_VALID_KIND is the maximum valid kind value.
const MAX_VALID_KIND = JSONKind

const (
	// IntegerFormat is a regex that describes the format integer number strings must match. It specifies
	// that integers may have at most 100 digits.
	IntegerFormat = `^-?[0-9]{1,100}$`

	// DecimalFormat is a regex that describes the format decimal number strings must match. It specifies
	// that decimals may have at most 50 digits before and after the decimal point and may have an optional
	// exponent of up to 2 digits. These restrictions ensure that the decimal can be accurately represented
	// by a wide variety of implementations.
	DecimalFormat = `^-?[0-9]{1,50}(\.[0-9]{1,50})?([eE][-+]?[0-9]{1,2})?$`
)

// Validate returns an errContains if the kind is invalid.
func (t Kind) Validate() error {
	if t <= InvalidKind {
		return fmt.Errorf("unknown type: %d", t)
	}
	if t > JSONKind {
		return fmt.Errorf("invalid type: %d", t)
	}
	return nil
}

// String returns a string representation of the kind.
func (t Kind) String() string {
	switch t {
	case StringKind:
		return "string"
	case BytesKind:
		return "bytes"
	case Int8Kind:
		return "int8"
	case Uint8Kind:
		return "uint8"
	case Int16Kind:
		return "int16"
	case Uint16Kind:
		return "uint16"
	case Int32Kind:
		return "int32"
	case Uint32Kind:
		return "uint32"
	case Int64Kind:
		return "int64"
	case Uint64Kind:
		return "uint64"
	case DecimalStringKind:
		return "decimal"
	case IntegerStringKind:
		return "integer"
	case BoolKind:
		return "bool"
	case TimeKind:
		return "time"
	case DurationKind:
		return "duration"
	case Float32Kind:
		return "float32"
	case Float64Kind:
		return "float64"
	case Bech32AddressKind:
		return "bech32address"
	case EnumKind:
		return "enum"
	case JSONKind:
		return "json"
	default:
		return fmt.Sprintf("invalid(%d)", t)
	}
}

// ValidateValueType returns an errContains if the value does not conform to the expected go type.
// Some fields may accept nil values, however, this method does not have any notion of
// nullability. This method only validates that the go type of the value is correct for the kind
// and does not validate string or json formats. Kind.ValidateValue does a more thorough validation
// of number and json string formatting.
func (t Kind) ValidateValueType(value interface{}) error {
	switch t {
	case StringKind:
		_, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
	case BytesKind:
		_, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("expected []byte, got %T", value)
		}
	case Int8Kind:
		_, ok := value.(int8)
		if !ok {
			return fmt.Errorf("expected int8, got %T", value)
		}
	case Uint8Kind:
		_, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("expected uint8, got %T", value)
		}
	case Int16Kind:
		_, ok := value.(int16)
		if !ok {
			return fmt.Errorf("expected int16, got %T", value)
		}
	case Uint16Kind:
		_, ok := value.(uint16)
		if !ok {
			return fmt.Errorf("expected uint16, got %T", value)
		}
	case Int32Kind:
		_, ok := value.(int32)
		if !ok {
			return fmt.Errorf("expected int32, got %T", value)
		}
	case Uint32Kind:
		_, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("expected uint32, got %T", value)
		}
	case Int64Kind:
		_, ok := value.(int64)
		if !ok {
			return fmt.Errorf("expected int64, got %T", value)
		}
	case Uint64Kind:
		_, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("expected uint64, got %T", value)
		}
	case IntegerStringKind:
		_, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}

	case DecimalStringKind:
		_, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
	case BoolKind:
		_, ok := value.(bool)
		if !ok {
			return fmt.Errorf("expected bool, got %T", value)
		}
	case TimeKind:
		_, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("expected time.Time, got %T", value)
		}
	case DurationKind:
		_, ok := value.(time.Duration)
		if !ok {
			return fmt.Errorf("expected time.Duration, got %T", value)
		}
	case Float32Kind:
		_, ok := value.(float32)
		if !ok {
			return fmt.Errorf("expected float32, got %T", value)
		}
	case Float64Kind:
		_, ok := value.(float64)
		if !ok {
			return fmt.Errorf("expected float64, got %T", value)
		}
	case Bech32AddressKind:
		_, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("expected []byte, got %T", value)
		}
	case EnumKind:
		_, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
	case JSONKind:
		_, ok := value.(json.RawMessage)
		if !ok {
			return fmt.Errorf("expected json.RawMessage, got %T", value)
		}
	default:
		return fmt.Errorf("invalid type: %d", t)
	}
	return nil
}

// ValidateValue returns an errContains if the value does not conform to the expected go type and format.
// It is more thorough, but slower, than Kind.ValidateValueType and validates that Integer, Decimal and JSON
// values are formatted correctly. It cannot validate enum values because Kind's do not have enum schemas.
func (t Kind) ValidateValue(value interface{}) error {
	err := t.ValidateValueType(value)
	if err != nil {
		return err
	}

	switch t {
	case StringKind:
		if !utf8.ValidString(value.(string)) {
			return fmt.Errorf("expected valid utf-8 string, got %s", value)
		}
	case IntegerStringKind:
		if !integerRegex.Match([]byte(value.(string))) {
			return fmt.Errorf("expected base10 integer, got %s", value)
		}
	case DecimalStringKind:
		if !decimalRegex.Match([]byte(value.(string))) {
			return fmt.Errorf("expected decimal number, got %s", value)
		}
	case JSONKind:
		if !json.Valid(value.(json.RawMessage)) {
			return fmt.Errorf("expected valid JSON, got %s", value)
		}
	default:
		return nil
	}
	return nil
}

var (
	integerRegex = regexp.MustCompile(IntegerFormat)
	decimalRegex = regexp.MustCompile(DecimalFormat)
)

// KindForGoValue finds the simplest kind that can represent the given go value. It will not, however,
// return kinds such as IntegerStringKind, DecimalStringKind, Bech32AddressKind, or EnumKind which all can be
// represented as strings.
func KindForGoValue(value interface{}) Kind {
	switch value.(type) {
	case string:
		return StringKind
	case []byte:
		return BytesKind
	case int8:
		return Int8Kind
	case uint8:
		return Uint8Kind
	case int16:
		return Int16Kind
	case uint16:
		return Uint16Kind
	case int32:
		return Int32Kind
	case uint32:
		return Uint32Kind
	case int64:
		return Int64Kind
	case uint64:
		return Uint64Kind
	case float32:
		return Float32Kind
	case float64:
		return Float64Kind
	case bool:
		return BoolKind
	case time.Time:
		return TimeKind
	case time.Duration:
		return DurationKind
	case json.RawMessage:
		return JSONKind
	default:
		return InvalidKind
	}
}
