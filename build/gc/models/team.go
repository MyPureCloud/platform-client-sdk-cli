package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamDud struct { 
    Id string `json:"id"`


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    MemberCount int `json:"memberCount"`


    SelfUri string `json:"selfUri"`

}

// Team
type Team struct { 
    


    // Name - The team name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - Team information.
    Description string `json:"description"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Team) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Team) MarshalJSON() ([]byte, error) {
    type Alias Team

    if TeamMarshalled {
        return []byte("{}"), nil
    }
    TeamMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

