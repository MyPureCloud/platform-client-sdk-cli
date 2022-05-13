package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UcintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UcintegrationDud struct { 
    Id string `json:"id"`


    


    UcIntegrationKey string `json:"ucIntegrationKey"`


    IntegrationPresenceSource string `json:"integrationPresenceSource"`


    PbxPermission string `json:"pbxPermission"`


    Icon Ucicon `json:"icon"`


    I10n map[string]Uci10n `json:"i10n"`


    SelfUri string `json:"selfUri"`

}

// Ucintegration - UC Integration UI configuration data
type Ucintegration struct { 
    


    // Name
    Name string `json:"name"`


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Ucintegration) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ucintegration) MarshalJSON() ([]byte, error) {
    type Alias Ucintegration

    if UcintegrationMarshalled {
        return []byte("{}"), nil
    }
    UcintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

