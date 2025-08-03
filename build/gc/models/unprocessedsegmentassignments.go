package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnprocessedsegmentassignmentsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnprocessedsegmentassignmentsDud struct { 
    


    


    

}

// Unprocessedsegmentassignments
type Unprocessedsegmentassignments struct { 
    // Assign - The segment assignments to apply.
    Assign Segmentassignments `json:"assign"`


    // Unassign - The segment unassignments to apply.
    Unassign Segmentunassignments `json:"unassign"`


    // Count - The total number of segment assignments and unassignments that were not successfully processed.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Unprocessedsegmentassignments) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unprocessedsegmentassignments) MarshalJSON() ([]byte, error) {
    type Alias Unprocessedsegmentassignments

    if UnprocessedsegmentassignmentsMarshalled {
        return []byte("{}"), nil
    }
    UnprocessedsegmentassignmentsMarshalled = true

    return json.Marshal(&struct {
        
        Assign Segmentassignments `json:"assign"`
        
        Unassign Segmentunassignments `json:"unassign"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

