package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LongtermrequirementsservicegoaldetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LongtermrequirementsservicegoaldetailDud struct { 
    


    

}

// Longtermrequirementsservicegoaldetail
type Longtermrequirementsservicegoaldetail struct { 
    // Id - Service goal template ID
    Id string `json:"id"`


    // ServiceGoals - The service goals used to generate the requirements 
    ServiceGoals Longtermrequirementsservicegoal `json:"serviceGoals"`

}

// String returns a JSON representation of the model
func (o *Longtermrequirementsservicegoaldetail) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Longtermrequirementsservicegoaldetail) MarshalJSON() ([]byte, error) {
    type Alias Longtermrequirementsservicegoaldetail

    if LongtermrequirementsservicegoaldetailMarshalled {
        return []byte("{}"), nil
    }
    LongtermrequirementsservicegoaldetailMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ServiceGoals Longtermrequirementsservicegoal `json:"serviceGoals"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

