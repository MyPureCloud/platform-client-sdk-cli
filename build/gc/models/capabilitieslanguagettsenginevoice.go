package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitieslanguagettsenginevoiceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitieslanguagettsenginevoiceDud struct { 
    

}

// Capabilitieslanguagettsenginevoice
type Capabilitieslanguagettsenginevoice struct { 
    // IsDefault
    IsDefault bool `json:"isDefault"`

}

// String returns a JSON representation of the model
func (o *Capabilitieslanguagettsenginevoice) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitieslanguagettsenginevoice) MarshalJSON() ([]byte, error) {
    type Alias Capabilitieslanguagettsenginevoice

    if CapabilitieslanguagettsenginevoiceMarshalled {
        return []byte("{}"), nil
    }
    CapabilitieslanguagettsenginevoiceMarshalled = true

    return json.Marshal(&struct {
        
        IsDefault bool `json:"isDefault"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

