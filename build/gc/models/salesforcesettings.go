package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SalesforcesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SalesforcesettingsDud struct { 
    


    


    


    

}

// Salesforcesettings
type Salesforcesettings struct { 
    // Channel - Filter source by channel.
    Channel string `json:"channel"`


    // Language - Filter source by language.
    Language string `json:"language"`


    // Categories - Filter source by categories.
    Categories []string `json:"categories"`


    // BaseUrl - The base URL to resources.
    BaseUrl string `json:"baseUrl"`

}

// String returns a JSON representation of the model
func (o *Salesforcesettings) String() string {
    
    
     o.Categories = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Salesforcesettings) MarshalJSON() ([]byte, error) {
    type Alias Salesforcesettings

    if SalesforcesettingsMarshalled {
        return []byte("{}"), nil
    }
    SalesforcesettingsMarshalled = true

    return json.Marshal(&struct {
        
        Channel string `json:"channel"`
        
        Language string `json:"language"`
        
        Categories []string `json:"categories"`
        
        BaseUrl string `json:"baseUrl"`
        *Alias
    }{

        


        


        
        Categories: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

