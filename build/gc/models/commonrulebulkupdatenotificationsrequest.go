package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommonrulebulkupdatenotificationsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommonrulebulkupdatenotificationsrequestDud struct { 
    


    

}

// Commonrulebulkupdatenotificationsrequest
type Commonrulebulkupdatenotificationsrequest struct { 
    // RuleIds - The user supplied rules ids to be updated
    RuleIds []string `json:"ruleIds"`


    // Properties - The rule properties to be updated
    Properties Modifiableruleproperties `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Commonrulebulkupdatenotificationsrequest) String() string {
     o.RuleIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commonrulebulkupdatenotificationsrequest) MarshalJSON() ([]byte, error) {
    type Alias Commonrulebulkupdatenotificationsrequest

    if CommonrulebulkupdatenotificationsrequestMarshalled {
        return []byte("{}"), nil
    }
    CommonrulebulkupdatenotificationsrequestMarshalled = true

    return json.Marshal(&struct {
        
        RuleIds []string `json:"ruleIds"`
        
        Properties Modifiableruleproperties `json:"properties"`
        *Alias
    }{

        
        RuleIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

