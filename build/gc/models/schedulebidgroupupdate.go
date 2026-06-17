package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidgroupupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidgroupupdateDud struct { 
    


    


    


    


    


    


    

}

// Schedulebidgroupupdate
type Schedulebidgroupupdate struct { 
    // Name - The name of the schedule bid group
    Name string `json:"name"`


    // ManagementUnitId - The ID of the management unit to which this bid group belongs
    ManagementUnitId string `json:"managementUnitId"`


    // AgentIds - The IDs of the agents who participate in this bid group
    AgentIds Setwrapperstring `json:"agentIds"`


    // WorkPlanIds - The IDs of the work plans used in this bid group
    WorkPlanIds Setwrapperstring `json:"workPlanIds"`


    // WorkPlanRotations - The work plan rotations used in this bid group
    WorkPlanRotations Listwrapperbidgroupworkplanrotationrequest `json:"workPlanRotations"`


    // PlanningGroupIds - The IDs of the planning groups selected in this bid group
    PlanningGroupIds Setwrapperstring `json:"planningGroupIds"`


    // ScheduleSets - The schedule sets generated for this bid group
    ScheduleSets Listwrapperschedulesetrequest `json:"scheduleSets"`

}

// String returns a JSON representation of the model
func (o *Schedulebidgroupupdate) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidgroupupdate) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidgroupupdate

    if SchedulebidgroupupdateMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidgroupupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnitId string `json:"managementUnitId"`
        
        AgentIds Setwrapperstring `json:"agentIds"`
        
        WorkPlanIds Setwrapperstring `json:"workPlanIds"`
        
        WorkPlanRotations Listwrapperbidgroupworkplanrotationrequest `json:"workPlanRotations"`
        
        PlanningGroupIds Setwrapperstring `json:"planningGroupIds"`
        
        ScheduleSets Listwrapperschedulesetrequest `json:"scheduleSets"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

