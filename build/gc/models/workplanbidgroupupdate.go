package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidgroupupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidgroupupdateDud struct { 
    


    


    


    


    

}

// Workplanbidgroupupdate
type Workplanbidgroupupdate struct { 
    // Name - The name of the work plan bid group
    Name string `json:"name"`


    // ManagementUnitId - The management unit ID used for this work plan bid group
    ManagementUnitId string `json:"managementUnitId"`


    // AgentIds - Agent IDs who participate in this bid group
    AgentIds Listwrapperstring `json:"agentIds"`


    // WorkPlans - The list of work plans used in this bid group
    WorkPlans Listwrapperbidgroupworkplanrequest `json:"workPlans"`


    // PlanningGroupIds - The planning group IDs selected in this bid group
    PlanningGroupIds Listwrapperstring `json:"planningGroupIds"`

}

// String returns a JSON representation of the model
func (o *Workplanbidgroupupdate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidgroupupdate) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidgroupupdate

    if WorkplanbidgroupupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidgroupupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnitId string `json:"managementUnitId"`
        
        AgentIds Listwrapperstring `json:"agentIds"`
        
        WorkPlans Listwrapperbidgroupworkplanrequest `json:"workPlans"`
        
        PlanningGroupIds Listwrapperstring `json:"planningGroupIds"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

