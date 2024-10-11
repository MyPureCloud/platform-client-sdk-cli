package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MuagentsworkplansresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MuagentsworkplansresultDud struct { 
    


    


    

}

// Muagentsworkplansresult
type Muagentsworkplansresult struct { 
    // Entities
    Entities []Agentworkplans `json:"entities"`


    // ReferenceStartWeekDate - The reference date in yyyy-MM-dd format rolled back to nearest BU start day of week. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    ReferenceStartWeekDate time.Time `json:"referenceStartWeekDate"`


    // WorkPlanLookup - Map containing lookup values for agent work plans. The integer keys serves as lookup keys for effective work plan from workPlanLookupKeysPerWeek property
    WorkPlanLookup map[string]Workplanreference `json:"workPlanLookup"`

}

// String returns a JSON representation of the model
func (o *Muagentsworkplansresult) String() string {
     o.Entities = []Agentworkplans{{}} 
    
     o.WorkPlanLookup = map[string]Workplanreference{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Muagentsworkplansresult) MarshalJSON() ([]byte, error) {
    type Alias Muagentsworkplansresult

    if MuagentsworkplansresultMarshalled {
        return []byte("{}"), nil
    }
    MuagentsworkplansresultMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Agentworkplans `json:"entities"`
        
        ReferenceStartWeekDate time.Time `json:"referenceStartWeekDate"`
        
        WorkPlanLookup map[string]Workplanreference `json:"workPlanLookup"`
        *Alias
    }{

        
        Entities: []Agentworkplans{{}},
        


        


        
        WorkPlanLookup: map[string]Workplanreference{"": {}},
        

        Alias: (*Alias)(u),
    })
}

