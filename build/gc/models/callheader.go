package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallheaderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallheaderDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Callheader
type Callheader struct { 
    


    // Name
    Name string `json:"name"`


    // Headers - parsed SIP headers
    Headers map[string][]string `json:"headers"`


    

}

// String returns a JSON representation of the model
func (o *Callheader) String() string {
    
     o.Headers = map[string][]string{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callheader) MarshalJSON() ([]byte, error) {
    type Alias Callheader

    if CallheaderMarshalled {
        return []byte("{}"), nil
    }
    CallheaderMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Headers map[string][]string `json:"headers"`
        *Alias
    }{

        


        


        
        Headers: map[string][]string{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

