package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HelpsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HelpsettingsDud struct { 
    


    

}

// Helpsettings
type Helpsettings struct { 
    // Keyword - List of keywords for compliance
    Keyword []string `json:"keyword"`


    // Response - The response configuration for the keywords
    Response Complianceresponse `json:"response"`

}

// String returns a JSON representation of the model
func (o *Helpsettings) String() string {
     o.Keyword = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Helpsettings) MarshalJSON() ([]byte, error) {
    type Alias Helpsettings

    if HelpsettingsMarshalled {
        return []byte("{}"), nil
    }
    HelpsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Keyword []string `json:"keyword"`
        
        Response Complianceresponse `json:"response"`
        *Alias
    }{

        
        Keyword: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

