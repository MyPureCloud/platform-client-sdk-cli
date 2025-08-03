package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LongtermrequirementserrordetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LongtermrequirementserrordetailDud struct { 
    


    

}

// Longtermrequirementserrordetail
type Longtermrequirementserrordetail struct { 
    // InternalErrorCode - The error code
    InternalErrorCode string `json:"internalErrorCode"`


    // Description - The description of the error code
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Longtermrequirementserrordetail) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Longtermrequirementserrordetail) MarshalJSON() ([]byte, error) {
    type Alias Longtermrequirementserrordetail

    if LongtermrequirementserrordetailMarshalled {
        return []byte("{}"), nil
    }
    LongtermrequirementserrordetailMarshalled = true

    return json.Marshal(&struct {
        
        InternalErrorCode string `json:"internalErrorCode"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

