package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeofflineconfigurationnetworkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeofflineconfigurationnetworkDud struct { 
    

}

// Edgeofflineconfigurationnetwork
type Edgeofflineconfigurationnetwork struct { 
    // Wan - Settings for the Edge WAN interface.
    Wan Edgeofflineconfigurationinterface `json:"wan"`

}

// String returns a JSON representation of the model
func (o *Edgeofflineconfigurationnetwork) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeofflineconfigurationnetwork) MarshalJSON() ([]byte, error) {
    type Alias Edgeofflineconfigurationnetwork

    if EdgeofflineconfigurationnetworkMarshalled {
        return []byte("{}"), nil
    }
    EdgeofflineconfigurationnetworkMarshalled = true

    return json.Marshal(&struct {
        
        Wan Edgeofflineconfigurationinterface `json:"wan"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

