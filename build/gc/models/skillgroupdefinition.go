package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillgroupdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillgroupdefinitionDud struct { 
    Id string `json:"id"`


    


    


    


    MemberCount int `json:"memberCount"`


    DateModified time.Time `json:"dateModified"`


    DateCreated time.Time `json:"dateCreated"`


    SelfUri string `json:"selfUri"`

}

// Skillgroupdefinition
type Skillgroupdefinition struct { 
    


    // Name - The group name.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - Group description
    Description string `json:"description"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Skillgroupdefinition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillgroupdefinition) MarshalJSON() ([]byte, error) {
    type Alias Skillgroupdefinition

    if SkillgroupdefinitionMarshalled {
        return []byte("{}"), nil
    }
    SkillgroupdefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

