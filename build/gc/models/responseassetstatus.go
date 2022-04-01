package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseassetstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseassetstatusDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    ErrorCode string `json:"errorCode"`


    ErrorMessage string `json:"errorMessage"`

}

// Responseassetstatus
type Responseassetstatus struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Responseassetstatus) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseassetstatus) MarshalJSON() ([]byte, error) {
    type Alias Responseassetstatus

    if ResponseassetstatusMarshalled {
        return []byte("{}"), nil
    }
    ResponseassetstatusMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

