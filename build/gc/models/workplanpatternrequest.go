package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanpatternrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanpatternrequestDud struct { 
    

}

// Workplanpatternrequest
type Workplanpatternrequest struct { 
    // WorkPlanIds - List of work plan IDs in order of rotation on a weekly basis. Values in the list cannot be null or empty
    WorkPlanIds []string `json:"workPlanIds"`

}

// String returns a JSON representation of the model
func (o *Workplanpatternrequest) String() string {
     o.WorkPlanIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanpatternrequest) MarshalJSON() ([]byte, error) {
    type Alias Workplanpatternrequest

    if WorkplanpatternrequestMarshalled {
        return []byte("{}"), nil
    }
    WorkplanpatternrequestMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlanIds []string `json:"workPlanIds"`
        *Alias
    }{

        
        WorkPlanIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

