package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserappconfigurationinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserappconfigurationinfoDud struct { 
    Current Integrationconfiguration `json:"current"`


    Effective Effectiveconfiguration `json:"effective"`

}

// Userappconfigurationinfo - Configuration information for the integration
type Userappconfigurationinfo struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Userappconfigurationinfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userappconfigurationinfo) MarshalJSON() ([]byte, error) {
    type Alias Userappconfigurationinfo

    if UserappconfigurationinfoMarshalled {
        return []byte("{}"), nil
    }
    UserappconfigurationinfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

