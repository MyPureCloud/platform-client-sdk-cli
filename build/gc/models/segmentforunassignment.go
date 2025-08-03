package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentforunassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentforunassignmentDud struct { 
    

}

// Segmentforunassignment
type Segmentforunassignment struct { 
    // Id - The ID of the segment to be unassigned.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Segmentforunassignment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentforunassignment) MarshalJSON() ([]byte, error) {
    type Alias Segmentforunassignment

    if SegmentforunassignmentMarshalled {
        return []byte("{}"), nil
    }
    SegmentforunassignmentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

