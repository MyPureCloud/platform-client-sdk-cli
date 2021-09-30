package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FeaturestateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FeaturestateDud struct { 
    

}

// Featurestate
type Featurestate struct { 
    // Enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Featurestate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Featurestate) MarshalJSON() ([]byte, error) {
    type Alias Featurestate

    if FeaturestateMarshalled {
        return []byte("{}"), nil
    }
    FeaturestateMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

