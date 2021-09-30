package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillstoremoveMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillstoremoveDud struct { 
    


    


    

}

// Skillstoremove
type Skillstoremove struct { 
    // Name
    Name string `json:"name"`


    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Skillstoremove) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillstoremove) MarshalJSON() ([]byte, error) {
    type Alias Skillstoremove

    if SkillstoremoveMarshalled {
        return []byte("{}"), nil
    }
    SkillstoremoveMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

