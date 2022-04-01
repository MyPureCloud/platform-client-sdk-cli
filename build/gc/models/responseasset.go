package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseassetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseassetDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Responseasset
type Responseasset struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // ContentLength - response asset size in bytes
    ContentLength int `json:"contentLength"`


    // ContentLocation - response asset location.
    ContentLocation string `json:"contentLocation"`


    // ContentType - MIME type of response asset
    ContentType string `json:"contentType"`


    // DateCreated - Created date of the response asset. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // CreatedBy - User who created the response asset
    CreatedBy Domainentityref `json:"createdBy"`


    // DateModified - Last modified date of the response asset. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - User who last modified the response asset
    ModifiedBy Domainentityref `json:"modifiedBy"`


    // Responses - Canned responses actively using this asset
    Responses []Domainentityref `json:"responses"`


    

}

// String returns a JSON representation of the model
func (o *Responseasset) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Responses = []Domainentityref{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseasset) MarshalJSON() ([]byte, error) {
    type Alias Responseasset

    if ResponseassetMarshalled {
        return []byte("{}"), nil
    }
    ResponseassetMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        ContentLength int `json:"contentLength"`
        
        ContentLocation string `json:"contentLocation"`
        
        ContentType string `json:"contentType"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy Domainentityref `json:"modifiedBy"`
        
        Responses []Domainentityref `json:"responses"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Responses: []Domainentityref{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

