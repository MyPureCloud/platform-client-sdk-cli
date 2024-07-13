package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanjoblistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanjoblistingDud struct { 
    

}

// Activityplanjoblisting
type Activityplanjoblisting struct { 
    // Entities
    Entities []Activityplanjobresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Activityplanjoblisting) String() string {
     o.Entities = []Activityplanjobresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanjoblisting) MarshalJSON() ([]byte, error) {
    type Alias Activityplanjoblisting

    if ActivityplanjoblistingMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanjoblistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Activityplanjobresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Activityplanjobresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

