package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchalternativeshifttradesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchalternativeshifttradesrequestDud struct { 
    


    

}

// Searchalternativeshifttradesrequest
type Searchalternativeshifttradesrequest struct { 
    // ManagementUnitIds - The list of management unit IDs for this alternative shift trade search. Either managementUnitIds or agentIds is required
    ManagementUnitIds []string `json:"managementUnitIds"`


    // AgentIds - The list of agent IDs for this alternative shift trade search. Either managementUnitIds or agentIds is required
    AgentIds []string `json:"agentIds"`

}

// String returns a JSON representation of the model
func (o *Searchalternativeshifttradesrequest) String() string {
     o.ManagementUnitIds = []string{""} 
     o.AgentIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchalternativeshifttradesrequest) MarshalJSON() ([]byte, error) {
    type Alias Searchalternativeshifttradesrequest

    if SearchalternativeshifttradesrequestMarshalled {
        return []byte("{}"), nil
    }
    SearchalternativeshifttradesrequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        AgentIds []string `json:"agentIds"`
        *Alias
    }{

        
        ManagementUnitIds: []string{""},
        


        
        AgentIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

