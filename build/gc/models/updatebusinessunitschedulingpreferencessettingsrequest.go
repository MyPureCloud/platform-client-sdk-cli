package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatebusinessunitschedulingpreferencessettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatebusinessunitschedulingpreferencessettingsrequestDud struct { 
    

}

// Updatebusinessunitschedulingpreferencessettingsrequest
type Updatebusinessunitschedulingpreferencessettingsrequest struct { 
    // Enabled - Whether scheduling preferences are enabled for the business unit
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Updatebusinessunitschedulingpreferencessettingsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatebusinessunitschedulingpreferencessettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatebusinessunitschedulingpreferencessettingsrequest

    if UpdatebusinessunitschedulingpreferencessettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatebusinessunitschedulingpreferencessettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

