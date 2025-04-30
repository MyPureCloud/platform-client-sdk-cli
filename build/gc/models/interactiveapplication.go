package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InteractiveapplicationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InteractiveapplicationDud struct { 
    


    

}

// Interactiveapplication
type Interactiveapplication struct { 
    // Name - The name of the message app.
    Name string `json:"name"`


    // Url - Contains the data that is sent to the message app.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Interactiveapplication) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Interactiveapplication) MarshalJSON() ([]byte, error) {
    type Alias Interactiveapplication

    if InteractiveapplicationMarshalled {
        return []byte("{}"), nil
    }
    InteractiveapplicationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

