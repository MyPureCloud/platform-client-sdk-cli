package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperbushorttermforecastreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperbushorttermforecastreferenceDud struct { 
    

}

// Valuewrapperbushorttermforecastreference
type Valuewrapperbushorttermforecastreference struct { 
    // Value - The value for the associated field
    Value Bushorttermforecastreference `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperbushorttermforecastreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperbushorttermforecastreference) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperbushorttermforecastreference

    if ValuewrapperbushorttermforecastreferenceMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperbushorttermforecastreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Value Bushorttermforecastreference `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

