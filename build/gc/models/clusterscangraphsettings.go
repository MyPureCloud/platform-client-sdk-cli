package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClusterscangraphsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClusterscangraphsettingsDud struct { 
    

}

// Clusterscangraphsettings
type Clusterscangraphsettings struct { 
    // AutomaticMergingEnabled - Whether to enable automatic merging of discovered clusters
    AutomaticMergingEnabled bool `json:"automaticMergingEnabled"`

}

// String returns a JSON representation of the model
func (o *Clusterscangraphsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clusterscangraphsettings) MarshalJSON() ([]byte, error) {
    type Alias Clusterscangraphsettings

    if ClusterscangraphsettingsMarshalled {
        return []byte("{}"), nil
    }
    ClusterscangraphsettingsMarshalled = true

    return json.Marshal(&struct {
        
        AutomaticMergingEnabled bool `json:"automaticMergingEnabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

