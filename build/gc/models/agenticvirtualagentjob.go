package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentjobDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    Errors []Errorbody `json:"errors"`


    SelfUri string `json:"selfUri"`

}

// Agenticvirtualagentjob
type Agenticvirtualagentjob struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentjob) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentjob) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentjob

    if AgenticvirtualagentjobMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentjobMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

