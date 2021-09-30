package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailabletimeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailabletimeDud struct { 
    DateStart time.Time `json:"dateStart"`


    LengthInMinutes int `json:"lengthInMinutes"`


    IsPaid bool `json:"isPaid"`


    ActivityCategory string `json:"activityCategory"`


    WfmSchedule Wfmschedulereference `json:"wfmSchedule"`

}

// Availabletime
type Availabletime struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Availabletime) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availabletime) MarshalJSON() ([]byte, error) {
    type Alias Availabletime

    if AvailabletimeMarshalled {
        return []byte("{}"), nil
    }
    AvailabletimeMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

