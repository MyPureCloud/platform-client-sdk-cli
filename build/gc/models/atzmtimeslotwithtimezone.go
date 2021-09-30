package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AtzmtimeslotwithtimezoneMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AtzmtimeslotwithtimezoneDud struct { 
    


    


    

}

// Atzmtimeslotwithtimezone
type Atzmtimeslotwithtimezone struct { 
    // EarliestCallableTime - The earliest time to dial a contact. Valid format is HH:mm
    EarliestCallableTime string `json:"earliestCallableTime"`


    // LatestCallableTime - The latest time to dial a contact. Valid format is HH:mm
    LatestCallableTime string `json:"latestCallableTime"`


    // TimeZoneId - The time zone to use for contacts that cannot be mapped.
    TimeZoneId string `json:"timeZoneId"`

}

// String returns a JSON representation of the model
func (o *Atzmtimeslotwithtimezone) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Atzmtimeslotwithtimezone) MarshalJSON() ([]byte, error) {
    type Alias Atzmtimeslotwithtimezone

    if AtzmtimeslotwithtimezoneMarshalled {
        return []byte("{}"), nil
    }
    AtzmtimeslotwithtimezoneMarshalled = true

    return json.Marshal(&struct { 
        EarliestCallableTime string `json:"earliestCallableTime"`
        
        LatestCallableTime string `json:"latestCallableTime"`
        
        TimeZoneId string `json:"timeZoneId"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

