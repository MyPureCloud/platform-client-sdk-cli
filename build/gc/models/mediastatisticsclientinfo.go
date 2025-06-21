package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediastatisticsclientinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediastatisticsclientinfoDud struct { 
    


    


    

}

// Mediastatisticsclientinfo
type Mediastatisticsclientinfo struct { 
    // OriginAppName - The name of the application that created the media session.
    OriginAppName string `json:"originAppName"`


    // OriginAppId - The ID of the application that created the media session.
    OriginAppId string `json:"originAppId"`


    // OriginAppVersion - The version of the application that created the media session.
    OriginAppVersion string `json:"originAppVersion"`

}

// String returns a JSON representation of the model
func (o *Mediastatisticsclientinfo) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediastatisticsclientinfo) MarshalJSON() ([]byte, error) {
    type Alias Mediastatisticsclientinfo

    if MediastatisticsclientinfoMarshalled {
        return []byte("{}"), nil
    }
    MediastatisticsclientinfoMarshalled = true

    return json.Marshal(&struct {
        
        OriginAppName string `json:"originAppName"`
        
        OriginAppId string `json:"originAppId"`
        
        OriginAppVersion string `json:"originAppVersion"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

