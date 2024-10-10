package main

import "testing"

func TestCamelCase(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"_my_field_name_2", "XMyFieldName_2"},
		{"my_field_name_2", "MyFieldName_2"},
		{"my_field_name", "MyFieldName"},
		{"_my_field_name", "XMyFieldName"},
		{"singleword", "Singleword"},
		{"_singleWord", "XSingleWord"},
		{"", ""},
	}

	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			if got := camelCase(tt.input); got != tt.want {
				t.Errorf("camelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
