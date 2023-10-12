package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemquerysortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemquerysortDud struct { 
    


    

}

// Workitemquerysort
type Workitemquerysort struct { 
    // Name - Specify an attribute for sorting. It's possible to use an attribute both for sorting and in the query at the same time, but these restrictions apply: Only the operators EQ, LT, LTE, GT, GTE, BETWEEN and BEGINS_WITH are supported and the attribute can't be present in more than one filter.
    Name string `json:"name"`


    // Ascending - Sort Ascending
    Ascending bool `json:"ascending"`

}

// String returns a JSON representation of the model
func (o *Workitemquerysort) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemquerysort) MarshalJSON() ([]byte, error) {
    type Alias Workitemquerysort

    if WorkitemquerysortMarshalled {
        return []byte("{}"), nil
    }
    WorkitemquerysortMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Ascending bool `json:"ascending"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

