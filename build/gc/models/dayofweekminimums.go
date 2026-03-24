package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DayofweekminimumsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DayofweekminimumsDud struct { 
    


    

}

// Dayofweekminimums
type Dayofweekminimums struct { 
    // MinimumValue - The minimum staff value to be applied to this planning group
    MinimumValue float64 `json:"minimumValue"`


    // DaysOfWeek - Days of week this minimum staff value applies to
    DaysOfWeek []string `json:"daysOfWeek"`

}

// String returns a JSON representation of the model
func (o *Dayofweekminimums) String() string {
    
     o.DaysOfWeek = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dayofweekminimums) MarshalJSON() ([]byte, error) {
    type Alias Dayofweekminimums

    if DayofweekminimumsMarshalled {
        return []byte("{}"), nil
    }
    DayofweekminimumsMarshalled = true

    return json.Marshal(&struct {
        
        MinimumValue float64 `json:"minimumValue"`
        
        DaysOfWeek []string `json:"daysOfWeek"`
        *Alias
    }{

        


        
        DaysOfWeek: []string{""},
        

        Alias: (*Alias)(u),
    })
}

