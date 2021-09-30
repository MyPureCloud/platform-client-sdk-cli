package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RunnowresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RunnowresponseDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Runnowresponse
type Runnowresponse struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Runnowresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Runnowresponse) MarshalJSON() ([]byte, error) {
    type Alias Runnowresponse

    if RunnowresponseMarshalled {
        return []byte("{}"), nil
    }
    RunnowresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

