package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentchecklistitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentchecklistitemDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentchecklistitem
type Agentchecklistitem struct { 
    


    // Name - Name of the checklist item.
    Name string `json:"name"`


    // Description - Description of the checklist item.
    Description string `json:"description"`


    // AutomatedCheckEnabled - Flag to indicate whether automated check is enabled for this checklist item.
    AutomatedCheckEnabled bool `json:"automatedCheckEnabled"`


    // Important - Flag to indicate whether this checklist item is marked as important.
    Important bool `json:"important"`


    

}

// String returns a JSON representation of the model
func (o *Agentchecklistitem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentchecklistitem) MarshalJSON() ([]byte, error) {
    type Alias Agentchecklistitem

    if AgentchecklistitemMarshalled {
        return []byte("{}"), nil
    }
    AgentchecklistitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        AutomatedCheckEnabled bool `json:"automatedCheckEnabled"`
        
        Important bool `json:"important"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

