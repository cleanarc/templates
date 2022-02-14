package command

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cleanarc/react-gen/pkg/cli/flags"
)

func Test<% .ComponentNameUpper %>_Name(t *testing.T) {
	tests := []struct {
		name         string
		<% .ComponentNameLower %>     <% .ComponentNameUpper %>
		expectedName string
	}{
		{
			name:         "name",
			<% .ComponentNameLower %>:     <% .ComponentNameUpper %>{},
			expectedName: "<% .ComponentNameLower %>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedName, test.<% .ComponentNameLower %>.Name())
		})
	}
}

func Test<% .ComponentNameUpper %>_Usage(t *testing.T) {
	tests := []struct {
		name          string
		<% .ComponentNameLower %>      <% .ComponentNameUpper %>
		expectedUsage string
	}{
		{
			name:          "usage",
			<% .ComponentNameLower %>:      <% .ComponentNameUpper %>{},
			expectedUsage: "Generates a <% .ComponentNameUpper %> component",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedUsage, test.<% .ComponentNameLower %>.Usage())
		})
	}
}

func Test<% .ComponentNameUpper %>_UsageText(t *testing.T) {
	tests := []struct {
		name         string
		<% .ComponentNameLower %>     <% .ComponentNameUpper %>
		expectedText string
	}{
		{
			name:         "usage_text",
			<% .ComponentNameLower %>:     <% .ComponentNameUpper %>{},
			expectedText: "react-gen <% .ComponentNameLower %> [command options] <output-directory>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedText, test.<% .ComponentNameLower %>.UsageText())
		})
	}
}

func Test<% .ComponentNameUpper %>_Flags(t *testing.T) {
	tests := []struct {
		name          string
		<% .ComponentNameLower %>      <% .ComponentNameUpper %>
		expectedFlags []flags.Flag
	}{
		{
			name:     "flags",
			<% .ComponentNameLower %>: <% .ComponentNameUpper %>{},
			expectedFlags: []flags.Flag{
				ComponentName,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedFlags, test.<% .ComponentNameLower %>.Flags())
		})
	}
}

func Test<% .ComponentNameUpper %>_Action(t *testing.T) {
	tests := []struct {
		name     string
		<% .ComponentNameLower %> <% .ComponentNameUpper %>
	}{
		{
			name:     "action_set",
			<% .ComponentNameLower %>: <% .ComponentNameUpper %>{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.NotNil(t, test.<% .ComponentNameLower %>.Action())
		})
	}
}

func Test<% .ComponentNameUpper %>Flag_Type(t *testing.T) {
	tests := []struct {
		name         string
		flag         <% .ComponentNameLower %>Flag
		expectedType flags.FlagType
	}{
		{
			name:         "component_name_type",
			flag:         ComponentName,
			expectedType: flags.String,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedType, test.flag.Type())
		})
	}
}

func Test<% .ComponentNameUpper %>Flag_Name(t *testing.T) {
	tests := []struct {
		name         string
		flag         <% .ComponentNameLower %>Flag
		expectedName string
	}{
		{
			name:         "component_name",
			flag:         ComponentName,
			expectedName: "name",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedName, test.flag.Name())
		})
	}
}

func Test<% .ComponentNameUpper %>Flag_Aliases(t *testing.T) {
	tests := []struct {
		name            string
		flag            <% .ComponentNameLower %>Flag
		expectedAliases []string
	}{
		{
			name:            "component_name_aliases",
			flag:            ComponentName,
			expectedAliases: []string{"n"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedAliases, test.flag.Aliases())
		})
	}
}

func Test<% .ComponentNameUpper %>Flag_Usage(t *testing.T) {
	tests := []struct {
		name          string
		flag          <% .ComponentNameLower %>Flag
		expectedUsage string
	}{
		{
			name:          "component_name_usage",
			flag:          ComponentName,
			expectedUsage: "A custom name for your <% .ComponentNameLower %> component",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedUsage, test.flag.Usage())
		})
	}
}

func Test<% .ComponentNameUpper %>Flag_Required(t *testing.T) {
	tests := []struct {
		name             string
		flag             <% .ComponentNameLower %>Flag
		expectedRequired bool
	}{
		{
			name:             "component_name_required",
			flag:             ComponentName,
			expectedRequired: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedRequired, test.flag.Required())
		})
	}
}

func Test<% .ComponentNameUpper %>Flag_DefaultValue(t *testing.T) {
	tests := []struct {
		name                 string
		flag                 <% .ComponentNameLower %>Flag
		expectedDefaultValue string
	}{
		{
			name:                 "component_name_default_value",
			flag:                 ComponentName,
			expectedDefaultValue: "<% .ComponentNameUpper %>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedDefaultValue, test.flag.DefaultValue())
		})
	}
}
