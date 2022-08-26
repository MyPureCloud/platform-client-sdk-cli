package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabelupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabelupdaterequestDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Labelupdaterequest
type Labelupdaterequest struct { 
    


    // Name - The name of the label.
    Name string `json:"name"`


    // Color - The color for the label.
    Color string `json:"color"`


    

}

// String returns a JSON representation of the model
func (o *Labelupdaterequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Labelupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Labelupdaterequest

    if LabelupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    LabelupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Color string `json:"color"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

