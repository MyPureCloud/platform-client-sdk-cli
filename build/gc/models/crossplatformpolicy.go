package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformpolicyDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Crossplatformpolicy
type Crossplatformpolicy struct { 
    


    // Name
    Name string `json:"name"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // Order
    Order int `json:"order"`


    // Description
    Description string `json:"description"`


    // Enabled
    Enabled bool `json:"enabled"`


    // MediaPolicies - Conditions and actions per media type
    MediaPolicies Crossplatformmediapolicies `json:"mediaPolicies"`


    // Conditions - Conditions
    Conditions Policyconditions `json:"conditions"`


    // Actions - Actions
    Actions Crossplatformpolicyactions `json:"actions"`


    // PolicyErrors
    PolicyErrors Policyerrors `json:"policyErrors"`


    

}

// String returns a JSON representation of the model
func (o *Crossplatformpolicy) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformpolicy) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformpolicy

    if CrossplatformpolicyMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformpolicyMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        Order int `json:"order"`
        
        Description string `json:"description"`
        
        Enabled bool `json:"enabled"`
        
        MediaPolicies Crossplatformmediapolicies `json:"mediaPolicies"`
        
        Conditions Policyconditions `json:"conditions"`
        
        Actions Crossplatformpolicyactions `json:"actions"`
        
        PolicyErrors Policyerrors `json:"policyErrors"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

