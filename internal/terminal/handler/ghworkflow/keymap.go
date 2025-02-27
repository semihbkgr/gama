package ghworkflow

import (
	teakey "github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	TabSwitch teakey.Binding
}

func (k keyMap) ShortHelp() []teakey.Binding {
	return []teakey.Binding{k.TabSwitch}
}

func (k keyMap) FullHelp() [][]teakey.Binding {
	return [][]teakey.Binding{
		{k.TabSwitch},
	}
}

var keys = keyMap{
	TabSwitch: teakey.NewBinding(
		teakey.WithKeys("shift+left", "shift+right"),
		teakey.WithHelp("shift + (← | →)", "switch tab"),
	),
}

func (m *ModelGithubWorkflow) ViewHelp() string {
	return m.Help.View(m.Keys)
}
