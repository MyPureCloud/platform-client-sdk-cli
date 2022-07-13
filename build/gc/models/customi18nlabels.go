package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Customi18nlabelsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Customi18nlabelsDud struct { 
    


    

}

// Customi18nlabels - The localization settings for homescreen
type Customi18nlabels struct { 
    // Language - Language of localized labels in homescreen app (eg. en-us, de-de)
    Language string `json:"language"`


    // LocalizedLabels - Contains localized labels used in homescreen app
    LocalizedLabels []Localizedlabels `json:"localizedLabels"`

}

// String returns a JSON representation of the model
func (o *Customi18nlabels) String() string {
    
     o.LocalizedLabels = []Localizedlabels{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customi18nlabels) MarshalJSON() ([]byte, error) {
    type Alias Customi18nlabels

    if Customi18nlabelsMarshalled {
        return []byte("{}"), nil
    }
    Customi18nlabelsMarshalled = true

    return json.Marshal(&struct {
        
        Language string `json:"language"`
        
        LocalizedLabels []Localizedlabels `json:"localizedLabels"`
        *Alias
    }{

        


        
        LocalizedLabels: []Localizedlabels{{}},
        

        Alias: (*Alias)(u),
    })
}

