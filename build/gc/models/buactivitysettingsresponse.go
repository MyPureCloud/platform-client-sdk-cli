package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuactivitysettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuactivitysettingsresponseDud struct { 
    

}

// Buactivitysettingsresponse
type Buactivitysettingsresponse struct { 
    // DefaultActivityCode - Default Activity Code settings
    DefaultActivityCode Activitycodereference `json:"defaultActivityCode"`

}

// String returns a JSON representation of the model
func (o *Buactivitysettingsresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buactivitysettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Buactivitysettingsresponse

    if BuactivitysettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    BuactivitysettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        DefaultActivityCode Activitycodereference `json:"defaultActivityCode"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

