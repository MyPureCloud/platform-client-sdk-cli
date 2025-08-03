package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentforassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentforassignmentDud struct { 
    

}

// Segmentforassignment
type Segmentforassignment struct { 
    // Id - The ID of the segment to be assigned.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Segmentforassignment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentforassignment) MarshalJSON() ([]byte, error) {
    type Alias Segmentforassignment

    if SegmentforassignmentMarshalled {
        return []byte("{}"), nil
    }
    SegmentforassignmentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

