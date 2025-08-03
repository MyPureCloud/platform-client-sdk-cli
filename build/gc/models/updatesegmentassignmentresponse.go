package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatesegmentassignmentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatesegmentassignmentresponseDud struct { 
    

}

// Updatesegmentassignmentresponse
type Updatesegmentassignmentresponse struct { 
    // UnprocessedItems - The segment assignments and unassignments which could not be processed.
    UnprocessedItems Unprocessedsegmentassignments `json:"unprocessedItems"`

}

// String returns a JSON representation of the model
func (o *Updatesegmentassignmentresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatesegmentassignmentresponse) MarshalJSON() ([]byte, error) {
    type Alias Updatesegmentassignmentresponse

    if UpdatesegmentassignmentresponseMarshalled {
        return []byte("{}"), nil
    }
    UpdatesegmentassignmentresponseMarshalled = true

    return json.Marshal(&struct {
        
        UnprocessedItems Unprocessedsegmentassignments `json:"unprocessedItems"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

