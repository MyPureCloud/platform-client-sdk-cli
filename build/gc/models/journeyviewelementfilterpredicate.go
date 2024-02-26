package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewelementfilterpredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewelementfilterpredicateDud struct { 
    


    


    


    

}

// Journeyviewelementfilterpredicate - A filter on an element within a journey view
type Journeyviewelementfilterpredicate struct { 
    // Dimension - the element's attribute being filtered on
    Dimension string `json:"dimension"`


    // Values - the values of the attribute to filter on
    Values []string `json:"values"`


    // Operator - Optional operator, default is Matches. Valid values: Matches
    Operator string `json:"operator"`


    // NoValue - set this to true if no specific value to be considered
    NoValue bool `json:"noValue"`

}

// String returns a JSON representation of the model
func (o *Journeyviewelementfilterpredicate) String() string {
    
     o.Values = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewelementfilterpredicate) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewelementfilterpredicate

    if JourneyviewelementfilterpredicateMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewelementfilterpredicateMarshalled = true

    return json.Marshal(&struct {
        
        Dimension string `json:"dimension"`
        
        Values []string `json:"values"`
        
        Operator string `json:"operator"`
        
        NoValue bool `json:"noValue"`
        *Alias
    }{

        


        
        Values: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

