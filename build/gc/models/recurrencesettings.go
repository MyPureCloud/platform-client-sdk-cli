package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecurrencesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecurrencesettingsDud struct { 
    


    


    

}

// Recurrencesettings
type Recurrencesettings struct { 
    // EndAfter - Settings controlling when to end the recurrence for the activity plan
    EndAfter Recurrenceendsettings `json:"endAfter"`


    // RecurrencePeriod - The recurrence period of the activity plan
    RecurrencePeriod Recurrenceperiod `json:"recurrencePeriod"`


    // MinimumTimeBetweenOccurrences - Constraint indicating the minimum time in hours between recurrences of the activity plan
    MinimumTimeBetweenOccurrences Recurrenceperiod `json:"minimumTimeBetweenOccurrences"`

}

// String returns a JSON representation of the model
func (o *Recurrencesettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recurrencesettings) MarshalJSON() ([]byte, error) {
    type Alias Recurrencesettings

    if RecurrencesettingsMarshalled {
        return []byte("{}"), nil
    }
    RecurrencesettingsMarshalled = true

    return json.Marshal(&struct {
        
        EndAfter Recurrenceendsettings `json:"endAfter"`
        
        RecurrencePeriod Recurrenceperiod `json:"recurrencePeriod"`
        
        MinimumTimeBetweenOccurrences Recurrenceperiod `json:"minimumTimeBetweenOccurrences"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

