package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatebenefitassessmentjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatebenefitassessmentjobrequestDud struct { 
    

}

// Createbenefitassessmentjobrequest
type Createbenefitassessmentjobrequest struct { 
    // DivisionIds - The list of division ids for routing queues that are to be assessed for Predictive Routing benefit.
    DivisionIds []string `json:"divisionIds"`

}

// String returns a JSON representation of the model
func (o *Createbenefitassessmentjobrequest) String() string {
    
    
     o.DivisionIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createbenefitassessmentjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Createbenefitassessmentjobrequest

    if CreatebenefitassessmentjobrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatebenefitassessmentjobrequestMarshalled = true

    return json.Marshal(&struct { 
        DivisionIds []string `json:"divisionIds"`
        
        *Alias
    }{
        

        
        DivisionIds: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

