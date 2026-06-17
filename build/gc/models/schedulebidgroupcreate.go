package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidgroupcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidgroupcreateDud struct { 
    


    


    


    


    


    

}

// Schedulebidgroupcreate
type Schedulebidgroupcreate struct { 
    // Name - The name of the schedule bid group
    Name string `json:"name"`


    // ManagementUnitId - The ID of the management unit to which this bid group belongs
    ManagementUnitId string `json:"managementUnitId"`


    // AgentIds - The IDs of the agents who participate in this bid group
    AgentIds []string `json:"agentIds"`


    // WorkPlanIds - The IDs of the work plans used in this bid group
    WorkPlanIds []string `json:"workPlanIds"`


    // WorkPlanRotations - The work plan rotations used in this bid group
    WorkPlanRotations []Bidgroupworkplanrotationrequest `json:"workPlanRotations"`


    // PlanningGroupIds - The IDs of the planning groups selected in this bid group
    PlanningGroupIds []string `json:"planningGroupIds"`

}

// String returns a JSON representation of the model
func (o *Schedulebidgroupcreate) String() string {
    
    
     o.AgentIds = []string{""} 
     o.WorkPlanIds = []string{""} 
     o.WorkPlanRotations = []Bidgroupworkplanrotationrequest{{}} 
     o.PlanningGroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidgroupcreate) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidgroupcreate

    if SchedulebidgroupcreateMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidgroupcreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnitId string `json:"managementUnitId"`
        
        AgentIds []string `json:"agentIds"`
        
        WorkPlanIds []string `json:"workPlanIds"`
        
        WorkPlanRotations []Bidgroupworkplanrotationrequest `json:"workPlanRotations"`
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        *Alias
    }{

        


        


        
        AgentIds: []string{""},
        


        
        WorkPlanIds: []string{""},
        


        
        WorkPlanRotations: []Bidgroupworkplanrotationrequest{{}},
        


        
        PlanningGroupIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

