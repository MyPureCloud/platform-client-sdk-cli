package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidgroupcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidgroupcreateDud struct { 
    


    


    


    


    

}

// Workplanbidgroupcreate
type Workplanbidgroupcreate struct { 
    // Name - The name of the work plan bid group
    Name string `json:"name"`


    // ManagementUnitId - The management unit ID this bid group belongs to
    ManagementUnitId string `json:"managementUnitId"`


    // AgentIds - Agent IDs who participate in this bid group
    AgentIds []string `json:"agentIds"`


    // WorkPlans - The list of work plans used in this bid group
    WorkPlans []Bidgroupworkplanrequest `json:"workPlans"`


    // PlanningGroupIds - The planning group IDs selected in this bid group
    PlanningGroupIds []string `json:"planningGroupIds"`

}

// String returns a JSON representation of the model
func (o *Workplanbidgroupcreate) String() string {
    
    
     o.AgentIds = []string{""} 
     o.WorkPlans = []Bidgroupworkplanrequest{{}} 
     o.PlanningGroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidgroupcreate) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidgroupcreate

    if WorkplanbidgroupcreateMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidgroupcreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnitId string `json:"managementUnitId"`
        
        AgentIds []string `json:"agentIds"`
        
        WorkPlans []Bidgroupworkplanrequest `json:"workPlans"`
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        *Alias
    }{

        


        


        
        AgentIds: []string{""},
        


        
        WorkPlans: []Bidgroupworkplanrequest{{}},
        


        
        PlanningGroupIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

