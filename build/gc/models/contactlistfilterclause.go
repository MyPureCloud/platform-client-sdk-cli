package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistfilterclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistfilterclauseDud struct { 
    


    

}

// Contactlistfilterclause
type Contactlistfilterclause struct { 
    // FilterType - How to join predicates together.
    FilterType string `json:"filterType"`


    // Predicates - Conditions to filter the contacts by.
    Predicates []Contactlistfilterpredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Contactlistfilterclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Contactlistfilterpredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistfilterclause) MarshalJSON() ([]byte, error) {
    type Alias Contactlistfilterclause

    if ContactlistfilterclauseMarshalled {
        return []byte("{}"), nil
    }
    ContactlistfilterclauseMarshalled = true

    return json.Marshal(&struct { 
        FilterType string `json:"filterType"`
        
        Predicates []Contactlistfilterpredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Contactlistfilterpredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

