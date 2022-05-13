package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmapactiontemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmapactiontemplateDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Actionmapactiontemplate
type Actionmapactiontemplate struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Actionmapactiontemplate) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmapactiontemplate) MarshalJSON() ([]byte, error) {
    type Alias Actionmapactiontemplate

    if ActionmapactiontemplateMarshalled {
        return []byte("{}"), nil
    }
    ActionmapactiontemplateMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

