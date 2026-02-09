package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationdetailcontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationdetailcontainerDud struct { 
    


    

}

// Taskmanagementobservationdetailcontainer
type Taskmanagementobservationdetailcontainer struct { 
    // TypeDetails - Information about worktypes referenced in the results. Present when 'type' is included in the expands parameter.
    TypeDetails []Worktypereference `json:"typeDetails"`


    // AssigneeDetails - Information about assignees referenced in the results. Present when 'assignee' is included in the expands parameter.
    AssigneeDetails []Userreferencewithname `json:"assigneeDetails"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationdetailcontainer) String() string {
     o.TypeDetails = []Worktypereference{{}} 
     o.AssigneeDetails = []Userreferencewithname{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationdetailcontainer) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationdetailcontainer

    if TaskmanagementobservationdetailcontainerMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationdetailcontainerMarshalled = true

    return json.Marshal(&struct {
        
        TypeDetails []Worktypereference `json:"typeDetails"`
        
        AssigneeDetails []Userreferencewithname `json:"assigneeDetails"`
        *Alias
    }{

        
        TypeDetails: []Worktypereference{{}},
        


        
        AssigneeDetails: []Userreferencewithname{{}},
        

        Alias: (*Alias)(u),
    })
}

