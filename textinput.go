package azuki

import "time"

type TextInputComponent struct {
	BaseComponent
	Value       TextSource `json:"initialValue,omitzero"`
	DebounceMS  uint       `json:"debounce,omitzero"`
	Target      string     `json:"target,omitzero"`
	Name        string     `json:"name"`
	Placeholder TextSource `json:"placeholder,omitzero"`
	HiddenInput bool       `json:"hiddenInput,omitzero"`
}

func TextInput() TextInputComponent {
	return TextInputComponent{
		BaseComponent: newBaseComponent(ComponentTypeTextInput),
	}
}

func (ti TextInputComponent) WithName(k string) TextInputComponent {
	ti.Name = k
	return ti
}

func (ti TextInputComponent) WithValue(t TextSource) TextInputComponent {
	ti.Value = t
	return ti
}

func (ti TextInputComponent) WithDebounce(t time.Duration) TextInputComponent {
	ti.DebounceMS = uint(t.Milliseconds())
	return ti
}

func (ti TextInputComponent) WithTarget(t string) TextInputComponent {
	ti.Target = t
	return ti
}

func (ti TextInputComponent) WithPlaceholder(t TextSource) TextInputComponent {
	ti.Placeholder = t
	return ti
}

func (ti TextInputComponent) WithHiddenInput() TextInputComponent {
	ti.HiddenInput = true
	return ti
}
