package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RollbackdecisiontableversionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RollbackdecisiontableversionrequestDud struct { 
    

}

// Rollbackdecisiontableversionrequest
type Rollbackdecisiontableversionrequest struct { 
    // RollbackReason - Optional note recorded on the target version when rollback succeeds. Present while the version is Published after rollback; cleared when Superseded. Max: 200 characters.
    RollbackReason string `json:"rollbackReason"`

}

// String returns a JSON representation of the model
func (o *Rollbackdecisiontableversionrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Rollbackdecisiontableversionrequest) MarshalJSON() ([]byte, error) {
    type Alias Rollbackdecisiontableversionrequest

    if RollbackdecisiontableversionrequestMarshalled {
        return []byte("{}"), nil
    }
    RollbackdecisiontableversionrequestMarshalled = true

    return json.Marshal(&struct {
        
        RollbackReason string `json:"rollbackReason"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

