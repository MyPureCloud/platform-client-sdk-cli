package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FreetriallimitMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FreetriallimitDud struct { 
    


    


    


    

}

// Freetriallimit
type Freetriallimit struct { 
    // Key
    Key string `json:"key"`


    // DefaultValue
    DefaultValue int `json:"defaultValue"`


    // Description
    Description string `json:"description"`


    // Resource
    Resource string `json:"resource"`

}

// String returns a JSON representation of the model
func (o *Freetriallimit) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Freetriallimit) MarshalJSON() ([]byte, error) {
    type Alias Freetriallimit

    if FreetriallimitMarshalled {
        return []byte("{}"), nil
    }
    FreetriallimitMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        DefaultValue int `json:"defaultValue"`
        
        Description string `json:"description"`
        
        Resource string `json:"resource"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

