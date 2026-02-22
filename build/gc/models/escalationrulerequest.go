package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EscalationrulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EscalationrulerequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Escalationrulerequest
type Escalationrulerequest struct { 
    // Name - The name of the escalation rule.
    Name string `json:"name"`


    // MatchCriteria - The criteria that defines when a social media message should be escalated.
    MatchCriteria string `json:"matchCriteria"`


    // Priority - The priority of the escalation rule. The lower the number the higer the priority. Once a rule is matched others are skipped.
    Priority int `json:"priority"`


    // DivisionId - The ID of the division the social escalation rule belongs to.
    DivisionId string `json:"divisionId"`


    // Description - A description of the social escalation rule.
    Description string `json:"description"`


    // Status - The status of the escalation rule.
    Status string `json:"status"`


    // OpenEscalation - The target integration configuration used for an open message escalation.
    OpenEscalation Escalationtarget `json:"openEscalation"`


    // FacebookEscalation - The target integration configuration used for a Facebook message escalation.
    FacebookEscalation Escalationtarget `json:"facebookEscalation"`


    // InstagramEscalation - The target integration configuration used for an Instagram message escalation.
    InstagramEscalation Escalationtarget `json:"instagramEscalation"`


    // TwitterEscalation - The target integration configuration used for a X (formerly Twitter) message escalation.
    TwitterEscalation Escalationtarget `json:"twitterEscalation"`


    // GoogleBusinessProfileEscalation - The target integration configuration used for a Google Business Profile message escalation.
    GoogleBusinessProfileEscalation Escalationtarget `json:"googleBusinessProfileEscalation"`

}

// String returns a JSON representation of the model
func (o *Escalationrulerequest) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Escalationrulerequest) MarshalJSON() ([]byte, error) {
    type Alias Escalationrulerequest

    if EscalationrulerequestMarshalled {
        return []byte("{}"), nil
    }
    EscalationrulerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        MatchCriteria string `json:"matchCriteria"`
        
        Priority int `json:"priority"`
        
        DivisionId string `json:"divisionId"`
        
        Description string `json:"description"`
        
        Status string `json:"status"`
        
        OpenEscalation Escalationtarget `json:"openEscalation"`
        
        FacebookEscalation Escalationtarget `json:"facebookEscalation"`
        
        InstagramEscalation Escalationtarget `json:"instagramEscalation"`
        
        TwitterEscalation Escalationtarget `json:"twitterEscalation"`
        
        GoogleBusinessProfileEscalation Escalationtarget `json:"googleBusinessProfileEscalation"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

