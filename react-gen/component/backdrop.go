package components

// <% .ComponentNameUpper %> type representing a <% .ComponentNameUpper %> react component
type <% .ComponentNameUpper %> struct {
	Options
}

// New<% .ComponentNameUpper %> creates a new <% .ComponentNameUpper %> value with the given options
func New<% .ComponentNameUpper %>(opts ...Option) <% .ComponentNameUpper %> {
	b := <% .ComponentNameUpper %>{}

	for _, wrap := range opts {
		wrap(&b.Options)
	}

	return b
}

// Name the name of the React component type
func (b <% .ComponentNameUpper %>) Name() string {
	return "<% .ComponentNameUpper %>"
}

// Deps list of lower level components which the <% .ComponentNameUpper %> component uses
func (b <% .ComponentNameUpper %>) Deps() []Component {
	// <% .ComponentNameUpper %> is a vanilla component and does not have any dependencies
	return nil
}

// Opts returns the options stored within this component
func (b <% .ComponentNameUpper %>) Opts() Options {
	return b.Options
}

// AdditionalFiles any files other than the main jsx/js/css files for the
// component that you want to include from the component's template dir.
func (b <% .ComponentNameUpper %>) AdditionalFiles() []string {
	return nil
}
