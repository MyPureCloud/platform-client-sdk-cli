package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperactivityplanservicegoalimpactoverridesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperactivityplanservicegoalimpactoverridesDud struct { 
    

}

// Valuewrapperactivityplanservicegoalimpactoverrides
type Valuewrapperactivityplanservicegoalimpactoverrides struct { 
    // Value - The value for the associated field
    Value Activityplanservicegoalimpactoverrides `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperactivityplanservicegoalimpactoverrides) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperactivityplanservicegoalimpactoverrides) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperactivityplanservicegoalimpactoverrides

    if ValuewrapperactivityplanservicegoalimpactoverridesMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperactivityplanservicegoalimpactoverridesMarshalled = true

    return json.Marshal(&struct {
        
        Value Activityplanservicegoalimpactoverrides `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

