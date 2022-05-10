package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RegisterarchitectjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RegisterarchitectjobresponseDud struct { 
    Id string `json:"id"`


    PresignedUrl string `json:"presignedUrl"`


    Headers map[string]string `json:"headers"`


    SelfUri string `json:"selfUri"`

}

// Registerarchitectjobresponse
type Registerarchitectjobresponse struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Registerarchitectjobresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Registerarchitectjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Registerarchitectjobresponse

    if RegisterarchitectjobresponseMarshalled {
        return []byte("{}"), nil
    }
    RegisterarchitectjobresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

