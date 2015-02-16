package Storage

type Event struct {
    Id          string                  `json:"id"`
    Type        string                  `json:"name"`
    Data        map[string]interface{}  `json:"properties"`
}

type Events map[string]Event
