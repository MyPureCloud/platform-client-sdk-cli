package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatesegmentassignmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatesegmentassignmentrequestDud struct { 
    


    

}

// Updatesegmentassignmentrequest
type Updatesegmentassignmentrequest struct { 
    // Assign - The segment assignments to apply.
    Assign Segmentassignments `json:"assign"`


    // Unassign - The segment unassignments to apply.
    Unassign Segmentunassignments `json:"unassign"`

}

// String returns a JSON representation of the model
func (o *Updatesegmentassignmentrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatesegmentassignmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatesegmentassignmentrequest

    if UpdatesegmentassignmentrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatesegmentassignmentrequestMarshalled = true

    return json.Marshal(&struct {
        
        Assign Segmentassignments `json:"assign"`
        
        Unassign Segmentunassignments `json:"unassign"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

