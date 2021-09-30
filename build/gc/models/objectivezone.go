package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ObjectivezoneMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ObjectivezoneDud struct { 
    


    


    


    


    


    


    

}

// Objectivezone
type Objectivezone struct { 
    // Label - label
    Label string `json:"label"`


    // DirectionType - direction type
    DirectionType string `json:"directionType"`


    // ZoneType - zone type
    ZoneType string `json:"zoneType"`


    // UpperLimitPoints - upper limit points
    UpperLimitPoints int `json:"upperLimitPoints"`


    // LowerLimitPoints - lower limit points
    LowerLimitPoints int `json:"lowerLimitPoints"`


    // UpperLimitValue - upper limit value
    UpperLimitValue int `json:"upperLimitValue"`


    // LowerLimitValue - lower limit value
    LowerLimitValue int `json:"lowerLimitValue"`

}

// String returns a JSON representation of the model
func (o *Objectivezone) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Objectivezone) MarshalJSON() ([]byte, error) {
    type Alias Objectivezone

    if ObjectivezoneMarshalled {
        return []byte("{}"), nil
    }
    ObjectivezoneMarshalled = true

    return json.Marshal(&struct { 
        Label string `json:"label"`
        
        DirectionType string `json:"directionType"`
        
        ZoneType string `json:"zoneType"`
        
        UpperLimitPoints int `json:"upperLimitPoints"`
        
        LowerLimitPoints int `json:"lowerLimitPoints"`
        
        UpperLimitValue int `json:"upperLimitValue"`
        
        LowerLimitValue int `json:"lowerLimitValue"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

