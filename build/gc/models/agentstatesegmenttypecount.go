package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstatesegmenttypecountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstatesegmenttypecountDud struct { 
    


    

}

// Agentstatesegmenttypecount
type Agentstatesegmenttypecount struct { 
    // SegmentType - Segment type
    SegmentType string `json:"segmentType"`


    // Count - Count of segment type
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Agentstatesegmenttypecount) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstatesegmenttypecount) MarshalJSON() ([]byte, error) {
    type Alias Agentstatesegmenttypecount

    if AgentstatesegmenttypecountMarshalled {
        return []byte("{}"), nil
    }
    AgentstatesegmenttypecountMarshalled = true

    return json.Marshal(&struct {
        
        SegmentType string `json:"segmentType"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

