package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BullseyeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BullseyeDud struct { 
    

}

// Bullseye
type Bullseye struct { 
    // Rings - The bullseye rings configured for this queue.
    Rings []Ring `json:"rings"`

}

// String returns a JSON representation of the model
func (o *Bullseye) String() string {
     o.Rings = []Ring{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bullseye) MarshalJSON() ([]byte, error) {
    type Alias Bullseye

    if BullseyeMarshalled {
        return []byte("{}"), nil
    }
    BullseyeMarshalled = true

    return json.Marshal(&struct {
        
        Rings []Ring `json:"rings"`
        *Alias
    }{

        
        Rings: []Ring{{}},
        

        Alias: (*Alias)(u),
    })
}

