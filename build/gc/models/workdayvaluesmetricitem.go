package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkdayvaluesmetricitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkdayvaluesmetricitemDud struct { 
    MetricDefinition Domainentityref `json:"metricDefinition"`


    Average float64 `json:"average"`


    UnitType string `json:"unitType"`


    Trend []Workdayvaluestrenditem `json:"trend"`

}

// Workdayvaluesmetricitem
type Workdayvaluesmetricitem struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Workdayvaluesmetricitem) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workdayvaluesmetricitem) MarshalJSON() ([]byte, error) {
    type Alias Workdayvaluesmetricitem

    if WorkdayvaluesmetricitemMarshalled {
        return []byte("{}"), nil
    }
    WorkdayvaluesmetricitemMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

