package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew<% .ComponentNameUpper %>(t *testing.T) {
	tests := []struct {
		name             string
		<% .ComponentNameLower %>         <% .ComponentNameUpper %>
		expected<% .ComponentNameUpper %> <% .ComponentNameUpper %>
	}{
		{
			name:     "new",
			<% .ComponentNameLower %>: New<% .ComponentNameUpper %>(),
			expected<% .ComponentNameUpper %>: <% .ComponentNameUpper %>{
				Options: Options{},
			},
		},
		{
			name:     "new_with_opts",
			<% .ComponentNameLower %>: New<% .ComponentNameUpper %>(WithDepsInDir()),
			expected<% .ComponentNameUpper %>: <% .ComponentNameUpper %>{
				Options: Options{
					DepsInDir: true,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected<% .ComponentNameUpper %>, test.<% .ComponentNameLower %>)
		})
	}
}

func Test<% .ComponentNameUpper %>_Name(t *testing.T) {
	tests := []struct {
		name         string
		<% .ComponentNameLower %>     <% .ComponentNameUpper %>
		expectedName string
	}{
		{
			name:         "name",
			<% .ComponentNameLower %>:     <% .ComponentNameUpper %>{},
			expectedName: "<% .ComponentNameUpper %>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedName, test.<% .ComponentNameLower %>.Name())
		})
	}
}

func Test<% .ComponentNameUpper %>_Deps(t *testing.T) {
	tests := []struct {
		name     string
		<% .ComponentNameLower %> <% .ComponentNameUpper %>
	}{
		{
			name:     "deps",
			<% .ComponentNameLower %>: <% .ComponentNameUpper %>{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Nil(t, test.<% .ComponentNameLower %>.Deps())
		})
	}
}

func Test<% .ComponentNameUpper %>_Opts(t *testing.T) {
	tests := []struct {
		name         string
		<% .ComponentNameLower %>     <% .ComponentNameUpper %>
		expectedOpts Options
	}{
		{
			name:     "opts",
			<% .ComponentNameLower %>: New<% .ComponentNameUpper %>(WithDepsInDir()),
			expectedOpts: Options{
				DepsInDir: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedOpts, test.<% .ComponentNameLower %>.Opts())
		})
	}
}

func Test<% .ComponentNameUpper %>_AdditionalFiles(t *testing.T) {
	tests := []struct {
		name     string
		<% .ComponentNameLower %> <% .ComponentNameUpper %>
	}{
		{
			name:     "additional_files",
			<% .ComponentNameLower %>: New<% .ComponentNameUpper %>(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Nil(t, test.<% .ComponentNameLower %>.AdditionalFiles())
		})
	}
}
