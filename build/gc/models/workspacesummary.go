package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkspacesummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkspacesummaryDud struct { 
    


    

}

// Workspacesummary
type Workspacesummary struct { 
    // TotalDocumentCount
    TotalDocumentCount int `json:"totalDocumentCount"`


    // TotalDocumentByteCount
    TotalDocumentByteCount int `json:"totalDocumentByteCount"`

}

// String returns a JSON representation of the model
func (o *Workspacesummary) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workspacesummary) MarshalJSON() ([]byte, error) {
    type Alias Workspacesummary

    if WorkspacesummaryMarshalled {
        return []byte("{}"), nil
    }
    WorkspacesummaryMarshalled = true

    return json.Marshal(&struct {
        
        TotalDocumentCount int `json:"totalDocumentCount"`
        
        TotalDocumentByteCount int `json:"totalDocumentByteCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

