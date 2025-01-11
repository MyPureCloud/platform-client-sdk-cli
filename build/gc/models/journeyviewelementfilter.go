package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewelementfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewelementfilterDud struct { 
    


    


    

}

// Journeyviewelementfilter - A set of filters on an element within a journey view
type Journeyviewelementfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses. Valid values: And
    VarType string `json:"type"`


    // Predicates - predicates
    Predicates []Journeyviewelementfilterpredicate `json:"predicates"`


    // NumberPredicates - numberPredicates
    NumberPredicates []Journeyviewelementfilternumberpredicate `json:"numberPredicates"`

}

// String returns a JSON representation of the model
func (o *Journeyviewelementfilter) String() string {
    
     o.Predicates = []Journeyviewelementfilterpredicate{{}} 
     o.NumberPredicates = []Journeyviewelementfilternumberpredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewelementfilter) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewelementfilter

    if JourneyviewelementfilterMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewelementfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Journeyviewelementfilterpredicate `json:"predicates"`
        
        NumberPredicates []Journeyviewelementfilternumberpredicate `json:"numberPredicates"`
        *Alias
    }{

        


        
        Predicates: []Journeyviewelementfilterpredicate{{}},
        


        
        NumberPredicates: []Journeyviewelementfilternumberpredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

