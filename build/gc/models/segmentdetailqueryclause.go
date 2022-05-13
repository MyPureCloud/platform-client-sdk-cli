package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentdetailqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentdetailqueryclauseDud struct { 
    


    

}

// Segmentdetailqueryclause
type Segmentdetailqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Segmentdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Segmentdetailqueryclause) String() string {
    
     o.Predicates = []Segmentdetailquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentdetailqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Segmentdetailqueryclause

    if SegmentdetailqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    SegmentdetailqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Segmentdetailquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Segmentdetailquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

