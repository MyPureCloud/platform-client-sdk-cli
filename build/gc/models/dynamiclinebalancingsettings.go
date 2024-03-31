package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamiclinebalancingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamiclinebalancingsettingsDud struct { 
    


    

}

// Dynamiclinebalancingsettings
type Dynamiclinebalancingsettings struct { 
    // Enabled - Indicates that this campaign is subject of dynamic line balancing
    Enabled bool `json:"enabled"`


    // RelativeWeight - Relative weight of this campaign in dynamic line balancing
    RelativeWeight int `json:"relativeWeight"`

}

// String returns a JSON representation of the model
func (o *Dynamiclinebalancingsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamiclinebalancingsettings) MarshalJSON() ([]byte, error) {
    type Alias Dynamiclinebalancingsettings

    if DynamiclinebalancingsettingsMarshalled {
        return []byte("{}"), nil
    }
    DynamiclinebalancingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        RelativeWeight int `json:"relativeWeight"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

