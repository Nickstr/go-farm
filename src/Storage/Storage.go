package Storage

import (
    "github.com/boltdb/bolt"
    "time"
    "fmt"
    "encoding/json"
)

type Storage struct {
    db  *bolt.DB
}

func OpenDatabase() *Storage {
    s := new(Storage)
    s.db, _ = bolt.Open("tillage.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
    return s
}

func (s *Storage) createBucketIfNotExists(name string) {
    s.db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte(name))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        return nil
    })
}

func (s *Storage) StoreEvent(farm string, id string, event []byte) {
    s.createBucketIfNotExists(farm)
    s.db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(farm))
        return b.Put([]byte(id), event)
    })
}

func (s *Storage) GetEventsForAggregate(farm string) (e Events) {
    e = make(map[string]Event)
    s.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(farm))
        b.ForEach(func(k, v []byte) error {
            var event Event
            err := json.Unmarshal(v, &event)
            if err == nil {
                e["test"] = event
            }
            return nil
        })
        return nil
    })

    return e
}
