package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationimageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationimageDud struct { 
    


    

}

// Locationimage
type Locationimage struct { 
    // Resolution - Height and/or width of image. ex: 640x480 or x128
    Resolution string `json:"resolution"`


    // ImageUri
    ImageUri string `json:"imageUri"`

}

// String returns a JSON representation of the model
func (o *Locationimage) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Locationimage) MarshalJSON() ([]byte, error) {
    type Alias Locationimage

    if LocationimageMarshalled {
        return []byte("{}"), nil
    }
    LocationimageMarshalled = true

    return json.Marshal(&struct {
        
        Resolution string `json:"resolution"`
        
        ImageUri string `json:"imageUri"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

