package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommonrulebulkdeleterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommonrulebulkdeleterequestDud struct { 
    

}

// Commonrulebulkdeleterequest
type Commonrulebulkdeleterequest struct { 
    // RuleIds - The user supplied rule ids to be deleted
    RuleIds []string `json:"ruleIds"`

}

// String returns a JSON representation of the model
func (o *Commonrulebulkdeleterequest) String() string {
     o.RuleIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commonrulebulkdeleterequest) MarshalJSON() ([]byte, error) {
    type Alias Commonrulebulkdeleterequest

    if CommonrulebulkdeleterequestMarshalled {
        return []byte("{}"), nil
    }
    CommonrulebulkdeleterequestMarshalled = true

    return json.Marshal(&struct {
        
        RuleIds []string `json:"ruleIds"`
        *Alias
    }{

        
        RuleIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

