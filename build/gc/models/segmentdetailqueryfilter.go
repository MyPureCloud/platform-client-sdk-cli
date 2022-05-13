package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentdetailqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentdetailqueryfilterDud struct { 
    


    


    

}

// Segmentdetailqueryfilter
type Segmentdetailqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Segmentdetailqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Segmentdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Segmentdetailqueryfilter) String() string {
    
     o.Clauses = []Segmentdetailqueryclause{{}} 
     o.Predicates = []Segmentdetailquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentdetailqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Segmentdetailqueryfilter

    if SegmentdetailqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    SegmentdetailqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Segmentdetailqueryclause `json:"clauses"`
        
        Predicates []Segmentdetailquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Segmentdetailqueryclause{{}},
        


        
        Predicates: []Segmentdetailquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

