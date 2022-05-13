package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanningperiodsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanningperiodsettingsDud struct { 
    


    

}

// Planningperiodsettings
type Planningperiodsettings struct { 
    // WeekCount - Planning period length in weeks
    WeekCount int `json:"weekCount"`


    // StartDate - Start date of the planning period in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`

}

// String returns a JSON representation of the model
func (o *Planningperiodsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planningperiodsettings) MarshalJSON() ([]byte, error) {
    type Alias Planningperiodsettings

    if PlanningperiodsettingsMarshalled {
        return []byte("{}"), nil
    }
    PlanningperiodsettingsMarshalled = true

    return json.Marshal(&struct {
        
        WeekCount int `json:"weekCount"`
        
        StartDate time.Time `json:"startDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

