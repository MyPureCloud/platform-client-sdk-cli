package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PushintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PushintegrationDud struct { 
    


    

}

// Pushintegration
type Pushintegration struct { 
    // Id - The mobile push integration id associated with the deployment
    Id string `json:"id"`


    // Provider - The integration provider associated with the deployment
    Provider string `json:"provider"`

}

// String returns a JSON representation of the model
func (o *Pushintegration) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pushintegration) MarshalJSON() ([]byte, error) {
    type Alias Pushintegration

    if PushintegrationMarshalled {
        return []byte("{}"), nil
    }
    PushintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Provider string `json:"provider"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

