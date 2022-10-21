package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HristimeofftypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HristimeofftypeDud struct { 
    


    


    


    

}

// Hristimeofftype
type Hristimeofftype struct { 
    // Id - The ID of the time off type configured in integration.
    Id string `json:"id"`


    // Name - The name of the time off type configured in integration.
    Name string `json:"name"`


    // HrisIntegrationId - The ID of the integration.
    HrisIntegrationId string `json:"hrisIntegrationId"`


    // SecondaryId - Secondary ID of the time off type, if configured in integration.
    SecondaryId string `json:"secondaryId"`

}

// String returns a JSON representation of the model
func (o *Hristimeofftype) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Hristimeofftype) MarshalJSON() ([]byte, error) {
    type Alias Hristimeofftype

    if HristimeofftypeMarshalled {
        return []byte("{}"), nil
    }
    HristimeofftypeMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        HrisIntegrationId string `json:"hrisIntegrationId"`
        
        SecondaryId string `json:"secondaryId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

