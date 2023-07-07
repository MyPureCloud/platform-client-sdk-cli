package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcentercategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcentercategoryDud struct { 
    


    


    

}

// Supportcentercategory
type Supportcentercategory struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // Image
    Image Supportcenterimage `json:"image"`

}

// String returns a JSON representation of the model
func (o *Supportcentercategory) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcentercategory) MarshalJSON() ([]byte, error) {
    type Alias Supportcentercategory

    if SupportcentercategoryMarshalled {
        return []byte("{}"), nil
    }
    SupportcentercategoryMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        Image Supportcenterimage `json:"image"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

