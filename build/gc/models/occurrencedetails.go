package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OccurrencedetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OccurrencedetailsDud struct { 
    


    

}

// Occurrencedetails
type Occurrencedetails struct { 
    // DateOfNextOccurrence - The date of the next start or end occurrence for the recurrence as an ISO-8601 string
    DateOfNextOccurrence time.Time `json:"dateOfNextOccurrence"`


    // NumberOfOccurrences - The number of start or end occurrences that have been processed for the recurrence.
    NumberOfOccurrences int `json:"numberOfOccurrences"`

}

// String returns a JSON representation of the model
func (o *Occurrencedetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Occurrencedetails) MarshalJSON() ([]byte, error) {
    type Alias Occurrencedetails

    if OccurrencedetailsMarshalled {
        return []byte("{}"), nil
    }
    OccurrencedetailsMarshalled = true

    return json.Marshal(&struct {
        
        DateOfNextOccurrence time.Time `json:"dateOfNextOccurrence"`
        
        NumberOfOccurrences int `json:"numberOfOccurrences"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

