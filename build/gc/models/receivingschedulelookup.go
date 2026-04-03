package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReceivingschedulelookupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReceivingschedulelookupDud struct { 
    

}

// Receivingschedulelookup
type Receivingschedulelookup struct { 
    // Id - The ID of the schedule
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Receivingschedulelookup) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Receivingschedulelookup) MarshalJSON() ([]byte, error) {
    type Alias Receivingschedulelookup

    if ReceivingschedulelookupMarshalled {
        return []byte("{}"), nil
    }
    ReceivingschedulelookupMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

