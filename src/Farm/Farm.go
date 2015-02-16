package Farm

import (
    "Storage"
)

type Farm struct {
    id          string
    name        string
}

func (state *Farm) transition(event Storage.Event) {
    switch {
        case event.Type == "FarmCreated":
            state.name = event.Data["name"].(string)
    }
}

func CreateFromHistory(events Storage.Events) *Farm {
    state := &Farm{}
    for _, event := range events {
        state.transition(event)
    }
    return state
}
