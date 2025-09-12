package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataextractionfileurlMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataextractionfileurlDud struct { 
    


    

}

// Dataextractionfileurl
type Dataextractionfileurl struct { 
    // Id - The unique identifier for the file
    Id string `json:"id"`


    // SignedUrl - Cloudfront URL for the file
    SignedUrl string `json:"signedUrl"`

}

// String returns a JSON representation of the model
func (o *Dataextractionfileurl) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataextractionfileurl) MarshalJSON() ([]byte, error) {
    type Alias Dataextractionfileurl

    if DataextractionfileurlMarshalled {
        return []byte("{}"), nil
    }
    DataextractionfileurlMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SignedUrl string `json:"signedUrl"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

