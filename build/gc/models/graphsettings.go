package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GraphsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GraphsettingsDud struct { 
    


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Graphsettings
type Graphsettings struct { 
    // AutomaticMergingEnabled - Whether to enable automatic merging of discovered clusters
    AutomaticMergingEnabled bool `json:"automaticMergingEnabled"`


    


    

}

// String returns a JSON representation of the model
func (o *Graphsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Graphsettings) MarshalJSON() ([]byte, error) {
    type Alias Graphsettings

    if GraphsettingsMarshalled {
        return []byte("{}"), nil
    }
    GraphsettingsMarshalled = true

    return json.Marshal(&struct {
        
        AutomaticMergingEnabled bool `json:"automaticMergingEnabled"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

