package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanstaffinggroupallocationsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanstaffinggroupallocationsresponseDud struct { 
    


    


    


    

}

// Capacityplanstaffinggroupallocationsresponse
type Capacityplanstaffinggroupallocationsresponse struct { 
    // CapacityPlan - The capacity plan to which the staffing groups belong
    CapacityPlan Capacityplanreference `json:"capacityPlan"`


    // FullTimeEquivalentWeeklyHours - The weekly hours used to calculate full time equivalent agents
    FullTimeEquivalentWeeklyHours float64 `json:"fullTimeEquivalentWeeklyHours"`


    // DownloadUrl - The URL to download the staffing group allocations
    DownloadUrl string `json:"downloadUrl"`


    // DownloadTemplate - Staffing groups allocation results always come through downloadUrl, the schema included here is just for documentation
    DownloadTemplate Staffinggroupallocationsresponsetemplate `json:"downloadTemplate"`

}

// String returns a JSON representation of the model
func (o *Capacityplanstaffinggroupallocationsresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanstaffinggroupallocationsresponse) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanstaffinggroupallocationsresponse

    if CapacityplanstaffinggroupallocationsresponseMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanstaffinggroupallocationsresponseMarshalled = true

    return json.Marshal(&struct {
        
        CapacityPlan Capacityplanreference `json:"capacityPlan"`
        
        FullTimeEquivalentWeeklyHours float64 `json:"fullTimeEquivalentWeeklyHours"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadTemplate Staffinggroupallocationsresponsetemplate `json:"downloadTemplate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

