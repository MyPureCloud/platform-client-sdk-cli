package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagenttargetreferencedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagenttargetreferencedefinitionDud struct { 
    


    


    

}

// Agenticvirtualagenttargetreferencedefinition - Target reference definition for input or output data type properties
type Agenticvirtualagenttargetreferencedefinition struct { 
    // Id - The ID of the target Custom Conversation Attribute schema
    Id string `json:"id"`


    // Name - The name of the target attribute field on the Custom Conversation Attribute schema
    Name string `json:"name"`


    // SelfUri - The URI for this resource.
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttargetreferencedefinition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagenttargetreferencedefinition) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagenttargetreferencedefinition

    if AgenticvirtualagenttargetreferencedefinitionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagenttargetreferencedefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

