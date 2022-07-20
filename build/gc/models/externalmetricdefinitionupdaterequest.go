package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalmetricdefinitionupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalmetricdefinitionupdaterequestDud struct { 
    


    


    


    

}

// Externalmetricdefinitionupdaterequest
type Externalmetricdefinitionupdaterequest struct { 
    // Name - The name of the External Metric Definition
    Name string `json:"name"`


    // Precision - The decimal precision of the External Metric Definition. Must be at least 0 and at most 5
    Precision int `json:"precision"`


    // DefaultObjectiveType - The default objective type of the External Metric Definition
    DefaultObjectiveType string `json:"defaultObjectiveType"`


    // Enabled - True if the External Metric Definition is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Externalmetricdefinitionupdaterequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalmetricdefinitionupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Externalmetricdefinitionupdaterequest

    if ExternalmetricdefinitionupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    ExternalmetricdefinitionupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Precision int `json:"precision"`
        
        DefaultObjectiveType string `json:"defaultObjectiveType"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

