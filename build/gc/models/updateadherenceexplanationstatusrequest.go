package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateadherenceexplanationstatusrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateadherenceexplanationstatusrequestDud struct { 
    

}

// Updateadherenceexplanationstatusrequest
type Updateadherenceexplanationstatusrequest struct { 
    // Status - The status of the adherence explanation
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Updateadherenceexplanationstatusrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateadherenceexplanationstatusrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateadherenceexplanationstatusrequest

    if UpdateadherenceexplanationstatusrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateadherenceexplanationstatusrequestMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

