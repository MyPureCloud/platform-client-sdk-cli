package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RegionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RegionresponseDud struct { 
    

}

// Regionresponse
type Regionresponse struct { 
    // RegionName - Name of the valid linking region, ie. us-east-1
    RegionName string `json:"regionName"`

}

// String returns a JSON representation of the model
func (o *Regionresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Regionresponse) MarshalJSON() ([]byte, error) {
    type Alias Regionresponse

    if RegionresponseMarshalled {
        return []byte("{}"), nil
    }
    RegionresponseMarshalled = true

    return json.Marshal(&struct {
        
        RegionName string `json:"regionName"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

