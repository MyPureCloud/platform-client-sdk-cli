package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BackgroundassistantuserauthenticatedinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BackgroundassistantuserauthenticatedinfoDud struct { 
    

}

// Backgroundassistantuserauthenticatedinfo
type Backgroundassistantuserauthenticatedinfo struct { 
    // BackgroundAssistantId - Id of Genesys Cloud Background Assistant
    BackgroundAssistantId string `json:"backgroundAssistantId"`

}

// String returns a JSON representation of the model
func (o *Backgroundassistantuserauthenticatedinfo) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Backgroundassistantuserauthenticatedinfo) MarshalJSON() ([]byte, error) {
    type Alias Backgroundassistantuserauthenticatedinfo

    if BackgroundassistantuserauthenticatedinfoMarshalled {
        return []byte("{}"), nil
    }
    BackgroundassistantuserauthenticatedinfoMarshalled = true

    return json.Marshal(&struct {
        
        BackgroundAssistantId string `json:"backgroundAssistantId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

