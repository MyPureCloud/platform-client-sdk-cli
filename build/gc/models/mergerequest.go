package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MergerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MergerequestDud struct { 
    


    

}

// Mergerequest
type Mergerequest struct { 
    // SourceContactId - The ID of the source contact for the merge operation
    SourceContactId string `json:"sourceContactId"`


    // TargetContactId - The ID of the target contact for the merge operation
    TargetContactId string `json:"targetContactId"`

}

// String returns a JSON representation of the model
func (o *Mergerequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mergerequest) MarshalJSON() ([]byte, error) {
    type Alias Mergerequest

    if MergerequestMarshalled {
        return []byte("{}"), nil
    }
    MergerequestMarshalled = true

    return json.Marshal(&struct {
        
        SourceContactId string `json:"sourceContactId"`
        
        TargetContactId string `json:"targetContactId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

