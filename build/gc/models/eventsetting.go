package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventsettingDud struct { 
    

}

// Eventsetting
type Eventsetting struct { 
    // Typing - Settings regarding typing events
    Typing Typingsetting `json:"typing"`

}

// String returns a JSON representation of the model
func (o *Eventsetting) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventsetting) MarshalJSON() ([]byte, error) {
    type Alias Eventsetting

    if EventsettingMarshalled {
        return []byte("{}"), nil
    }
    EventsettingMarshalled = true

    return json.Marshal(&struct {
        
        Typing Typingsetting `json:"typing"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

