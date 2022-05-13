package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformpolicycreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformpolicycreateDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Crossplatformpolicycreate
type Crossplatformpolicycreate struct { 
    


    // Name - The policy name.
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
func (o *Crossplatformpolicycreate) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformpolicycreate) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformpolicycreate

    if CrossplatformpolicycreateMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformpolicycreateMarshalled = true

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

