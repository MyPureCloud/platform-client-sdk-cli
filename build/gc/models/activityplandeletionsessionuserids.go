package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplandeletionsessionuseridsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplandeletionsessionuseridsDud struct { 
    

}

// Activityplandeletionsessionuserids
type Activityplandeletionsessionuserids struct { 
    // Ids - The user Ids to delete from this activity plan session
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Activityplandeletionsessionuserids) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplandeletionsessionuserids) MarshalJSON() ([]byte, error) {
    type Alias Activityplandeletionsessionuserids

    if ActivityplandeletionsessionuseridsMarshalled {
        return []byte("{}"), nil
    }
    ActivityplandeletionsessionuseridsMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

