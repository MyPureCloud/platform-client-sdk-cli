package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanrequestDud struct { 
    


    


    


    


    


    


    

}

// Capacityplanrequest
type Capacityplanrequest struct { 
    // Name - The name of the capacity plan
    Name string `json:"name"`


    // Description - Description of the capacity plan
    Description string `json:"description"`


    // StartBusinessUnitDate - The start date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`


    // EndBusinessUnitDate - The end date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`


    // Forecast - The selected forecast for this capacity plan. Null when main forecast is used in the future
    Forecast Bushorttermforecastreference `json:"forecast"`


    // FullTimeEquivalentWeeklyHours - The weekly hours used to calculate full time equivalent agents
    FullTimeEquivalentWeeklyHours float64 `json:"fullTimeEquivalentWeeklyHours"`


    // StaffingGroupAllocations - List of staffing group allocations to be used for the capacity plan
    StaffingGroupAllocations []Createstaffinggroupallocation `json:"staffingGroupAllocations"`

}

// String returns a JSON representation of the model
func (o *Capacityplanrequest) String() string {
    
    
    
    
    
    
     o.StaffingGroupAllocations = []Createstaffinggroupallocation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanrequest

    if CapacityplanrequestMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`
        
        EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`
        
        Forecast Bushorttermforecastreference `json:"forecast"`
        
        FullTimeEquivalentWeeklyHours float64 `json:"fullTimeEquivalentWeeklyHours"`
        
        StaffingGroupAllocations []Createstaffinggroupallocation `json:"staffingGroupAllocations"`
        *Alias
    }{

        


        


        


        


        


        


        
        StaffingGroupAllocations: []Createstaffinggroupallocation{{}},
        

        Alias: (*Alias)(u),
    })
}

