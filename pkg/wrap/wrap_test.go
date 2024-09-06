package wrap

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			fn1 := func() error {
				fn := func(int) error { return nil }
				err := fn(1)
				if err != nil {
					return err
				}

				return nil
			}

			w := NewW()
			w.Run(fn1)

		})
	}
}
