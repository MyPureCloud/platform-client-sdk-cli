package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionplanninggroupsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionplanninggroupsDud struct { 
    

}

// Sessionplanninggroups
type Sessionplanninggroups struct { 
    // DownloadUrl - URL to fetch the planning groups. This field is populated only if session state is Complete
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Sessionplanninggroups) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionplanninggroups) MarshalJSON() ([]byte, error) {
    type Alias Sessionplanninggroups

    if SessionplanninggroupsMarshalled {
        return []byte("{}"), nil
    }
    SessionplanninggroupsMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

