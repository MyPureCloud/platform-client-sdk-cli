package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentassignedshiftsetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentassignedshiftsetDud struct { 
    


    


    


    

}

// Agentassignedshiftset
type Agentassignedshiftset struct { 
    // Id - The ID of the shift set
    Id string `json:"id"`


    // EffectiveWorkPlan - The work plan or work plan rotation used for generating the shift set
    EffectiveWorkPlan Shiftseteffectiveworkplan `json:"effectiveWorkPlan"`


    // Shifts - The scheduled shifts
    Shifts []Schedulebidscheduledshift `json:"shifts"`


    // Agents - The details of the agents assigned to this shift set
    Agents []Assignedagentdetails `json:"agents"`

}

// String returns a JSON representation of the model
func (o *Agentassignedshiftset) String() string {
    
    
     o.Shifts = []Schedulebidscheduledshift{{}} 
     o.Agents = []Assignedagentdetails{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentassignedshiftset) MarshalJSON() ([]byte, error) {
    type Alias Agentassignedshiftset

    if AgentassignedshiftsetMarshalled {
        return []byte("{}"), nil
    }
    AgentassignedshiftsetMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        EffectiveWorkPlan Shiftseteffectiveworkplan `json:"effectiveWorkPlan"`
        
        Shifts []Schedulebidscheduledshift `json:"shifts"`
        
        Agents []Assignedagentdetails `json:"agents"`
        *Alias
    }{

        


        


        
        Shifts: []Schedulebidscheduledshift{{}},
        


        
        Agents: []Assignedagentdetails{{}},
        

        Alias: (*Alias)(u),
    })
}

