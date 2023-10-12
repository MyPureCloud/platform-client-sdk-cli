package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentassignmenteventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentassignmenteventDud struct { 
    

}

// Segmentassignmentevent
type Segmentassignmentevent struct { 
    // Segment - The segment which was assigned.
    Segment Addressableentityref `json:"segment"`

}

// String returns a JSON representation of the model
func (o *Segmentassignmentevent) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentassignmentevent) MarshalJSON() ([]byte, error) {
    type Alias Segmentassignmentevent

    if SegmentassignmenteventMarshalled {
        return []byte("{}"), nil
    }
    SegmentassignmenteventMarshalled = true

    return json.Marshal(&struct {
        
        Segment Addressableentityref `json:"segment"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

