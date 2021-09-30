package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigndivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigndivisionviewDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Campaigndivisionview
type Campaigndivisionview struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Campaigndivisionview) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigndivisionview) MarshalJSON() ([]byte, error) {
    type Alias Campaigndivisionview

    if CampaigndivisionviewMarshalled {
        return []byte("{}"), nil
    }
    CampaigndivisionviewMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

