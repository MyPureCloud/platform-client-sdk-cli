package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntakesettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntakesettingDud struct { 
    


    


    

}

// Intakesetting
type Intakesetting struct { 
    // Property - The property name for this intake setting.
    Property string `json:"property"`


    // Required - Defines if this property is required for intake
    Required bool `json:"required"`


    // DisplayOrder - The order where this property should be displayed
    DisplayOrder int `json:"displayOrder"`

}

// String returns a JSON representation of the model
func (o *Intakesetting) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intakesetting) MarshalJSON() ([]byte, error) {
    type Alias Intakesetting

    if IntakesettingMarshalled {
        return []byte("{}"), nil
    }
    IntakesettingMarshalled = true

    return json.Marshal(&struct {
        
        Property string `json:"property"`
        
        Required bool `json:"required"`
        
        DisplayOrder int `json:"displayOrder"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

