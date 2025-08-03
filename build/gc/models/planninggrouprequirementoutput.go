package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggrouprequirementoutputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggrouprequirementoutputDud struct { 
    


    


    


    

}

// Planninggrouprequirementoutput
type Planninggrouprequirementoutput struct { 
    // Id - The ID of the planning group to which this output applies
    Id string `json:"id"`


    // Intervals - List of interval values that correspond with the requirements granularity that was requested
    Intervals []int `json:"intervals"`


    // ErrorDetails - Error details if the intervals cannot be provided for this planning group because of missing data or internal error
    ErrorDetails []Longtermrequirementserrordetail `json:"errorDetails"`


    // ServiceGoalDetails - The service goal details used to generate the requirements
    ServiceGoalDetails Longtermrequirementsservicegoaldetail `json:"serviceGoalDetails"`

}

// String returns a JSON representation of the model
func (o *Planninggrouprequirementoutput) String() string {
    
     o.Intervals = []int{0} 
     o.ErrorDetails = []Longtermrequirementserrordetail{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggrouprequirementoutput) MarshalJSON() ([]byte, error) {
    type Alias Planninggrouprequirementoutput

    if PlanninggrouprequirementoutputMarshalled {
        return []byte("{}"), nil
    }
    PlanninggrouprequirementoutputMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Intervals []int `json:"intervals"`
        
        ErrorDetails []Longtermrequirementserrordetail `json:"errorDetails"`
        
        ServiceGoalDetails Longtermrequirementsservicegoaldetail `json:"serviceGoalDetails"`
        *Alias
    }{

        


        
        Intervals: []int{0},
        


        
        ErrorDetails: []Longtermrequirementserrordetail{{}},
        


        

        Alias: (*Alias)(u),
    })
}

