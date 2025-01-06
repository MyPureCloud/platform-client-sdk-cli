package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MonthlyplanningperiodsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MonthlyplanningperiodsettingsDud struct { 
    

}

// Monthlyplanningperiodsettings
type Monthlyplanningperiodsettings struct { 
    // StartDate - Start date of the monthly planning period in yyyy-MM-dd format. The date must represent the first day of the given month. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`

}

// String returns a JSON representation of the model
func (o *Monthlyplanningperiodsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Monthlyplanningperiodsettings) MarshalJSON() ([]byte, error) {
    type Alias Monthlyplanningperiodsettings

    if MonthlyplanningperiodsettingsMarshalled {
        return []byte("{}"), nil
    }
    MonthlyplanningperiodsettingsMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

