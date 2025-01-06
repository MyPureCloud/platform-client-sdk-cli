package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstatequeryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstatequeryrequestDud struct { 
    


    


    


    


    


    

}

// Agentstatequeryrequest
type Agentstatequeryrequest struct { 
    // UserFilter - Filters that target user-level data
    UserFilter Agentstateuserfilter `json:"userFilter"`


    // SessionFilter - Filters that target session-level data
    SessionFilter Agentstatesessionfilter `json:"sessionFilter"`


    // UserOrderBy - Search user order dimension names; default to userName
    UserOrderBy string `json:"userOrderBy"`


    // UserOrder - Search user order direction; default to asc
    UserOrder string `json:"userOrder"`


    // SessionOrderBy - Search session order dimension names; default to segmentStart
    SessionOrderBy string `json:"sessionOrderBy"`


    // SessionOrder - Search session order direction; default to desc
    SessionOrder string `json:"sessionOrder"`

}

// String returns a JSON representation of the model
func (o *Agentstatequeryrequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstatequeryrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentstatequeryrequest

    if AgentstatequeryrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentstatequeryrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserFilter Agentstateuserfilter `json:"userFilter"`
        
        SessionFilter Agentstatesessionfilter `json:"sessionFilter"`
        
        UserOrderBy string `json:"userOrderBy"`
        
        UserOrder string `json:"userOrder"`
        
        SessionOrderBy string `json:"sessionOrderBy"`
        
        SessionOrder string `json:"sessionOrder"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

