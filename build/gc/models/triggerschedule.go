package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TriggerscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TriggerscheduleDud struct { 
    


    


    


    


    


    

}

// Triggerschedule - Schedule configuration for a scheduled trigger
type Triggerschedule struct { 
    // Minutes - Minutes on which the trigger should fire. Coma separated list of up to 2 values 0-59
    Minutes string `json:"minutes"`


    // Hours - Hours on which the trigger should fire. 0-23 or '*' for every hour.
    Hours string `json:"hours"`


    // DaysOfMonth - Days of month on which the trigger should fire. 1-31 or '*' for every or '?' for any
    DaysOfMonth string `json:"daysOfMonth"`


    // Months - Months on which the trigger should fire. 1-12 or '*' for every
    Months string `json:"months"`


    // DaysOfWeek - Days of week on which the trigger should fire. 1-7 or '*' for every or '?' for any
    DaysOfWeek string `json:"daysOfWeek"`


    // Timezone - Timezone for the trigger schedule
    Timezone string `json:"timezone"`

}

// String returns a JSON representation of the model
func (o *Triggerschedule) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Triggerschedule) MarshalJSON() ([]byte, error) {
    type Alias Triggerschedule

    if TriggerscheduleMarshalled {
        return []byte("{}"), nil
    }
    TriggerscheduleMarshalled = true

    return json.Marshal(&struct {
        
        Minutes string `json:"minutes"`
        
        Hours string `json:"hours"`
        
        DaysOfMonth string `json:"daysOfMonth"`
        
        Months string `json:"months"`
        
        DaysOfWeek string `json:"daysOfWeek"`
        
        Timezone string `json:"timezone"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

