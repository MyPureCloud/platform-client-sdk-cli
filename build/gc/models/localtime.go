package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocaltimeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocaltimeDud struct { 
    


    


    


    

}

// Localtime
type Localtime struct { 
    // Hour
    Hour int `json:"hour"`


    // Minute
    Minute int `json:"minute"`


    // Second
    Second int `json:"second"`


    // Nano
    Nano int `json:"nano"`

}

// String returns a JSON representation of the model
func (o *Localtime) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Localtime) MarshalJSON() ([]byte, error) {
    type Alias Localtime

    if LocaltimeMarshalled {
        return []byte("{}"), nil
    }
    LocaltimeMarshalled = true

    return json.Marshal(&struct {
        
        Hour int `json:"hour"`
        
        Minute int `json:"minute"`
        
        Second int `json:"second"`
        
        Nano int `json:"nano"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

