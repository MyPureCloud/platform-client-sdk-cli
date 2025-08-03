package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentunassignmentsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentunassignmentsDud struct { 
    

}

// Segmentunassignments
type Segmentunassignments struct { 
    // Segments - The segments to be unassigned.
    Segments []Segmentforunassignment `json:"segments"`

}

// String returns a JSON representation of the model
func (o *Segmentunassignments) String() string {
     o.Segments = []Segmentforunassignment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentunassignments) MarshalJSON() ([]byte, error) {
    type Alias Segmentunassignments

    if SegmentunassignmentsMarshalled {
        return []byte("{}"), nil
    }
    SegmentunassignmentsMarshalled = true

    return json.Marshal(&struct {
        
        Segments []Segmentforunassignment `json:"segments"`
        *Alias
    }{

        
        Segments: []Segmentforunassignment{{}},
        

        Alias: (*Alias)(u),
    })
}

