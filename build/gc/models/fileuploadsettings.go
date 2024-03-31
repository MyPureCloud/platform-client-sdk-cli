package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FileuploadsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FileuploadsettingsDud struct { 
    


    

}

// Fileuploadsettings - File upload settings for messenger
type Fileuploadsettings struct { 
    // EnableAttachments - whether or not attachments is enabled
    EnableAttachments bool `json:"enableAttachments"`


    // Modes - The list of supported file upload modes
    Modes []Fileuploadmode `json:"modes"`

}

// String returns a JSON representation of the model
func (o *Fileuploadsettings) String() string {
    
     o.Modes = []Fileuploadmode{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fileuploadsettings) MarshalJSON() ([]byte, error) {
    type Alias Fileuploadsettings

    if FileuploadsettingsMarshalled {
        return []byte("{}"), nil
    }
    FileuploadsettingsMarshalled = true

    return json.Marshal(&struct {
        
        EnableAttachments bool `json:"enableAttachments"`
        
        Modes []Fileuploadmode `json:"modes"`
        *Alias
    }{

        


        
        Modes: []Fileuploadmode{{}},
        

        Alias: (*Alias)(u),
    })
}

