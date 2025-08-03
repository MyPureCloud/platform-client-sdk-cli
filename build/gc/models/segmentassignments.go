package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentassignmentsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentassignmentsDud struct { 
    

}

// Segmentassignments
type Segmentassignments struct { 
    // Segments - The segments to be assigned.
    Segments []Segmentforassignment `json:"segments"`

}

// String returns a JSON representation of the model
func (o *Segmentassignments) String() string {
     o.Segments = []Segmentforassignment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentassignments) MarshalJSON() ([]byte, error) {
    type Alias Segmentassignments

    if SegmentassignmentsMarshalled {
        return []byte("{}"), nil
    }
    SegmentassignmentsMarshalled = true

    return json.Marshal(&struct {
        
        Segments []Segmentforassignment `json:"segments"`
        *Alias
    }{

        
        Segments: []Segmentforassignment{{}},
        

        Alias: (*Alias)(u),
    })
}

