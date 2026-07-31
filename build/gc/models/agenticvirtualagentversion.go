package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentversionDud struct { 
    Version string `json:"version"`


    


    VirtualAgent Addressableentityref `json:"virtualAgent"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    

}

// Agenticvirtualagentversion
type Agenticvirtualagentversion struct { 
    


    // SelfUri
    SelfUri string `json:"selfUri"`


    


    // Status - The current status of the virtual agent version.
    Status string `json:"status"`


    


    


    // Definition - The definition of the virtual agent.
    Definition Agenticvirtualagentversiondefinition `json:"definition"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentversion) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentversion) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentversion

    if AgenticvirtualagentversionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentversionMarshalled = true

    return json.Marshal(&struct {
        
        SelfUri string `json:"selfUri"`
        
        Status string `json:"status"`
        
        Definition Agenticvirtualagentversiondefinition `json:"definition"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

