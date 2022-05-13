package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LicenseorgtoggleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LicenseorgtoggleDud struct { 
    


    

}

// Licenseorgtoggle
type Licenseorgtoggle struct { 
    // FeatureName
    FeatureName string `json:"featureName"`


    // Enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Licenseorgtoggle) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Licenseorgtoggle) MarshalJSON() ([]byte, error) {
    type Alias Licenseorgtoggle

    if LicenseorgtoggleMarshalled {
        return []byte("{}"), nil
    }
    LicenseorgtoggleMarshalled = true

    return json.Marshal(&struct {
        
        FeatureName string `json:"featureName"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

