package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3startmanualsyncrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3startmanualsyncrequestDud struct { 
    

}

// V3startmanualsyncrequest
type V3startmanualsyncrequest struct { 
    // VarType -  Optional field that specifies the synchronization type. For SharePoint only Full synchronization is supported, therefore that is the default. For FileUpload both Full and Incremental is supported, default is Incremental.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *V3startmanualsyncrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3startmanualsyncrequest) MarshalJSON() ([]byte, error) {
    type Alias V3startmanualsyncrequest

    if V3startmanualsyncrequestMarshalled {
        return []byte("{}"), nil
    }
    V3startmanualsyncrequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

