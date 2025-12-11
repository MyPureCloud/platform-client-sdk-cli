package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentchecklistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentchecklistDud struct { 
    Id string `json:"id"`


    


    


    


    CreatedBy Userreference `json:"createdBy"`


    ModifiedBy Userreference `json:"modifiedBy"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Agentchecklist
type Agentchecklist struct { 
    


    // Name - Agent Checklist Name.
    Name string `json:"name"`


    // Language - Agent Checklist Language.
    Language string `json:"language"`


    // ChecklistItems - Agent Checklist Items.
    ChecklistItems []Agentchecklistitem `json:"checklistItems"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Agentchecklist) String() string {
    
    
     o.ChecklistItems = []Agentchecklistitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentchecklist) MarshalJSON() ([]byte, error) {
    type Alias Agentchecklist

    if AgentchecklistMarshalled {
        return []byte("{}"), nil
    }
    AgentchecklistMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Language string `json:"language"`
        
        ChecklistItems []Agentchecklistitem `json:"checklistItems"`
        *Alias
    }{

        


        


        


        
        ChecklistItems: []Agentchecklistitem{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

