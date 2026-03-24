package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcewitherrorresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcewitherrorresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// V3sourcewitherrorresponse
type V3sourcewitherrorresponse struct { 
    


    // Name - The name of the source.
    Name string `json:"name"`


    // ConnectionId - The connectionId of the source.
    ConnectionId string `json:"connectionId"`


    // VarType - The type of the source.
    VarType string `json:"type"`


    // TriggerType - The trigger type of the source.
    TriggerType string `json:"triggerType"`


    // Status - The current status of the source.
    Status string `json:"status"`


    // CreatedBy - The user who created the source.
    CreatedBy Userreference `json:"createdBy"`


    // ModifiedBy - The user who modified the document.
    ModifiedBy Userreference `json:"modifiedBy"`


    // DateCreated - Source creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Source last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // LastSync - The last synchronization of the source.
    LastSync V3sourcelastsynchronization `json:"lastSync"`


    // VarError - Optional error details of an errored source.
    VarError Errorbody `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *V3sourcewitherrorresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcewitherrorresponse) MarshalJSON() ([]byte, error) {
    type Alias V3sourcewitherrorresponse

    if V3sourcewitherrorresponseMarshalled {
        return []byte("{}"), nil
    }
    V3sourcewitherrorresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ConnectionId string `json:"connectionId"`
        
        VarType string `json:"type"`
        
        TriggerType string `json:"triggerType"`
        
        Status string `json:"status"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        LastSync V3sourcelastsynchronization `json:"lastSync"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

