package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewelementfilternumberpredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewelementfilternumberpredicateDud struct { 
    


    


    


    

}

// Journeyviewelementfilternumberpredicate - A numeric filter on an element within a journey view
type Journeyviewelementfilternumberpredicate struct { 
    // Dimension - the element's attribute being filtered on
    Dimension string `json:"dimension"`


    // Operator - Optional operator, default is Matches. Valid values: Matches
    Operator string `json:"operator"`


    // NoValue - set this to true if no specific value to be considered
    NoValue bool `json:"noValue"`


    // VarRange - the range of comparators to filter on
    VarRange Journeyviewelementfilterrange `json:"range"`

}

// String returns a JSON representation of the model
func (o *Journeyviewelementfilternumberpredicate) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewelementfilternumberpredicate) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewelementfilternumberpredicate

    if JourneyviewelementfilternumberpredicateMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewelementfilternumberpredicateMarshalled = true

    return json.Marshal(&struct {
        
        Dimension string `json:"dimension"`
        
        Operator string `json:"operator"`
        
        NoValue bool `json:"noValue"`
        
        VarRange Journeyviewelementfilterrange `json:"range"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

