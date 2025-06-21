package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuideversionpublishjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuideversionpublishjobrequestDud struct { 
    

}

// Guideversionpublishjobrequest - Request body for publishing a guide version
type Guideversionpublishjobrequest struct { 
    // GuideVersion - The attributes of the guide version to update as part of this publish job.
    GuideVersion Guideversionpublish `json:"guideVersion"`

}

// String returns a JSON representation of the model
func (o *Guideversionpublishjobrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guideversionpublishjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Guideversionpublishjobrequest

    if GuideversionpublishjobrequestMarshalled {
        return []byte("{}"), nil
    }
    GuideversionpublishjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        GuideVersion Guideversionpublish `json:"guideVersion"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

