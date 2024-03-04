package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowcxprojectMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowcxprojectDud struct { 
    


    

}

// Dialogflowcxproject
type Dialogflowcxproject struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Dialogflowcxproject) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowcxproject) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowcxproject

    if DialogflowcxprojectMarshalled {
        return []byte("{}"), nil
    }
    DialogflowcxprojectMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

