package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillreferenceDud struct { 
    


    


    

}

// Skillreference - A skill reference extracted from the expression
type Skillreference struct { 
    // Id - Skill UUID
    Id string `json:"id"`


    // Name - Skill name
    Name string `json:"name"`


    // IsLanguageSkill - Whether this is a language skill
    IsLanguageSkill bool `json:"isLanguageSkill"`

}

// String returns a JSON representation of the model
func (o *Skillreference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillreference) MarshalJSON() ([]byte, error) {
    type Alias Skillreference

    if SkillreferenceMarshalled {
        return []byte("{}"), nil
    }
    SkillreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        IsLanguageSkill bool `json:"isLanguageSkill"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

