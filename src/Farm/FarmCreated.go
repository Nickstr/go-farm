package Farm
import "Storage"
func FarmCreated() Storage.Event {
    properties := make(map[string]interface{})
    properties["name"] = "test4343"

    return Storage.Event {
        Id: "331223",
        Type: "FarmCreated",
        Data: properties,
    }
}
