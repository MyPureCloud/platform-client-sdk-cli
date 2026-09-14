package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivitypresencedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivitypresencedefinitionDud struct { 
    


    

}

// Useractivitypresencedefinition
type Useractivitypresencedefinition struct { 
    // Id - The globally unique identifier for the presence definition
    Id string `json:"id"`


    // SystemPresence - The system presence to which this definition maps
    SystemPresence string `json:"systemPresence"`

}

// String returns a JSON representation of the model
func (o *Useractivitypresencedefinition) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivitypresencedefinition) MarshalJSON() ([]byte, error) {
    type Alias Useractivitypresencedefinition

    if UseractivitypresencedefinitionMarshalled {
        return []byte("{}"), nil
    }
    UseractivitypresencedefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SystemPresence string `json:"systemPresence"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

