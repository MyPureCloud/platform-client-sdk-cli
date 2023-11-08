package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateverifierrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateverifierrequestDud struct { 
    


    


    

}

// Updateverifierrequest
type Updateverifierrequest struct { 
    // Name - The name of the verifier.
    Name string `json:"name"`


    // Enabled - Indicates whether this verifier will be enabled.
    Enabled bool `json:"enabled"`


    // VarDefault - Indicates whether this will be the default verifier.
    VarDefault bool `json:"default"`

}

// String returns a JSON representation of the model
func (o *Updateverifierrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateverifierrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateverifierrequest

    if UpdateverifierrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateverifierrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        
        VarDefault bool `json:"default"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

