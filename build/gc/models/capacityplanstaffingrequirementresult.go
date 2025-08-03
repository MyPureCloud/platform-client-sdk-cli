package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanstaffingrequirementresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanstaffingrequirementresultDud struct { 
    


    


    


    


    


    


    


    

}

// Capacityplanstaffingrequirementresult
type Capacityplanstaffingrequirementresult struct { 
    // BusinessUnit - The business unit to which the capacity plan belongs
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // CapacityPlan - The capacity plan for which requirements are generated
    CapacityPlan Capacityplanreference `json:"capacityPlan"`


    // Status - The status of the requirement generation of the capacity plan
    Status string `json:"status"`


    // ReferenceBusinessUnitDate - The reference date for interval-based data for the requirements. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    ReferenceBusinessUnitDate time.Time `json:"referenceBusinessUnitDate"`


    // Granularity - Granularity of the intervals
    Granularity string `json:"granularity"`


    // ErrorCode - The error code when status is 'Failed'
    ErrorCode string `json:"errorCode"`


    // DownloadUrl - The URL to get the requirements results for the capacity plan. It will be populated if the status is 'Complete'
    DownloadUrl string `json:"downloadUrl"`


    // DownloadTemplate - Staffing requirement results always come through downloadUrl, the schema included here is just for documentation
    DownloadTemplate Staffingrequirementresultresponsetemplate `json:"downloadTemplate"`

}

// String returns a JSON representation of the model
func (o *Capacityplanstaffingrequirementresult) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanstaffingrequirementresult) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanstaffingrequirementresult

    if CapacityplanstaffingrequirementresultMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanstaffingrequirementresultMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        CapacityPlan Capacityplanreference `json:"capacityPlan"`
        
        Status string `json:"status"`
        
        ReferenceBusinessUnitDate time.Time `json:"referenceBusinessUnitDate"`
        
        Granularity string `json:"granularity"`
        
        ErrorCode string `json:"errorCode"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadTemplate Staffingrequirementresultresponsetemplate `json:"downloadTemplate"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

