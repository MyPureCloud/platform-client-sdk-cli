package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateskilldivisionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateskilldivisionrequestDud struct { 
    

}

// Updateskilldivisionrequest
type Updateskilldivisionrequest struct { 
    // DivisionId - The division to which this skill will belong
    DivisionId string `json:"divisionId"`

}

// String returns a JSON representation of the model
func (o *Updateskilldivisionrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateskilldivisionrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateskilldivisionrequest

    if UpdateskilldivisionrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateskilldivisionrequestMarshalled = true

    return json.Marshal(&struct {
        
        DivisionId string `json:"divisionId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

