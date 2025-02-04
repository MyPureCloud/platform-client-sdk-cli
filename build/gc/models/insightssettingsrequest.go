package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightssettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightssettingsrequestDud struct { 
    

}

// Insightssettingsrequest
type Insightssettingsrequest struct { 
    // Enabled - The AI Insights setting
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Insightssettingsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightssettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Insightssettingsrequest

    if InsightssettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    InsightssettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

