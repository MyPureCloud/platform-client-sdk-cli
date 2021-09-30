package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptaggregatequeryfilterDud struct { 
    


    


    

}

// Transcriptaggregatequeryfilter
type Transcriptaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Transcriptaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Transcriptaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Transcriptaggregatequeryclause{{}} 
    
    
    
     o.Predicates = []Transcriptaggregatequerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Transcriptaggregatequeryfilter

    if TranscriptaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    TranscriptaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Transcriptaggregatequeryclause `json:"clauses"`
        
        Predicates []Transcriptaggregatequerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Transcriptaggregatequeryclause{{}},
        

        

        
        Predicates: []Transcriptaggregatequerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

