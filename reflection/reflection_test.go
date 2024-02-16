package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{

			"strict	with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Nishil", 12},
			[]string{"Nishil"},
		},
		{
			"nested structs",
			Person{
				"Nsihil",
				Profile{33, "London"},
			},
			[]string{"Nishi", "asdas"},
		},
		{
			"pointers to thing",
			&Person{
				"chirs",
				Profile{33, "Nishil"},
			},
			[]string{"Nihil"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
