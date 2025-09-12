package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RemoveentitiesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RemoveentitiesrequestDud struct { 
    

}

// Removeentitiesrequest
type Removeentitiesrequest struct { 
    // Entities - List of entities to be removed
    Entities []Removeentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Removeentitiesrequest) String() string {
     o.Entities = []Removeentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Removeentitiesrequest) MarshalJSON() ([]byte, error) {
    type Alias Removeentitiesrequest

    if RemoveentitiesrequestMarshalled {
        return []byte("{}"), nil
    }
    RemoveentitiesrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Removeentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Removeentity{{}},
        

        Alias: (*Alias)(u),
    })
}

