package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DependencycountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DependencycountDud struct { 
    EstimatedCount int `json:"estimatedCount"`

}

// Dependencycount - An estimated count of entities that depend on this entity, including indirect dependencies.
type Dependencycount struct { 
    

}

// String returns a JSON representation of the model
func (o *Dependencycount) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dependencycount) MarshalJSON() ([]byte, error) {
    type Alias Dependencycount

    if DependencycountMarshalled {
        return []byte("{}"), nil
    }
    DependencycountMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

