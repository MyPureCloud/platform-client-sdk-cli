package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestmetricsDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contestmetrics
type Contestmetrics struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Weight - The Contest Metrics weight, allowed values are 1-100
    Weight int `json:"weight"`


    // MinimumQualifier - The Contest Metrics minimum qualifier. Min value is 0, no Max value
    MinimumQualifier int `json:"minimumQualifier"`


    

}

// String returns a JSON representation of the model
func (o *Contestmetrics) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestmetrics) MarshalJSON() ([]byte, error) {
    type Alias Contestmetrics

    if ContestmetricsMarshalled {
        return []byte("{}"), nil
    }
    ContestmetricsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Weight int `json:"weight"`
        
        MinimumQualifier int `json:"minimumQualifier"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

