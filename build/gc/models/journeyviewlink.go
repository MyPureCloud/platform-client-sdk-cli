package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewlinkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewlinkDud struct { 
    


    


    


    


    

}

// Journeyviewlink - A link between elements in a journey view
type Journeyviewlink struct { 
    // Id - The identifier of the element downstream
    Id string `json:"id"`


    // ConstraintWithin - A time constraint on this link, which requires a customer to complete the downstream element within this amount of time to be counted.
    ConstraintWithin Journeyviewlinktimeconstraint `json:"constraintWithin"`


    // ConstraintAfter - A time constraint on this link, which requires a customer must complete the downstream element after this amount of time to be counted.
    ConstraintAfter Journeyviewlinktimeconstraint `json:"constraintAfter"`


    // EventCountType - The type of events that will be counted. Note: Concurrent will override any JourneyViewLinkTimeConstraint. Default is Sequential.
    EventCountType string `json:"eventCountType"`


    // JoinAttributes - Other (secondary) attributes on which this link should join the customers being counted
    JoinAttributes []string `json:"joinAttributes"`

}

// String returns a JSON representation of the model
func (o *Journeyviewlink) String() string {
    
    
    
    
     o.JoinAttributes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewlink) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewlink

    if JourneyviewlinkMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewlinkMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ConstraintWithin Journeyviewlinktimeconstraint `json:"constraintWithin"`
        
        ConstraintAfter Journeyviewlinktimeconstraint `json:"constraintAfter"`
        
        EventCountType string `json:"eventCountType"`
        
        JoinAttributes []string `json:"joinAttributes"`
        *Alias
    }{

        


        


        


        


        
        JoinAttributes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

