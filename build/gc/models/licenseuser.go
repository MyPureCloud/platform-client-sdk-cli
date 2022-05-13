package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LicenseuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LicenseuserDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Licenseuser
type Licenseuser struct { 
    


    // Licenses
    Licenses []Licensedefinition `json:"licenses"`


    

}

// String returns a JSON representation of the model
func (o *Licenseuser) String() string {
     o.Licenses = []Licensedefinition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Licenseuser) MarshalJSON() ([]byte, error) {
    type Alias Licenseuser

    if LicenseuserMarshalled {
        return []byte("{}"), nil
    }
    LicenseuserMarshalled = true

    return json.Marshal(&struct {
        
        Licenses []Licensedefinition `json:"licenses"`
        *Alias
    }{

        


        
        Licenses: []Licensedefinition{{}},
        


        

        Alias: (*Alias)(u),
    })
}

