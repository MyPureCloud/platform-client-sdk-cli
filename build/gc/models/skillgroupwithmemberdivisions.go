package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillgroupwithmemberdivisionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillgroupwithmemberdivisionsDud struct { 
    Id string `json:"id"`


    


    


    


    MemberCount int `json:"memberCount"`


    DateModified time.Time `json:"dateModified"`


    DateCreated time.Time `json:"dateCreated"`


    Status string `json:"status"`


    


    


    SelfUri string `json:"selfUri"`

}

// Skillgroupwithmemberdivisions
type Skillgroupwithmemberdivisions struct { 
    


    // Name - The group name.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - Group description
    Description string `json:"description"`


    


    


    


    


    // SkillConditions - Conditions for this group
    SkillConditions []Skillgroupcondition `json:"skillConditions"`


    // MemberDivisions - Member divisions for this skill group
    MemberDivisions []string `json:"memberDivisions"`


    

}

// String returns a JSON representation of the model
func (o *Skillgroupwithmemberdivisions) String() string {
    
    
    
     o.SkillConditions = []Skillgroupcondition{{}} 
     o.MemberDivisions = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillgroupwithmemberdivisions) MarshalJSON() ([]byte, error) {
    type Alias Skillgroupwithmemberdivisions

    if SkillgroupwithmemberdivisionsMarshalled {
        return []byte("{}"), nil
    }
    SkillgroupwithmemberdivisionsMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        SkillConditions []Skillgroupcondition `json:"skillConditions"`
        
        MemberDivisions []string `json:"memberDivisions"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        SkillConditions: []Skillgroupcondition{{}},
        


        
        MemberDivisions: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

