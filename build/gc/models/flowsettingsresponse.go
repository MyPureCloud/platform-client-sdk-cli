package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowsettingsresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowsettingsresponse - This is a table of settings per a loglevel that define what will be logged in executionData when enabled (true)
type Flowsettingsresponse struct { 
    


    // Name
    Name string `json:"name"`


    // VarType - The Flow Type
    VarType string `json:"type"`


    // ModifiedBy - User that last changed the log level setting.
    ModifiedBy Userreference `json:"modifiedBy"`


    // ModifiedByClient - OAuth client that last changed the log level setting.
    ModifiedByClient Domainentityref `json:"modifiedByClient"`


    // DateModified - The time this log level was set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // LogLevelCharacteristics - The log level set for this flow
    LogLevelCharacteristics Flowloglevel `json:"logLevelCharacteristics"`


    

}

// String returns a JSON representation of the model
func (o *Flowsettingsresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowsettingsresponse

    if FlowsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        ModifiedByClient Domainentityref `json:"modifiedByClient"`
        
        DateModified time.Time `json:"dateModified"`
        
        LogLevelCharacteristics Flowloglevel `json:"logLevelCharacteristics"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

