package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LimitdocsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LimitdocsDud struct { 
    


    


    


    


    


    

}

// Limitdocs
type Limitdocs struct { 
    // Key
    Key string `json:"key"`


    // DefaultValue
    DefaultValue int `json:"defaultValue"`


    // Description
    Description string `json:"description"`


    // Resource
    Resource string `json:"resource"`


    // Configurable
    Configurable bool `json:"configurable"`


    // Trackable
    Trackable bool `json:"trackable"`

}

// String returns a JSON representation of the model
func (o *Limitdocs) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Limitdocs) MarshalJSON() ([]byte, error) {
    type Alias Limitdocs

    if LimitdocsMarshalled {
        return []byte("{}"), nil
    }
    LimitdocsMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        DefaultValue int `json:"defaultValue"`
        
        Description string `json:"description"`
        
        Resource string `json:"resource"`
        
        Configurable bool `json:"configurable"`
        
        Trackable bool `json:"trackable"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

