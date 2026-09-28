package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplandeletionsessionidsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplandeletionsessionidsDud struct { 
    

}

// Activityplandeletionsessionids
type Activityplandeletionsessionids struct { 
    // Ids - The session Ids to delete from this activity plan occurrence
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Activityplandeletionsessionids) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplandeletionsessionids) MarshalJSON() ([]byte, error) {
    type Alias Activityplandeletionsessionids

    if ActivityplandeletionsessionidsMarshalled {
        return []byte("{}"), nil
    }
    ActivityplandeletionsessionidsMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

