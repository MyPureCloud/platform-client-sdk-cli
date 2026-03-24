package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsettingsresponseDud struct { 
    

}

// Workitemsettingsresponse
type Workitemsettingsresponse struct { 
    // Worktype - The worktype information for the workitem settings.
    Worktype Stepplansworktypereference `json:"worktype"`

}

// String returns a JSON representation of the model
func (o *Workitemsettingsresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Workitemsettingsresponse

    if WorkitemsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Worktype Stepplansworktypereference `json:"worktype"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

