package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestlistingDud struct { 
    

}

// Timeoffrequestlisting
type Timeoffrequestlisting struct { 
    // Entities - List of time off requests
    Entities []Timeoffrequest `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestlisting) String() string {
     o.Entities = []Timeoffrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestlisting) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestlisting

    if TimeoffrequestlistingMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Timeoffrequest `json:"entities"`
        *Alias
    }{

        
        Entities: []Timeoffrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

