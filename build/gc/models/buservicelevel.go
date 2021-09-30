package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuservicelevelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuservicelevelDud struct { 
    


    


    

}

// Buservicelevel - Service goal service level configuration
type Buservicelevel struct { 
    // Include - Whether to include service level targets in the associated configuration
    Include bool `json:"include"`


    // Percent - Service level target percent answered. Required if include == true
    Percent int `json:"percent"`


    // Seconds - Service level target answer time. Required if include == true
    Seconds int `json:"seconds"`

}

// String returns a JSON representation of the model
func (o *Buservicelevel) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buservicelevel) MarshalJSON() ([]byte, error) {
    type Alias Buservicelevel

    if BuservicelevelMarshalled {
        return []byte("{}"), nil
    }
    BuservicelevelMarshalled = true

    return json.Marshal(&struct { 
        Include bool `json:"include"`
        
        Percent int `json:"percent"`
        
        Seconds int `json:"seconds"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

