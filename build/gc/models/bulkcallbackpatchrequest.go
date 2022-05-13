package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkcallbackpatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkcallbackpatchrequestDud struct { 
    

}

// Bulkcallbackpatchrequest
type Bulkcallbackpatchrequest struct { 
    // PatchCallbackRequests - The list of requests to update callbacks in bulk
    PatchCallbackRequests []Patchcallbackrequest `json:"patchCallbackRequests"`

}

// String returns a JSON representation of the model
func (o *Bulkcallbackpatchrequest) String() string {
     o.PatchCallbackRequests = []Patchcallbackrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkcallbackpatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkcallbackpatchrequest

    if BulkcallbackpatchrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkcallbackpatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        PatchCallbackRequests []Patchcallbackrequest `json:"patchCallbackRequests"`
        *Alias
    }{

        
        PatchCallbackRequests: []Patchcallbackrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

