package main

import (
    "Farm"
    "Storage"
    "encoding/json"
    "fmt"
)

func main() {
    event := Farm.FarmCreated()
    e, _ := json.Marshal(event)
    s := Storage.OpenDatabase()
    s.StoreEvent("farm", "1", e)


    history := s.GetEventsForAggregate("farm")

    farm := Farm.CreateFromHistory(history)
    fmt.Println(farm)
}
