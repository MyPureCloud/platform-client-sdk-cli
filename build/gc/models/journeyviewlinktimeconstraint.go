package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewlinktimeconstraintMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewlinktimeconstraintDud struct { 
    


    

}

// Journeyviewlinktimeconstraint
type Journeyviewlinktimeconstraint struct { 
    // Unit - The unit for the link's time constraint
    Unit string `json:"unit"`


    // Value - The value for the link's time constraint
    Value int `json:"value"`

}

// String returns a JSON representation of the model
func (o *Journeyviewlinktimeconstraint) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewlinktimeconstraint) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewlinktimeconstraint

    if JourneyviewlinktimeconstraintMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewlinktimeconstraintMarshalled = true

    return json.Marshal(&struct {
        
        Unit string `json:"unit"`
        
        Value int `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

