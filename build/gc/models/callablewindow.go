package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallablewindowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallablewindowDud struct { 
    


    

}

// Callablewindow
type Callablewindow struct { 
    // Mapped - The time interval to place outbound calls, for contacts that can be mapped to a time zone.
    Mapped Atzmtimeslot `json:"mapped"`


    // Unmapped - The time interval and time zone to place outbound calls, for contacts that cannot be mapped to a time zone.
    Unmapped Atzmtimeslotwithtimezone `json:"unmapped"`

}

// String returns a JSON representation of the model
func (o *Callablewindow) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callablewindow) MarshalJSON() ([]byte, error) {
    type Alias Callablewindow

    if CallablewindowMarshalled {
        return []byte("{}"), nil
    }
    CallablewindowMarshalled = true

    return json.Marshal(&struct { 
        Mapped Atzmtimeslot `json:"mapped"`
        
        Unmapped Atzmtimeslotwithtimezone `json:"unmapped"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

