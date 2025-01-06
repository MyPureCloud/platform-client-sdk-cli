package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingtemplateheaderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingtemplateheaderDud struct { 
    


    


    

}

// Recordingtemplateheader
type Recordingtemplateheader struct { 
    // VarType - Template header type.
    VarType string `json:"type"`


    // Text - Header text.
    Text string `json:"text"`


    // Media - Media template header image.
    Media Recordingattachment `json:"media"`

}

// String returns a JSON representation of the model
func (o *Recordingtemplateheader) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingtemplateheader) MarshalJSON() ([]byte, error) {
    type Alias Recordingtemplateheader

    if RecordingtemplateheaderMarshalled {
        return []byte("{}"), nil
    }
    RecordingtemplateheaderMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Media Recordingattachment `json:"media"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

