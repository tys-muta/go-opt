package opt

import (
	"fmt"
	"testing"
)

func TestReflectInt64(t *testing.T) {
	options := &TestOptions{}

	var want int64 = 1
	if err := Reflect(options, WithTestOptionInt64(0)); err == nil {
		t.Errorf("failed to validate: int64")
	}
	if err := Reflect(options, WithTestOptionInt64(want)); err != nil {
		t.Errorf("failed to reflect: int64")
	}
	if options.OptionInt64 == nil {
		t.Errorf("option is nil: int64")
	}
	if int64(*options.OptionInt64) != want {
		t.Errorf("option is not match: int64")
	}
}

func TestReflectString(t *testing.T) {
	options := &TestOptions{}

	var want string = "foo"
	if err := Reflect(options, WithTestOptionString("")); err == nil {
		t.Errorf("failed to validate: string")
	}
	if err := Reflect(options, WithTestOptionString(want)); err != nil {
		t.Errorf("failed to reflect: string")
	}
	if options.OptionString == nil {
		t.Errorf("option is nil: string")
	}
	if string(*options.OptionString) != want {
		t.Errorf("option is not match: string")
	}
}

type TestOptions struct {
	OptionInt64  *testOptionInt64
	OptionString *testOptionString
}

type testOptionInt64 int64

func WithTestOptionInt64(v int64) testOptionInt64 {
	return testOptionInt64(v)
}

func (o testOptionInt64) Validate() error {
	if o < 1 {
		return fmt.Errorf("out of range")
	}
	return nil
}

func (o testOptionInt64) Apply(options interface{}) {
	if v, ok := options.(*TestOptions); ok {
		v.OptionInt64 = &o
	}
}

type testOptionString string

func WithTestOptionString(v string) testOptionString {
	return testOptionString(v)
}

func (o testOptionString) Validate() error {
	if o == "" {
		return fmt.Errorf("empty")
	}
	return nil
}

func (o testOptionString) Apply(options interface{}) {
	if v, ok := options.(*TestOptions); ok {
		v.OptionString = &o
	}
}
