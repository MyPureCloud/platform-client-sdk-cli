package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowhealthintentinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowhealthintentinfoDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Flowhealthintentinfo
type Flowhealthintentinfo struct { 
    


    // Name
    Name string `json:"name"`


    // LanguageHealth - Health computation info for each language.
    LanguageHealth map[string]Localehealth `json:"languageHealth"`


    

}

// String returns a JSON representation of the model
func (o *Flowhealthintentinfo) String() string {
    
     o.LanguageHealth = map[string]Localehealth{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowhealthintentinfo) MarshalJSON() ([]byte, error) {
    type Alias Flowhealthintentinfo

    if FlowhealthintentinfoMarshalled {
        return []byte("{}"), nil
    }
    FlowhealthintentinfoMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        LanguageHealth map[string]Localehealth `json:"languageHealth"`
        *Alias
    }{

        


        


        
        LanguageHealth: map[string]Localehealth{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

