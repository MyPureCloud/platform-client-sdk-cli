package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidgroupsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidgroupsummaryDud struct { 
    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Schedulebidgroupsummary
type Schedulebidgroupsummary struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name - The name assigned to this bid group
    Name string `json:"name"`


    // ManagementUnit - The management unit to which this bid group belongs
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // AgentCount - The number of agents in this bid group
    AgentCount int `json:"agentCount"`


    // WorkPlanCount - The number of work plans in this bid group or the number of work plans in rotations
    WorkPlanCount int `json:"workPlanCount"`


    // WorkPlanRotationCount - The number of work plan rotations used in this bid group
    WorkPlanRotationCount int `json:"workPlanRotationCount"`


    // PlanningGroupCount - The number of planning groups in this bid group
    PlanningGroupCount int `json:"planningGroupCount"`


    // ScheduleSetError - Schedule set optimization error details for this bid group. Present only when optimization fails
    ScheduleSetError Scheduleseterror `json:"scheduleSetError"`


    

}

// String returns a JSON representation of the model
func (o *Schedulebidgroupsummary) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidgroupsummary) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidgroupsummary

    if SchedulebidgroupsummaryMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidgroupsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        AgentCount int `json:"agentCount"`
        
        WorkPlanCount int `json:"workPlanCount"`
        
        WorkPlanRotationCount int `json:"workPlanRotationCount"`
        
        PlanningGroupCount int `json:"planningGroupCount"`
        
        ScheduleSetError Scheduleseterror `json:"scheduleSetError"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

