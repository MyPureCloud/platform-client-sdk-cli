package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeintegrationdefaultsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeintegrationdefaultsettingsDud struct { 
    

}

// Knowledgeintegrationdefaultsettings
type Knowledgeintegrationdefaultsettings struct { 
    // BaseUrl - The default base URL setting for the integration.
    BaseUrl string `json:"baseUrl"`

}

// String returns a JSON representation of the model
func (o *Knowledgeintegrationdefaultsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeintegrationdefaultsettings) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeintegrationdefaultsettings

    if KnowledgeintegrationdefaultsettingsMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeintegrationdefaultsettingsMarshalled = true

    return json.Marshal(&struct {
        
        BaseUrl string `json:"baseUrl"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

