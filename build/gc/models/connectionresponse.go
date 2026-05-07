package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectionresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Connectionresponse
type Connectionresponse struct { 
    


    // Name - The name of the connection.
    Name string `json:"name"`


    // VarType - Type of the connection.
    VarType string `json:"type"`


    // Integration - The reference to the integration associated with the connection.
    Integration Knowledgeintegrationreference `json:"integration"`


    // AuthenticationProperties - Authentication properties which can be used to initiate authentication of a user.
    AuthenticationProperties Authenticationproperties `json:"authenticationProperties"`


    // CreatedBy - Reference of the creator.
    CreatedBy Userreference `json:"createdBy"`


    // ModifiedBy - Reference of the user whom modified the connection.
    ModifiedBy Userreference `json:"modifiedBy"`


    // DateCreated - Date of the creation of connection. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date of the last modification made to the connection. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Status - Current status of the connection.
    Status string `json:"status"`


    // VarError - Optional error message of the connection.
    VarError Errorbody `json:"error"`


    // DateExpiry - Expiry date of the authentication credentials. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpiry time.Time `json:"dateExpiry"`


    

}

// String returns a JSON representation of the model
func (o *Connectionresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectionresponse) MarshalJSON() ([]byte, error) {
    type Alias Connectionresponse

    if ConnectionresponseMarshalled {
        return []byte("{}"), nil
    }
    ConnectionresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Integration Knowledgeintegrationreference `json:"integration"`
        
        AuthenticationProperties Authenticationproperties `json:"authenticationProperties"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Status string `json:"status"`
        
        VarError Errorbody `json:"error"`
        
        DateExpiry time.Time `json:"dateExpiry"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

