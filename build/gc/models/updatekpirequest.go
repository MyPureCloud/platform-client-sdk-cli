package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatekpirequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatekpirequestDud struct { 
    


    


    

}

// Updatekpirequest
type Updatekpirequest struct { 
    // Name - The name of the Key Performance Indicator.
    Name string `json:"name"`


    // Description - The description of the Key Performance Indicator.
    Description string `json:"description"`


    // Status - Key Performance Indicator status which can be Enabled or Disabled.
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Updatekpirequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatekpirequest) MarshalJSON() ([]byte, error) {
    type Alias Updatekpirequest

    if UpdatekpirequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatekpirequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

