package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NextoccurrencedetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NextoccurrencedetailsDud struct { 
    


    

}

// Nextoccurrencedetails
type Nextoccurrencedetails struct { 
    // StartOccurrenceDetails - The details for the next start occurrence for the recurrence.
    StartOccurrenceDetails Occurrencedetails `json:"startOccurrenceDetails"`


    // EndOccurrenceDetails - The details for the next end occurrence for the recurrence.
    EndOccurrenceDetails Occurrencedetails `json:"endOccurrenceDetails"`

}

// String returns a JSON representation of the model
func (o *Nextoccurrencedetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nextoccurrencedetails) MarshalJSON() ([]byte, error) {
    type Alias Nextoccurrencedetails

    if NextoccurrencedetailsMarshalled {
        return []byte("{}"), nil
    }
    NextoccurrencedetailsMarshalled = true

    return json.Marshal(&struct {
        
        StartOccurrenceDetails Occurrencedetails `json:"startOccurrenceDetails"`
        
        EndOccurrenceDetails Occurrencedetails `json:"endOccurrenceDetails"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

