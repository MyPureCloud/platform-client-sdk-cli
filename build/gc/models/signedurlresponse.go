package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SignedurlresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SignedurlresponseDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Signedurlresponse
type Signedurlresponse struct { 
    


    // Name
    Name string `json:"name"`


    // Url - Url of the downloaded pcap file
    Url string `json:"url"`


    

}

// String returns a JSON representation of the model
func (o *Signedurlresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Signedurlresponse) MarshalJSON() ([]byte, error) {
    type Alias Signedurlresponse

    if SignedurlresponseMarshalled {
        return []byte("{}"), nil
    }
    SignedurlresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Url string `json:"url"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

