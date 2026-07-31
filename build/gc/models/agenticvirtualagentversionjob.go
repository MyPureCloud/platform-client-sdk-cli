package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentversionjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentversionjobDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    Errors []Errorbody `json:"errors"`


    TokenCount int `json:"tokenCount"`


    VirtualAgentVersion Agenticvirtualagentversion `json:"virtualAgentVersion"`


    SelfUri string `json:"selfUri"`

}

// Agenticvirtualagentversionjob
type Agenticvirtualagentversionjob struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentversionjob) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentversionjob) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentversionjob

    if AgenticvirtualagentversionjobMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentversionjobMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

