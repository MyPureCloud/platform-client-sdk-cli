package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentfileresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentfileresponseDud struct { 
    


    


    


    


    

}

// Contentfileresponse
type Contentfileresponse struct { 
    // Name - The name of the file
    Name string `json:"name"`


    // VarType - The file format
    VarType string `json:"type"`


    // Checksum - The checksum of the file
    Checksum string `json:"checksum"`


    // Size - The size of the file in bytes
    Size int `json:"size"`


    // ContentUrl - Public download url for content. Needs to be expanded
    ContentUrl string `json:"contentUrl"`

}

// String returns a JSON representation of the model
func (o *Contentfileresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentfileresponse) MarshalJSON() ([]byte, error) {
    type Alias Contentfileresponse

    if ContentfileresponseMarshalled {
        return []byte("{}"), nil
    }
    ContentfileresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Checksum string `json:"checksum"`
        
        Size int `json:"size"`
        
        ContentUrl string `json:"contentUrl"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

