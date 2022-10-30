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
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Chris", "Tokyo"},
			[]string{"Chris", "Tokyo"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 20},
			[]string{"Chris"},
		},
		{
			"Struct with nested field",
			Person{"Chris", Profile{20, "Tokyo"}},
			[]string{"Chris", "Tokyo"},
		},
		{
			"Struct with pointer",
			&Person{"Chris", Profile{20, "Tokyo"}},
			[]string{"Chris", "Tokyo"},
		},
		{
			"Slices",
			[]Profile{
				{20, "Tokyo"},
				{30, "Osaka"},
			},
			[]string{"Tokyo", "Osaka"},
		},
		{
			"Array",
			[2]Profile{
				{20, "Tokyo"},
				{30, "Osaka"},
			},
			[]string{"Tokyo", "Osaka"},
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

	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo": "bar",
			"baz": "boz",
		}
		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "bar")
		assertContains(t, got, "boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "Tokyo"}
			aChannel <- Profile{30, "Osaka"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Tokyo", "Osaka"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with funcs", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{33, "Tokyo"}, Profile{30, "Osaka"}
		}

		var got []string
		want := []string{"Tokyo", "Osaka"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, heystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range heystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q bud it didn't", heystack, needle)
	}
}
