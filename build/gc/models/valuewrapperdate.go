package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperdateDud struct { 
    

}

// Valuewrapperdate
type Valuewrapperdate struct { 
    // Value - The value for the associated field. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Value time.Time `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperdate) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperdate

    if ValuewrapperdateMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperdateMarshalled = true

    return json.Marshal(&struct {
        
        Value time.Time `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

