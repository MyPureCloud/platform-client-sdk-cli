package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PresencedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PresencedefinitionDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Presencedefinition
type Presencedefinition struct { 
    // Id - description
    Id string `json:"id"`


    // SystemPresence
    SystemPresence string `json:"systemPresence"`


    

}

// String returns a JSON representation of the model
func (o *Presencedefinition) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Presencedefinition) MarshalJSON() ([]byte, error) {
    type Alias Presencedefinition

    if PresencedefinitionMarshalled {
        return []byte("{}"), nil
    }
    PresencedefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SystemPresence string `json:"systemPresence"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

