package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidgroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidgroupDud struct { 
    


    


    


    


    

}

// Workplanbidgroup
type Workplanbidgroup struct { 
    // Name - The name of the work plan bid group
    Name string `json:"name"`


    // ManagementUnit - The management unit this bid group belongs to
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // Agents - The list of agents who participate in this bid group
    Agents []Userreference `json:"agents"`


    // WorkPlans - The list of work plans used in this bid group
    WorkPlans []Bidgroupworkplanresponse `json:"workPlans"`


    // PlanningGroups - The list of planning groups selected in this bid group
    PlanningGroups []Planninggroupreference `json:"planningGroups"`

}

// String returns a JSON representation of the model
func (o *Workplanbidgroup) String() string {
    
    
     o.Agents = []Userreference{{}} 
     o.WorkPlans = []Bidgroupworkplanresponse{{}} 
     o.PlanningGroups = []Planninggroupreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidgroup) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidgroup

    if WorkplanbidgroupMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidgroupMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        Agents []Userreference `json:"agents"`
        
        WorkPlans []Bidgroupworkplanresponse `json:"workPlans"`
        
        PlanningGroups []Planninggroupreference `json:"planningGroups"`
        *Alias
    }{

        


        


        
        Agents: []Userreference{{}},
        


        
        WorkPlans: []Bidgroupworkplanresponse{{}},
        


        
        PlanningGroups: []Planninggroupreference{{}},
        

        Alias: (*Alias)(u),
    })
}

