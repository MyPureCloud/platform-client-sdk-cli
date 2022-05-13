package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SharedresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SharedresponseDud struct { 
    


    


    


    


    

}

// Sharedresponse
type Sharedresponse struct { 
    // Id
    Id string `json:"id"`


    // DownloadUri
    DownloadUri string `json:"downloadUri"`


    // ViewUri
    ViewUri string `json:"viewUri"`


    // Document
    Document Document `json:"document"`


    // Share
    Share Share `json:"share"`

}

// String returns a JSON representation of the model
func (o *Sharedresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sharedresponse) MarshalJSON() ([]byte, error) {
    type Alias Sharedresponse

    if SharedresponseMarshalled {
        return []byte("{}"), nil
    }
    SharedresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DownloadUri string `json:"downloadUri"`
        
        ViewUri string `json:"viewUri"`
        
        Document Document `json:"document"`
        
        Share Share `json:"share"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

