package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetwrapperassignmenteffectivedaterangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetwrapperassignmenteffectivedaterangeDud struct { 
    

}

// Setwrapperassignmenteffectivedaterange
type Setwrapperassignmenteffectivedaterange struct { 
    // Values
    Values []Assignmenteffectivedaterange `json:"values"`

}

// String returns a JSON representation of the model
func (o *Setwrapperassignmenteffectivedaterange) String() string {
     o.Values = []Assignmenteffectivedaterange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setwrapperassignmenteffectivedaterange) MarshalJSON() ([]byte, error) {
    type Alias Setwrapperassignmenteffectivedaterange

    if SetwrapperassignmenteffectivedaterangeMarshalled {
        return []byte("{}"), nil
    }
    SetwrapperassignmenteffectivedaterangeMarshalled = true

    return json.Marshal(&struct {
        
        Values []Assignmenteffectivedaterange `json:"values"`
        *Alias
    }{

        
        Values: []Assignmenteffectivedaterange{{}},
        

        Alias: (*Alias)(u),
    })
}

