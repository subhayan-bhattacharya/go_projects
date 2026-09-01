package config_test

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"configloader/config"
)

func setValidEnvironment(t *testing.T) {
	t.Helper()
	t.Setenv("PORT", "8080")
	t.Setenv("DATABASE_URL", "postgres://localhost:5432/db")
	t.Setenv("TIMEOUT", "30s")
	t.Setenv("DEBUG", "true")
}

func TestLoadValidConfig(t *testing.T) {
	setValidEnvironment(t)

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Load() returned an unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("Load() returned a nil config without an error")
	}

	if cfg.Port != 8080 {
		t.Errorf("Port = %d, want 8080", cfg.Port)
	}
	if cfg.DatabaseUrl != "postgres://localhost:5432/db" {
		t.Errorf("DatabaseUrl = %q, want %q", cfg.DatabaseUrl, "postgres://localhost:5432/db")
	}
	if cfg.Timeout != 30*time.Second {
		t.Errorf("Timeout = %s, want %s", cfg.Timeout, 30*time.Second)
	}
	if !cfg.Debug {
		t.Error("Debug = false, want true")
	}
}

func TestLoadReportsAllErrors(t *testing.T) {
	setValidEnvironment(t)
	t.Setenv("PORT", "abc")
	t.Setenv("DATABASE_URL", "")

	cfg, err := config.Load()
	if err == nil {
		t.Fatal("Load() returned nil error, want invalid PORT and missing DATABASE_URL errors")
	}
	if cfg != nil {
		t.Errorf("Load() returned config %+v with an error, want nil config", cfg)
	}
	if !errors.Is(err, config.ErrMissingField) {
		t.Error("errors.Is(err, ErrMissingField) = false, want true")
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Error("errors.Is(err, strconv.ErrSyntax) = false, want true")
	}

	var fieldErr *config.FieldError
	if !errors.As(err, &fieldErr) {
		t.Fatal("errors.As(err, *FieldError) = false, want true")
	}
	if fieldErr.Field != "PORT" || fieldErr.Value != "abc" {
		t.Errorf("first FieldError = {Field: %q, Value: %q}, want PORT/abc", fieldErr.Field, fieldErr.Value)
	}

	joined, ok := err.(interface{ Unwrap() []error })
	if !ok {
		t.Fatalf("error type %T does not support multi-error unwrapping", err)
	}
	if got := len(joined.Unwrap()); got != 2 {
		t.Errorf("joined error contains %d errors, want 2", got)
	}
}

func TestLoadInvalidDuration(t *testing.T) {
	setValidEnvironment(t)
	t.Setenv("TIMEOUT", "notaduration")

	cfg, err := config.Load()
	if err == nil {
		t.Fatal("Load() returned nil error for an invalid duration")
	}
	if cfg != nil {
		t.Errorf("Load() returned config %+v with an error, want nil config", cfg)
	}

	var fieldErr *config.FieldError
	if !errors.As(err, &fieldErr) {
		t.Fatal("errors.As(err, *FieldError) = false, want true")
	}
	if fieldErr.Field != "TIMEOUT" || fieldErr.Value != "notaduration" {
		t.Errorf("FieldError = {Field: %q, Value: %q}, want TIMEOUT/notaduration", fieldErr.Field, fieldErr.Value)
	}
	if errors.Is(err, config.ErrMissingField) {
		t.Error("invalid duration was incorrectly classified as a missing field")
	}
}

func TestLoadInvalidBool(t *testing.T) {
	setValidEnvironment(t)
	t.Setenv("DEBUG", "not-a-bool")

	_, err := config.Load()
	if err == nil {
		t.Fatal("Load() returned nil error for an invalid boolean")
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Error("errors.Is(err, strconv.ErrSyntax) = false, want true")
	}

	var fieldErr *config.FieldError
	if !errors.As(err, &fieldErr) {
		t.Fatal("errors.As(err, *FieldError) = false, want true")
	}
	if fieldErr.Field != "DEBUG" || fieldErr.Value != "not-a-bool" {
		t.Errorf("FieldError = {Field: %q, Value: %q}, want DEBUG/not-a-bool", fieldErr.Field, fieldErr.Value)
	}
}
