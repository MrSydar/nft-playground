package main

import "fmt"

func generateProduct[T any](consumer func([]T) error, sets ...[]T) error {
	var product func([]T, int) error
	product = func(set []T, pos int) error {
		if pos == len(sets) {
			s := make([]T, len(sets))
			copy(s, set)

			if err := consumer(s); err != nil {
				return fmt.Errorf("consumer error: %v", err)
			}
		} else {
			for _, item := range sets[pos] {
				product(append(set, item), pos+1)
			}
		}

		return nil
	}

	if err := product([]T{}, 0); err != nil {
		return fmt.Errorf("product loop error: %v", err)
	}

	return nil
}
