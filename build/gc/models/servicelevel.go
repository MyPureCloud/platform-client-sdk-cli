package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicelevelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicelevelDud struct { 
    


    

}

// Servicelevel
type Servicelevel struct { 
    // Percentage - The desired Service Level. A value between 0 and 1.
    Percentage float64 `json:"percentage"`


    // DurationMs - Service Level target in milliseconds.
    DurationMs int `json:"durationMs"`

}

// String returns a JSON representation of the model
func (o *Servicelevel) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicelevel) MarshalJSON() ([]byte, error) {
    type Alias Servicelevel

    if ServicelevelMarshalled {
        return []byte("{}"), nil
    }
    ServicelevelMarshalled = true

    return json.Marshal(&struct { 
        Percentage float64 `json:"percentage"`
        
        DurationMs int `json:"durationMs"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

