package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediareactionsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediareactionsrequestDud struct { 
    

}

// Opensocialmediareactionsrequest
type Opensocialmediareactionsrequest struct { 
    // Events - List of open social media reaction events
    Events []Opensocialmediareactionsnormalizedevent `json:"events"`

}

// String returns a JSON representation of the model
func (o *Opensocialmediareactionsrequest) String() string {
     o.Events = []Opensocialmediareactionsnormalizedevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediareactionsrequest) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediareactionsrequest

    if OpensocialmediareactionsrequestMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediareactionsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Events []Opensocialmediareactionsnormalizedevent `json:"events"`
        *Alias
    }{

        
        Events: []Opensocialmediareactionsnormalizedevent{{}},
        

        Alias: (*Alias)(u),
    })
}

