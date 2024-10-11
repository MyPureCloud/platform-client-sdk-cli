package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatemuagentworkplansbatchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatemuagentworkplansbatchresponseDud struct { 
    

}

// Updatemuagentworkplansbatchresponse
type Updatemuagentworkplansbatchresponse struct { 
    // Failures - The work plan update failures
    Failures []Updatemuagentworkplanfailureresponse `json:"failures"`

}

// String returns a JSON representation of the model
func (o *Updatemuagentworkplansbatchresponse) String() string {
     o.Failures = []Updatemuagentworkplanfailureresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatemuagentworkplansbatchresponse) MarshalJSON() ([]byte, error) {
    type Alias Updatemuagentworkplansbatchresponse

    if UpdatemuagentworkplansbatchresponseMarshalled {
        return []byte("{}"), nil
    }
    UpdatemuagentworkplansbatchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Failures []Updatemuagentworkplanfailureresponse `json:"failures"`
        *Alias
    }{

        
        Failures: []Updatemuagentworkplanfailureresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

