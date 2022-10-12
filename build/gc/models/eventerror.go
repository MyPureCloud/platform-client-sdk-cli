package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventerrorDud struct { 
    


    


    

}

// Eventerror
type Eventerror struct { 
    // EventId - The eventId (V4 UUID) for the event that encountered an error.
    EventId string `json:"eventId"`


    // Message - A message describing the error.
    Message string `json:"message"`


    // Retryable - The event for this eventId can be resubmitted if this value is true.
    Retryable bool `json:"retryable"`

}

// String returns a JSON representation of the model
func (o *Eventerror) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventerror) MarshalJSON() ([]byte, error) {
    type Alias Eventerror

    if EventerrorMarshalled {
        return []byte("{}"), nil
    }
    EventerrorMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        Message string `json:"message"`
        
        Retryable bool `json:"retryable"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

