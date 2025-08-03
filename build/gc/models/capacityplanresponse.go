package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Capacityplanresponse
type Capacityplanresponse struct { 
    


    // Name
    Name string `json:"name"`


    // Description - Description of the capacity plan
    Description string `json:"description"`


    // Forecast - The selected forecast for this capacity plan. Null when main forecast is used in the future
    Forecast Bushorttermforecastreference `json:"forecast"`


    // StartBusinessUnitDate - The start date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`


    // EndBusinessUnitDate - The end date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`


    // FullTimeEquivalentWeeklyHours - The weekly hours used to calculate full time equivalent agents
    FullTimeEquivalentWeeklyHours float64 `json:"fullTimeEquivalentWeeklyHours"`


    // Metadata - The metadata of this capacity plan
    Metadata Capacityplanmetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Capacityplanresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanresponse) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanresponse

    if CapacityplanresponseMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Forecast Bushorttermforecastreference `json:"forecast"`
        
        StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`
        
        EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`
        
        FullTimeEquivalentWeeklyHours float64 `json:"fullTimeEquivalentWeeklyHours"`
        
        Metadata Capacityplanmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

