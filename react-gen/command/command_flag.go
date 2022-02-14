package command

import "github.com/cleanarc/react-gen/pkg/cli/flags"

// <% .ComponentNameLower %>Flag enum type representing all the available options for the
// <% .ComponentNameLower %> command
type <% .ComponentNameLower %>Flag int

const (
	// ComponentName flag used to supply a custom name for a <% .ComponentNameLower %> component
	ComponentName <% .ComponentNameLower %>Flag = iota + 1
)

// Type denotes the type of each cli option
func (<% .ReceiverName %> <% .ComponentNameLower %>Flag) Type() flags.FlagType {
	switch <% .ReceiverName %> {
	case ComponentName:
		return flags.String
	default:
		return 0
	}
}

// Name returns the full name of a given flag as it should be used in the cli
func (<% .ReceiverName %> <% .ComponentNameLower %>Flag) Name() string {
	switch <% .ReceiverName %> {
	case ComponentName:
		return "name"
	default:
		return ""
	}
}

// Aliases returns any aliases defined for a given flag
func (<% .ReceiverName %> <% .ComponentNameLower %>Flag) Aliases() []string {
	switch <% .ReceiverName %> {
	case ComponentName:
		return []string{"n"}
	default:
		return nil
	}
}

// Usage returns the usage message for each flag
func (<% .ReceiverName %> <% .ComponentNameLower %>Flag) Usage() string {
	switch <% .ReceiverName %> {
	case ComponentName:
		return "A custom name for your <% .ComponentNameLower %> component"
	default:
		return ""
	}
}

// Required returns a bool denoting if a given flag is required
func (<% .ReceiverName %> <% .ComponentNameLower %>Flag) Required() bool {
	switch <% .ReceiverName %> {
	case ComponentName:
		return false
	default:
		return false
	}
}

// DefaultValue returns the value assigned to each flag, if none is specified
// by the user
func (<% .ReceiverName %> <% .ComponentNameLower %>Flag) DefaultValue() interface{} {
	switch <% .ReceiverName %> {
	case ComponentName:
		return "<% .ComponentNameUpper %>"
	default:
		return ""
	}
}
