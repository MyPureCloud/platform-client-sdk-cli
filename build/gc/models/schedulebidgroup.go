package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidgroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidgroupDud struct { 
    


    


    


    


    


    


    


    

}

// Schedulebidgroup
type Schedulebidgroup struct { 
    // Name - The name of the schedule bid group
    Name string `json:"name"`


    // ManagementUnit - The management unit to which this bid group belongs
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // Agents - The agents who participate in this bid group
    Agents []Userreference `json:"agents"`


    // WorkPlans - The work plans used in this bid group
    WorkPlans []Workplanreference `json:"workPlans"`


    // WorkPlanRotations - The work plan rotations used in this bid group
    WorkPlanRotations []Bidgroupworkplanrotationresponse `json:"workPlanRotations"`


    // PlanningGroups - The planning groups selected in this bid group
    PlanningGroups []Planninggroupreference `json:"planningGroups"`


    // DownloadUrl - The downloadUrl to fetch Schedule sets. It will be populated if the status of this bid is 'Optimized'
    DownloadUrl string `json:"downloadUrl"`


    // DownloadTemplate - Schedule sets always come through downloadUrl, the schema included here is just for documentation
    DownloadTemplate Bidgroupscheduleset `json:"downloadTemplate"`

}

// String returns a JSON representation of the model
func (o *Schedulebidgroup) String() string {
    
    
     o.Agents = []Userreference{{}} 
     o.WorkPlans = []Workplanreference{{}} 
     o.WorkPlanRotations = []Bidgroupworkplanrotationresponse{{}} 
     o.PlanningGroups = []Planninggroupreference{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidgroup) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidgroup

    if SchedulebidgroupMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidgroupMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        Agents []Userreference `json:"agents"`
        
        WorkPlans []Workplanreference `json:"workPlans"`
        
        WorkPlanRotations []Bidgroupworkplanrotationresponse `json:"workPlanRotations"`
        
        PlanningGroups []Planninggroupreference `json:"planningGroups"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadTemplate Bidgroupscheduleset `json:"downloadTemplate"`
        *Alias
    }{

        


        


        
        Agents: []Userreference{{}},
        


        
        WorkPlans: []Workplanreference{{}},
        


        
        WorkPlanRotations: []Bidgroupworkplanrotationresponse{{}},
        


        
        PlanningGroups: []Planninggroupreference{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

