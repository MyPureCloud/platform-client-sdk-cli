package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidgroupsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidgroupsummaryDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workplanbidgroupsummary
type Workplanbidgroupsummary struct { 
    


    // Name - The name assigned to this bid group
    Name string `json:"name"`


    // ManagementUnit - The management unit this bid group belongs to
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // AgentCount - The number of agents in this bid group
    AgentCount int `json:"agentCount"`


    // WorkPlanCount - The number of work plans in this bid group
    WorkPlanCount int `json:"workPlanCount"`


    // PlanningGroupCount - The number of planning groups in this bid group
    PlanningGroupCount int `json:"planningGroupCount"`


    

}

// String returns a JSON representation of the model
func (o *Workplanbidgroupsummary) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidgroupsummary) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidgroupsummary

    if WorkplanbidgroupsummaryMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidgroupsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        AgentCount int `json:"agentCount"`
        
        WorkPlanCount int `json:"workPlanCount"`
        
        PlanningGroupCount int `json:"planningGroupCount"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

