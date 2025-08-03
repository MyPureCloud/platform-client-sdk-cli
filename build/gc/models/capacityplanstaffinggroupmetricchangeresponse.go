package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanstaffinggroupmetricchangeresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanstaffinggroupmetricchangeresponseDud struct { 
    


    


    


    


    


    


    


    


    

}

// Capacityplanstaffinggroupmetricchangeresponse
type Capacityplanstaffinggroupmetricchangeresponse struct { 
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


    // StaffingGroups - The staffing groups affected by the metric change
    StaffingGroups []Staffinggroupreference `json:"staffingGroups"`


    // CreatedBy - The user who created the metric change
    CreatedBy Userreference `json:"createdBy"`


    // CreatedDate - The date the entity was created, in ISO-8601 format
    CreatedDate time.Time `json:"createdDate"`


    // Version - The version of the capacity plan
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Capacityplanstaffinggroupmetricchangeresponse) String() string {
    
    
    
    
    
     o.StaffingGroups = []Staffinggroupreference{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanstaffinggroupmetricchangeresponse) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanstaffinggroupmetricchangeresponse

    if CapacityplanstaffinggroupmetricchangeresponseMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanstaffinggroupmetricchangeresponseMarshalled = true

    return json.Marshal(&struct {
        
        NumberOfWeeks int `json:"numberOfWeeks"`
        
        WeekStartNumber int `json:"weekStartNumber"`
        
        Value float64 `json:"value"`
        
        Metric string `json:"metric"`
        
        Notes string `json:"notes"`
        
        StaffingGroups []Staffinggroupreference `json:"staffingGroups"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        


        
        StaffingGroups: []Staffinggroupreference{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

