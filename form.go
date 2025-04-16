package azuki

type FormComponent struct {
	BaseComponent
	Target   string         `json:"target,omitzero"`
	Children ChildrenSource `json:"children"`
}

func Form(target string) FormComponent {
	return FormComponent{
		BaseComponent: newBaseComponent(ComponentTypeForm),
		Target:        target,
	}
}

func (f FormComponent) WithChildren(children ...Component) FormComponent {
	f.Children = ChildrenSlice(children)
	return f
}
