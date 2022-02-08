package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportedcontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportedcontentDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    CreatedBy Domainentityref `json:"createdBy"`


    ModifiedBy Domainentityref `json:"modifiedBy"`


    Version int `json:"version"`


    


    SelfUri string `json:"selfUri"`

}

// Supportedcontent - Supported content profile for inbound and outbound messages
type Supportedcontent struct { 
    


    // Name - The name of the supported content profile
    Name string `json:"name"`


    


    


    


    


    


    // MediaTypes - Defines the allowable media that may be accepted for an inbound message or to be sent in an outbound message. The following is an example of allowing all inbound media, and for outbound all images and only mpeg video: {   \"mediaTypes\": {     \"allow\": {       \"inbound\": [{\"type\": \"*/*\"}],       \"outbound\": [{\"type\": \"image/*\"}, {\"type\": \"video/mpeg\"}]     }   } }
    MediaTypes Mediatypes `json:"mediaTypes"`


    

}

// String returns a JSON representation of the model
func (o *Supportedcontent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportedcontent) MarshalJSON() ([]byte, error) {
    type Alias Supportedcontent

    if SupportedcontentMarshalled {
        return []byte("{}"), nil
    }
    SupportedcontentMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        
        
        
        
        
        
        MediaTypes Mediatypes `json:"mediaTypes"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

