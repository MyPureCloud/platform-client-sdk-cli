package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuideversionresourcesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuideversionresourcesDud struct { 
    

}

// Guideversionresources
type Guideversionresources struct { 
    // DataActions - The data actions associated with this version of the guide.
    DataActions []Dataaction `json:"dataActions"`

}

// String returns a JSON representation of the model
func (o *Guideversionresources) String() string {
     o.DataActions = []Dataaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guideversionresources) MarshalJSON() ([]byte, error) {
    type Alias Guideversionresources

    if GuideversionresourcesMarshalled {
        return []byte("{}"), nil
    }
    GuideversionresourcesMarshalled = true

    return json.Marshal(&struct {
        
        DataActions []Dataaction `json:"dataActions"`
        *Alias
    }{

        
        DataActions: []Dataaction{{}},
        

        Alias: (*Alias)(u),
    })
}

