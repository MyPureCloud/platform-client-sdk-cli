package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustrequestDud struct { 
    Id string `json:"id"`


    CreatedBy Orguser `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    Trustee Organization `json:"trustee"`


    Users []Orguser `json:"users"`


    Groups []Trustgroup `json:"groups"`


    SelfUri string `json:"selfUri"`

}

// Trustrequest
type Trustrequest struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Trustrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustrequest) MarshalJSON() ([]byte, error) {
    type Alias Trustrequest

    if TrustrequestMarshalled {
        return []byte("{}"), nil
    }
    TrustrequestMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

