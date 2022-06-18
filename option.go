package option

import "fmt"

type Option interface {
	Validate() error
	Apply(interface{})
}

func Reflect(dst interface{}, src ...Option) error {
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
