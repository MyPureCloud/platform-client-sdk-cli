package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationactionDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Integrationaction
type Integrationaction struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Integrationaction) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationaction) MarshalJSON() ([]byte, error) {
    type Alias Integrationaction

    if IntegrationactionMarshalled {
        return []byte("{}"), nil
    }
    IntegrationactionMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

