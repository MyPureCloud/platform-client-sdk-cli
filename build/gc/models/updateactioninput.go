package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateactioninputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateactioninputDud struct { 
    


    


    


    

}

// Updateactioninput
type Updateactioninput struct { 
    // Category - Category of action, Can be up to 256 characters long
    Category string `json:"category"`


    // Name - Name of action, Can be up to 256 characters long
    Name string `json:"name"`


    // Config - Configuration to support request and response processing
    Config Actionconfig `json:"config"`


    // Version - Version of this action
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Updateactioninput) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateactioninput) MarshalJSON() ([]byte, error) {
    type Alias Updateactioninput

    if UpdateactioninputMarshalled {
        return []byte("{}"), nil
    }
    UpdateactioninputMarshalled = true

    return json.Marshal(&struct {
        
        Category string `json:"category"`
        
        Name string `json:"name"`
        
        Config Actionconfig `json:"config"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

