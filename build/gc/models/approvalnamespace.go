package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ApprovalnamespaceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ApprovalnamespaceDud struct { 
    Namespace string `json:"namespace"`


    Status string `json:"status"`


    VarType string `json:"type"`

}

// Approvalnamespace
type Approvalnamespace struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Approvalnamespace) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Approvalnamespace) MarshalJSON() ([]byte, error) {
    type Alias Approvalnamespace

    if ApprovalnamespaceMarshalled {
        return []byte("{}"), nil
    }
    ApprovalnamespaceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

