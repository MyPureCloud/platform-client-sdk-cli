package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ObjectiveMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ObjectiveDud struct { 
    Id string `json:"id"`


    


    


    


    

}

// Objective
type Objective struct { 
    


    // TemplateId - The id of this objective's base template
    TemplateId string `json:"templateId"`


    // Zones - Objective zone specifies min,max points and values for the associated metric
    Zones []Objectivezone `json:"zones"`


    // Enabled - A flag for whether this objective is enabled for the related metric
    Enabled bool `json:"enabled"`


    // DateStart - start date of the objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`

}

// String returns a JSON representation of the model
func (o *Objective) String() string {
    
    
    
    
    
    
    
    
     o.Zones = []Objectivezone{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Objective) MarshalJSON() ([]byte, error) {
    type Alias Objective

    if ObjectiveMarshalled {
        return []byte("{}"), nil
    }
    ObjectiveMarshalled = true

    return json.Marshal(&struct { 
        
        
        TemplateId string `json:"templateId"`
        
        Zones []Objectivezone `json:"zones"`
        
        Enabled bool `json:"enabled"`
        
        DateStart time.Time `json:"dateStart"`
        
        *Alias
    }{
        

        

        

        

        

        
        Zones: []Objectivezone{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

