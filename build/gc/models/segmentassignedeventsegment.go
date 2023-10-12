package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentassignedeventsegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentassignedeventsegmentDud struct { 
    


    


    


    


    

}

// Segmentassignedeventsegment
type Segmentassignedeventsegment struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // DisplayName - The display name of the segment.
    DisplayName string `json:"displayName"`


    // Version - The version of the segment.
    Version int `json:"version"`


    // Scope - The target entity that a segment applies to.
    Scope string `json:"scope"`

}

// String returns a JSON representation of the model
func (o *Segmentassignedeventsegment) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentassignedeventsegment) MarshalJSON() ([]byte, error) {
    type Alias Segmentassignedeventsegment

    if SegmentassignedeventsegmentMarshalled {
        return []byte("{}"), nil
    }
    SegmentassignedeventsegmentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        DisplayName string `json:"displayName"`
        
        Version int `json:"version"`
        
        Scope string `json:"scope"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

