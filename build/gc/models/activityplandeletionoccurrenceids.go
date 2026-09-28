package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplandeletionoccurrenceidsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplandeletionoccurrenceidsDud struct { 
    

}

// Activityplandeletionoccurrenceids
type Activityplandeletionoccurrenceids struct { 
    // Ids - The occurrence Ids to delete from this activity plan
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Activityplandeletionoccurrenceids) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplandeletionoccurrenceids) MarshalJSON() ([]byte, error) {
    type Alias Activityplandeletionoccurrenceids

    if ActivityplandeletionoccurrenceidsMarshalled {
        return []byte("{}"), nil
    }
    ActivityplandeletionoccurrenceidsMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

