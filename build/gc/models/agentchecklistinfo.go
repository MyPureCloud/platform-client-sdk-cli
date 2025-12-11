package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentchecklistinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentchecklistinfoDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentchecklistinfo
type Agentchecklistinfo struct { 
    // Id - ID of the checklist.
    Id string `json:"id"`


    // Name - Name of the checklist.
    Name string `json:"name"`


    // ChecklistItems - List of individual Checklist Items.
    ChecklistItems []Checklistitem `json:"checklistItems"`


    

}

// String returns a JSON representation of the model
func (o *Agentchecklistinfo) String() string {
    
    
     o.ChecklistItems = []Checklistitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentchecklistinfo) MarshalJSON() ([]byte, error) {
    type Alias Agentchecklistinfo

    if AgentchecklistinfoMarshalled {
        return []byte("{}"), nil
    }
    AgentchecklistinfoMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ChecklistItems []Checklistitem `json:"checklistItems"`
        *Alias
    }{

        


        


        
        ChecklistItems: []Checklistitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

