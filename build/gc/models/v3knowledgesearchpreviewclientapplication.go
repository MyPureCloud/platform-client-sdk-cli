package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3knowledgesearchpreviewclientapplicationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3knowledgesearchpreviewclientapplicationDud struct { 
    

}

// V3knowledgesearchpreviewclientapplication
type V3knowledgesearchpreviewclientapplication struct { 
    // VarType - The application type to simulate for the preview.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *V3knowledgesearchpreviewclientapplication) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3knowledgesearchpreviewclientapplication) MarshalJSON() ([]byte, error) {
    type Alias V3knowledgesearchpreviewclientapplication

    if V3knowledgesearchpreviewclientapplicationMarshalled {
        return []byte("{}"), nil
    }
    V3knowledgesearchpreviewclientapplicationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

