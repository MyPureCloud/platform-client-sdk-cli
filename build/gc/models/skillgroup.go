package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillgroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillgroupDud struct { 
    Id string `json:"id"`


    


    


    


    MemberCount int `json:"memberCount"`


    DateModified time.Time `json:"dateModified"`


    DateCreated time.Time `json:"dateCreated"`


    Status string `json:"status"`


    


    SelfUri string `json:"selfUri"`

}

// Skillgroup
type Skillgroup struct { 
    


    // Name - The group name.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - Group description
    Description string `json:"description"`


    


    


    


    


    // SkillConditions - Conditions for this group
    SkillConditions []Skillgroupcondition `json:"skillConditions"`


    

}

// String returns a JSON representation of the model
func (o *Skillgroup) String() string {
    
    
    
     o.SkillConditions = []Skillgroupcondition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillgroup) MarshalJSON() ([]byte, error) {
    type Alias Skillgroup

    if SkillgroupMarshalled {
        return []byte("{}"), nil
    }
    SkillgroupMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        SkillConditions []Skillgroupcondition `json:"skillConditions"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        SkillConditions: []Skillgroupcondition{{}},
        


        

        Alias: (*Alias)(u),
    })
}

