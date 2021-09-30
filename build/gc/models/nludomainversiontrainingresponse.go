package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludomainversiontrainingresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludomainversiontrainingresponseDud struct { 
    Message string `json:"message"`


    

}

// Nludomainversiontrainingresponse
type Nludomainversiontrainingresponse struct { 
    


    // Version
    Version Nludomainversion `json:"version"`

}

// String returns a JSON representation of the model
func (o *Nludomainversiontrainingresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludomainversiontrainingresponse) MarshalJSON() ([]byte, error) {
    type Alias Nludomainversiontrainingresponse

    if NludomainversiontrainingresponseMarshalled {
        return []byte("{}"), nil
    }
    NludomainversiontrainingresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Version Nludomainversion `json:"version"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

