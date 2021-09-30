package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentthumbnailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentthumbnailDud struct { 
    


    


    


    

}

// Documentthumbnail
type Documentthumbnail struct { 
    // Resolution
    Resolution string `json:"resolution"`


    // ImageUri
    ImageUri string `json:"imageUri"`


    // Height
    Height int `json:"height"`


    // Width
    Width int `json:"width"`

}

// String returns a JSON representation of the model
func (o *Documentthumbnail) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentthumbnail) MarshalJSON() ([]byte, error) {
    type Alias Documentthumbnail

    if DocumentthumbnailMarshalled {
        return []byte("{}"), nil
    }
    DocumentthumbnailMarshalled = true

    return json.Marshal(&struct { 
        Resolution string `json:"resolution"`
        
        ImageUri string `json:"imageUri"`
        
        Height int `json:"height"`
        
        Width int `json:"width"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

