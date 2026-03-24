package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseassetbulkrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseassetbulkrequestDud struct { 
    

}

// Responseassetbulkrequest
type Responseassetbulkrequest struct { 
    // AssetIds - List of Response Asset IDs (max 50)
    AssetIds []string `json:"assetIds"`

}

// String returns a JSON representation of the model
func (o *Responseassetbulkrequest) String() string {
     o.AssetIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseassetbulkrequest) MarshalJSON() ([]byte, error) {
    type Alias Responseassetbulkrequest

    if ResponseassetbulkrequestMarshalled {
        return []byte("{}"), nil
    }
    ResponseassetbulkrequestMarshalled = true

    return json.Marshal(&struct {
        
        AssetIds []string `json:"assetIds"`
        *Alias
    }{

        
        AssetIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

