package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitieslanguagettsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitieslanguagettsDud struct { 
    

}

// Capabilitieslanguagetts
type Capabilitieslanguagetts struct { 
    // Engines
    Engines map[string]Capabilitieslanguagettsengine `json:"engines"`

}

// String returns a JSON representation of the model
func (o *Capabilitieslanguagetts) String() string {
     o.Engines = map[string]Capabilitieslanguagettsengine{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitieslanguagetts) MarshalJSON() ([]byte, error) {
    type Alias Capabilitieslanguagetts

    if CapabilitieslanguagettsMarshalled {
        return []byte("{}"), nil
    }
    CapabilitieslanguagettsMarshalled = true

    return json.Marshal(&struct {
        
        Engines map[string]Capabilitieslanguagettsengine `json:"engines"`
        *Alias
    }{

        
        Engines: map[string]Capabilitieslanguagettsengine{"": {}},
        

        Alias: (*Alias)(u),
    })
}

