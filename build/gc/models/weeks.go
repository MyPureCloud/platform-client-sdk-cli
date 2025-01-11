package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeeksMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeeksDud struct { 
    


    

}

// Weeks
type Weeks struct { 
    // WeekOffset - The week offset from data start date
    WeekOffset int `json:"weekOffset"`


    // Values - data per interval (daily or quarter hour) for this week. The value could be double or null
    Values []float64 `json:"values"`

}

// String returns a JSON representation of the model
func (o *Weeks) String() string {
    
     o.Values = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weeks) MarshalJSON() ([]byte, error) {
    type Alias Weeks

    if WeeksMarshalled {
        return []byte("{}"), nil
    }
    WeeksMarshalled = true

    return json.Marshal(&struct {
        
        WeekOffset int `json:"weekOffset"`
        
        Values []float64 `json:"values"`
        *Alias
    }{

        


        
        Values: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

