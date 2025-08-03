package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatecapacityplanstaffinggroupmetricchangerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatecapacityplanstaffinggroupmetricchangerequestDud struct { 
    


    


    


    


    


    


    

}

// Createcapacityplanstaffinggroupmetricchangerequest
type Createcapacityplanstaffinggroupmetricchangerequest struct { 
    // NumberOfWeeks - The number of weeks to which the metric change applies
    NumberOfWeeks int `json:"numberOfWeeks"`


    // WeekStartNumber - The start number of the week (starting from 1) to which the metric change applies, related to numberOfWeeks
    WeekStartNumber int `json:"weekStartNumber"`


    // Value - The value of the metric
    Value float64 `json:"value"`


    // Metric - The metric which is going to be modified for the selected staffing groups
    Metric string `json:"metric"`


    // Notes - Notes about the staffing groups metric changes
    Notes string `json:"notes"`


    // StaffingGroupIds - The IDs of the staffing groups affected by the metric change
    StaffingGroupIds []string `json:"staffingGroupIds"`


    // Version - The version of the capacity plan
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Createcapacityplanstaffinggroupmetricchangerequest) String() string {
    
    
    
    
    
     o.StaffingGroupIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createcapacityplanstaffinggroupmetricchangerequest) MarshalJSON() ([]byte, error) {
    type Alias Createcapacityplanstaffinggroupmetricchangerequest

    if CreatecapacityplanstaffinggroupmetricchangerequestMarshalled {
        return []byte("{}"), nil
    }
    CreatecapacityplanstaffinggroupmetricchangerequestMarshalled = true

    return json.Marshal(&struct {
        
        NumberOfWeeks int `json:"numberOfWeeks"`
        
        WeekStartNumber int `json:"weekStartNumber"`
        
        Value float64 `json:"value"`
        
        Metric string `json:"metric"`
        
        Notes string `json:"notes"`
        
        StaffingGroupIds []string `json:"staffingGroupIds"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        


        
        StaffingGroupIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

