package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalmetricdefinitioncreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalmetricdefinitioncreaterequestDud struct { 
    


    


    


    


    


    

}

// Externalmetricdefinitioncreaterequest
type Externalmetricdefinitioncreaterequest struct { 
    // Name - The name of the External Metric Definition
    Name string `json:"name"`


    // Unit - The unit of the External Metric Definition
    Unit string `json:"unit"`


    // UnitDefinition - The unit definition of the External Metric Definition
    UnitDefinition string `json:"unitDefinition"`


    // Precision - The decimal precision of the External Metric Definition. Must be at least 0 and at most 5
    Precision int `json:"precision"`


    // DefaultObjectiveType - The default objective type of the External Metric Definition
    DefaultObjectiveType string `json:"defaultObjectiveType"`


    // Enabled - True if the External Metric Definition is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Externalmetricdefinitioncreaterequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalmetricdefinitioncreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Externalmetricdefinitioncreaterequest

    if ExternalmetricdefinitioncreaterequestMarshalled {
        return []byte("{}"), nil
    }
    ExternalmetricdefinitioncreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Unit string `json:"unit"`
        
        UnitDefinition string `json:"unitDefinition"`
        
        Precision int `json:"precision"`
        
        DefaultObjectiveType string `json:"defaultObjectiveType"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

