package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ModelingstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ModelingstatusresponseDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    ErrorDetails []Modelingprocessingerror `json:"errorDetails"`


    ModelingResultUri string `json:"modelingResultUri"`

}

// Modelingstatusresponse
type Modelingstatusresponse struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Modelingstatusresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Modelingstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Modelingstatusresponse

    if ModelingstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    ModelingstatusresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

