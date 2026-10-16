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
	t.Run("walk", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{"Chris", 23},
				[]string{"Chris"},
			},
			{
				"nested fields",
				Person{
					"Chris",
					Profile{23, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"pointers to things",
				&Person{
					"Chris",
					Profile{23, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"slices",
				[]Profile{
					{23, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
			{
				"arrays",
				[2]Profile{
					{23, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
			{
				"maps",
				map[string]string{
					"Chris": "London",
					"John":  "Reykjavík",
				},
				[]string{"London", "Reykjavík"},
			},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				var got []string
				walk(c.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, c.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, c.ExpectedCalls)
				}
			})
		}
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Chris": "London",
			"John":  "Reykjavík",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertLength(t, got, len(aMap))
		assertContains(t, got, "London")
		assertContains(t, got, "Reykjavík")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{23, "London"}
			aChannel <- Profile{34, "Reykjavík"}
			close(aChannel)
		}()

		var got []string
		want := []string{"London", "Reykjavík"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with func", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{Age: 23, City: "London"}, Profile{Age: 34, City: "Reykjavík"}
		}

		var got []string
		want := []string{"London", "Reykjavík"}
		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with cycle pointers", func(t *testing.T) {
		type Person struct {
			Name   string
			Friend *Person
		}
		p := Person{Name: "Chris"}
		p.Friend = &p

		var got []string
		want := []string{"Chris"}
		walk(&p, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	var contains bool
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

func assertLength(t testing.TB, got []string, want int) {
	t.Helper()
	if len(got) != want {
		t.Errorf("got %d values but expected %d", len(got), want)
	}
}
