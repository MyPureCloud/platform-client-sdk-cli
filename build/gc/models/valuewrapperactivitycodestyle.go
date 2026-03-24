package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperactivitycodestyleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperactivitycodestyleDud struct { 
    

}

// Valuewrapperactivitycodestyle
type Valuewrapperactivitycodestyle struct { 
    // Value - The value for the associated field
    Value Activitycodestyle `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperactivitycodestyle) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperactivitycodestyle) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperactivitycodestyle

    if ValuewrapperactivitycodestyleMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperactivitycodestyleMarshalled = true

    return json.Marshal(&struct {
        
        Value Activitycodestyle `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

