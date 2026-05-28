package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperinstantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperinstantDud struct { 
    

}

// Valuewrapperinstant
type Valuewrapperinstant struct { 
    // Value - The value for the associated field. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Value time.Time `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperinstant) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperinstant) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperinstant

    if ValuewrapperinstantMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperinstantMarshalled = true

    return json.Marshal(&struct {
        
        Value time.Time `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

