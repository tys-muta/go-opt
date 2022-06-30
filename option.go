package opt

import "fmt"

type Option interface {
	Validate() error
	Apply(any)
}

func Reflect(dst any, src ...Option) error {
	for _, option := range src {
		if option == nil {
			continue
		}
		if err := option.Validate(); err != nil {
			return fmt.Errorf("invalid: %w", err)
		}
		option.Apply(dst)
	}
	return nil
}

func Options[T Option](options ...T) []Option {
	dst := []Option{}
	for _, option := range options {
		dst = append(dst, option)
	}
	return dst
}
