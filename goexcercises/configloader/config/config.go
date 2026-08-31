package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

var ErrMissingField = errors.New("missing required field")

type FieldError struct {
	Field string
	Value string
	Err   error
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("Field %s : invalid value %s: %v", e.Field, e.Value, e.Err)
}

func (e *FieldError) Unwrap() error {
	return e.Err
}

type Config struct {
	Port        int
	DatabaseUrl string
	Timeout     time.Duration
	Debug       bool
}

func requiredBool(loadFunc func(string) (string, bool), field string) (bool, error) {
	var debug bool
	value, isSet := loadFunc(field)
	if !isSet || value == "" {
		return debug, &FieldError{
			Field: field,
			Err:   ErrMissingField,
		}
	}
	parsedDebug, err := strconv.ParseBool(field)
	if err != nil {
		return debug, &FieldError{
			Field: field,
			Value: value,
			Err:   ErrMissingField,
		}
	}
	return parsedDebug, nil
}

func requiredDuration(loadFunc func(string) (string, bool), field string) (time.Duration, error) {
	var duration time.Duration
	value, isSet := loadFunc(field)
	if !isSet || value == "" {
		return duration, &FieldError{
			Field: field,
			Err:   ErrMissingField,
		}
	}
	sec, err := strconv.Atoi(value)
	if err != nil {
		return duration, &FieldError{
			Field: field,
			Value: value,
			Err:   err,
		}
	}
	return time.Duration(sec), nil
}

func requiredString(loadFunc func(string) (string, bool), field string) (string, error) {
	value, isSet := loadFunc(field)
	if !isSet || value == "" {
		return "", &FieldError{
			Field: field,
			Err:   ErrMissingField,
		}
	}
	return value, nil
}

func requiredInt(loadFunc func(string) (string, bool), field string) (int, error) {
	value, isSet := loadFunc(field)
	if !isSet || value == "" {
		return 0, &FieldError{
			Field: field,
			Err:   ErrMissingField,
		}
	}
	converted, err := strconv.Atoi(value)
	if err != nil {
		return 0, &FieldError{
			Field: field,
			Value: value,
			Err:   err,
		}
	}
	return converted, nil
}

func Load() (*Config, error) {
	var errs []error
	cfg := &Config{}
	port, err := requiredInt(os.LookupEnv, "port")
	errs = append(errs, err)
	cfg.Port = port
	databaseUrl, err := requiredString(os.LookupEnv, "databaseUrl")
	errs = append(errs, err)
	cfg.DatabaseUrl = databaseUrl
	duration, err := requiredDuration(os.LookupEnv, "timeout")
	errs = append(errs, err)
	cfg.Timeout = duration
	debug, err := requiredBool(os.LookupEnv, "debug")
	errs = append(errs, err)
	cfg.Debug = debug
	if len(errs) != 0 {
		return nil, errors.Join(errs...)
	}
	return cfg, nil
}
