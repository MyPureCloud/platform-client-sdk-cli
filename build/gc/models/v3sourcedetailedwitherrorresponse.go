package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcedetailedwitherrorresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcedetailedwitherrorresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// V3sourcedetailedwitherrorresponse
type V3sourcedetailedwitherrorresponse struct { 
    


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


    // ScheduleSettings - Settings that determine when the source starts a sync.
    ScheduleSettings V3sourceschedulesettings `json:"scheduleSettings"`


    // Filters - Filters that determine what documents are synced.
    Filters V3sourcefilter `json:"filters"`


    // FilterDetails - Additional details to the source's filters.
    FilterDetails V3sourcefilterdetails `json:"filterDetails"`


    // VarError - Optional error details of an errored source.
    VarError Errorbody `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *V3sourcedetailedwitherrorresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcedetailedwitherrorresponse) MarshalJSON() ([]byte, error) {
    type Alias V3sourcedetailedwitherrorresponse

    if V3sourcedetailedwitherrorresponseMarshalled {
        return []byte("{}"), nil
    }
    V3sourcedetailedwitherrorresponseMarshalled = true

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
        
        ScheduleSettings V3sourceschedulesettings `json:"scheduleSettings"`
        
        Filters V3sourcefilter `json:"filters"`
        
        FilterDetails V3sourcefilterdetails `json:"filterDetails"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

