package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentassignmentlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentassignmentlistingDud struct { 
    


    


    


    

}

// Segmentassignmentlisting
type Segmentassignmentlisting struct { 
    // Entities
    Entities []Segmentassignment `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Segmentassignmentlisting) String() string {
     o.Entities = []Segmentassignment{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentassignmentlisting) MarshalJSON() ([]byte, error) {
    type Alias Segmentassignmentlisting

    if SegmentassignmentlistingMarshalled {
        return []byte("{}"), nil
    }
    SegmentassignmentlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Segmentassignment `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Segmentassignment{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

