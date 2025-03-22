package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperlocaldateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperlocaldateDud struct { 
    

}

// Valuewrapperlocaldate
type Valuewrapperlocaldate struct { 
    // Value - The value for the associated field. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    Value time.Time `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperlocaldate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperlocaldate) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperlocaldate

    if ValuewrapperlocaldateMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperlocaldateMarshalled = true

    return json.Marshal(&struct {
        
        Value time.Time `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

