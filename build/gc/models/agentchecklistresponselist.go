package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentchecklistresponselistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentchecklistresponselistDud struct { 
    


    


    


    

}

// Agentchecklistresponselist
type Agentchecklistresponselist struct { 
    // Entities
    Entities []Agentchecklistresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Agentchecklistresponselist) String() string {
     o.Entities = []Agentchecklistresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentchecklistresponselist) MarshalJSON() ([]byte, error) {
    type Alias Agentchecklistresponselist

    if AgentchecklistresponselistMarshalled {
        return []byte("{}"), nil
    }
    AgentchecklistresponselistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Agentchecklistresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Agentchecklistresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

