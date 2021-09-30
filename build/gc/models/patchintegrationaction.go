package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchintegrationactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchintegrationactionDud struct { 
    

}

// Patchintegrationaction
type Patchintegrationaction struct { 
    // Id - ID of the integration action to be invoked.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Patchintegrationaction) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchintegrationaction) MarshalJSON() ([]byte, error) {
    type Alias Patchintegrationaction

    if PatchintegrationactionMarshalled {
        return []byte("{}"), nil
    }
    PatchintegrationactionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

