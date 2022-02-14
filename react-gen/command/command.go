package command

import (
	"errors"

	"github.com/cleanarc/react-gen/pkg/cli/args"
	"github.com/cleanarc/react-gen/pkg/cli/context"
	"github.com/cleanarc/react-gen/pkg/cli/flags"

	"github.com/cleanarc/react-gen/internal/components"
	"github.com/cleanarc/react-gen/internal/generate"
)

var (
	// ErrInvalidArgs error to return when user does not use the
	// cli command properly
	ErrInvalidArgs = errors.New(
		"\nPlease specify the output directory:\n" +
			"  react-gen <% .ComponentNameLower %> <output-directory>\n\n" +
			"For example:\n" +
			"  react-gen <% .ComponentNameLower %> .\n\n" +
			"Run react-gen <% .ComponentNameLower %> --help to see all options.",
	)
)

// <% .ComponentNameUpper %> encapsulates the behaviors of the <% .ComponentNameLower %> subcommand
type <% .ComponentNameUpper %> struct{}

// Name the name of the command
func (<% .ReceiverName %> <% .ComponentNameUpper %>) Name() string {
	return "<% .ComponentNameLower %>"
}

// Usage short description of the command
func (<% .ReceiverName %> <% .ComponentNameUpper %>) Usage() string {
	return "Generates a <% .ComponentNameUpper %> component"
}

// UsageText format to use when using the command
func (<% .ReceiverName %> <% .ComponentNameUpper %>) UsageText() string {
	return "react-gen <% .ComponentNameLower %> [command options] <output-directory>"
}

// Flags the options for the command
func (<% .ReceiverName %> <% .ComponentNameUpper %>) Flags() []flags.Flag {
	return []flags.Flag{
		ComponentName,
	}
}

// Action the action to invoke for the <% .ComponentNameLower %> command
func (<% .ReceiverName %> <% .ComponentNameUpper %>) Action() func(ctx context.Context) error {
	return <% .ReceiverName %>.action
}

// action implementation of Action
func (<% .ReceiverName %> <% .ComponentNameUpper %>) action(ctx context.Context) error {
	componentName := flags.GetStringValue(ctx, ComponentName)

	var outputDir string

	if outputDir = args.GetFirst(ctx); outputDir == "" {
		return ErrInvalidArgs
	}

	return generate.Component(
		components.New<% .ComponentNameUpper %>(),
		componentName,
		outputDir,
	)
}
