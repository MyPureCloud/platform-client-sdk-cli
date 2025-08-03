package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatecapacityplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatecapacityplanrequestDud struct { 
    


    


    


    


    


    


    


    

}

// Updatecapacityplanrequest
type Updatecapacityplanrequest struct { 
    // Name - The name of the capacity plan
    Name string `json:"name"`


    // Description - Description of the capacity plan
    Description string `json:"description"`


    // StartBusinessUnitDate - The start date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`


    // EndBusinessUnitDate - The end date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`


    // Forecast - The selected forecast for this capacity plan
    Forecast Valuewrapperbushorttermforecastreference `json:"forecast"`


    // FullTimeEquivalentWeeklyHours - The weekly hours used to calculate full time equivalent agents
    FullTimeEquivalentWeeklyHours float64 `json:"fullTimeEquivalentWeeklyHours"`


    // UseLatestPlanningGroupStaffingGroupAssociation - Whether to use latest staffing group to planning group association
    UseLatestPlanningGroupStaffingGroupAssociation bool `json:"useLatestPlanningGroupStaffingGroupAssociation"`


    // Metadata - The metadata of this capacity plan
    Metadata Capacityplanmetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updatecapacityplanrequest) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatecapacityplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatecapacityplanrequest

    if UpdatecapacityplanrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatecapacityplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`
        
        EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`
        
        Forecast Valuewrapperbushorttermforecastreference `json:"forecast"`
        
        FullTimeEquivalentWeeklyHours float64 `json:"fullTimeEquivalentWeeklyHours"`
        
        UseLatestPlanningGroupStaffingGroupAssociation bool `json:"useLatestPlanningGroupStaffingGroupAssociation"`
        
        Metadata Capacityplanmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

