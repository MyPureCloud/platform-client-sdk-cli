package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrenddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrenddataDud struct { 
    


    


    

}

// Trenddata
type Trenddata struct { 
    // DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    // DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEndWorkday time.Time `json:"dateEndWorkday"`


    // PercentOfGoal - Percent of goal
    PercentOfGoal float64 `json:"percentOfGoal"`

}

// String returns a JSON representation of the model
func (o *Trenddata) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trenddata) MarshalJSON() ([]byte, error) {
    type Alias Trenddata

    if TrenddataMarshalled {
        return []byte("{}"), nil
    }
    TrenddataMarshalled = true

    return json.Marshal(&struct {
        
        DateStartWorkday time.Time `json:"dateStartWorkday"`
        
        DateEndWorkday time.Time `json:"dateEndWorkday"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

