package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeneratebuforecastrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeneratebuforecastrequestDud struct { 
    


    


    

}

// Generatebuforecastrequest
type Generatebuforecastrequest struct { 
    // Description - The description for the forecast
    Description string `json:"description"`


    // WeekCount - The number of weeks this forecast covers
    WeekCount int `json:"weekCount"`


    // CanUseForScheduling - Whether this forecast can be used for scheduling
    CanUseForScheduling bool `json:"canUseForScheduling"`

}

// String returns a JSON representation of the model
func (o *Generatebuforecastrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generatebuforecastrequest) MarshalJSON() ([]byte, error) {
    type Alias Generatebuforecastrequest

    if GeneratebuforecastrequestMarshalled {
        return []byte("{}"), nil
    }
    GeneratebuforecastrequestMarshalled = true

    return json.Marshal(&struct { 
        Description string `json:"description"`
        
        WeekCount int `json:"weekCount"`
        
        CanUseForScheduling bool `json:"canUseForScheduling"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

