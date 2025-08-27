package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EscalationruleresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EscalationruleresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Escalationruleresponse
type Escalationruleresponse struct { 
    // Id - ID of the escalation rule.
    Id string `json:"id"`


    // Name - The name of the escalation rule.
    Name string `json:"name"`


    // MatchCriteria - The criteria that defines when a social media message should be escalated.
    MatchCriteria string `json:"matchCriteria"`


    // Priority - The priority of the escalation rule.
    Priority int `json:"priority"`


    // DivisionId - The ID of the division the social escalation rule belongs to.
    DivisionId string `json:"divisionId"`


    // Description - A description of the social escalation rule.
    Description string `json:"description"`


    // DateCreated - Timestamp indicating when the escalation rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Timestamp indicating when the escalation rule was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


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


    

}

// String returns a JSON representation of the model
func (o *Escalationruleresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Escalationruleresponse) MarshalJSON() ([]byte, error) {
    type Alias Escalationruleresponse

    if EscalationruleresponseMarshalled {
        return []byte("{}"), nil
    }
    EscalationruleresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        MatchCriteria string `json:"matchCriteria"`
        
        Priority int `json:"priority"`
        
        DivisionId string `json:"divisionId"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Status string `json:"status"`
        
        OpenEscalation Escalationtarget `json:"openEscalation"`
        
        FacebookEscalation Escalationtarget `json:"facebookEscalation"`
        
        InstagramEscalation Escalationtarget `json:"instagramEscalation"`
        
        TwitterEscalation Escalationtarget `json:"twitterEscalation"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

