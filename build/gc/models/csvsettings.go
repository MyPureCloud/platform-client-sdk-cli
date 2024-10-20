package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvsettingsDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Csvsettings
type Csvsettings struct { 
    


    // ExternalSettingsId - Id of the external settings
    ExternalSettingsId string `json:"externalSettingsId"`


    // Mappings - Mappings for the transformation
    Mappings []Csvmappingentry `json:"mappings"`


    

}

// String returns a JSON representation of the model
func (o *Csvsettings) String() string {
    
     o.Mappings = []Csvmappingentry{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvsettings) MarshalJSON() ([]byte, error) {
    type Alias Csvsettings

    if CsvsettingsMarshalled {
        return []byte("{}"), nil
    }
    CsvsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ExternalSettingsId string `json:"externalSettingsId"`
        
        Mappings []Csvmappingentry `json:"mappings"`
        *Alias
    }{

        


        


        
        Mappings: []Csvmappingentry{{}},
        


        

        Alias: (*Alias)(u),
    })
}

