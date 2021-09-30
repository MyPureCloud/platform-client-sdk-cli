package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeintervalDud struct { 
    


    


    


    

}

// Timeinterval
type Timeinterval struct { 
    // Months
    Months int `json:"months"`


    // Weeks
    Weeks int `json:"weeks"`


    // Days
    Days int `json:"days"`


    // Hours
    Hours int `json:"hours"`

}

// String returns a JSON representation of the model
func (o *Timeinterval) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeinterval) MarshalJSON() ([]byte, error) {
    type Alias Timeinterval

    if TimeintervalMarshalled {
        return []byte("{}"), nil
    }
    TimeintervalMarshalled = true

    return json.Marshal(&struct { 
        Months int `json:"months"`
        
        Weeks int `json:"weeks"`
        
        Days int `json:"days"`
        
        Hours int `json:"hours"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

