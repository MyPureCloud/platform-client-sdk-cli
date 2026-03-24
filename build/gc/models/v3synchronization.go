package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// V3synchronization
type V3synchronization struct { 
    


    // VarType - The type of the synchronization.
    VarType string `json:"type"`


    // CreatedBy - The user who started the synchronization if the source is manually synchronized or the user who created the source for scheduled synchronization.
    CreatedBy Userreference `json:"createdBy"`


    // Source - The source of the synchronization.
    Source V3sourceref `json:"source"`


    // DateStart - The start time of the synchronization. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The end time of the synchronization. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // DateSourceIntervalStart - The start time of the interval to be synchronized from the source. Source item changes during that interval are included in this synchronization. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateSourceIntervalStart time.Time `json:"dateSourceIntervalStart"`


    // DateSourceIntervalEnd - The end time of the interval to be synchronized from the source. Source item changes during that interval are included in this synchronization. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateSourceIntervalEnd time.Time `json:"dateSourceIntervalEnd"`


    // TriggerType - The trigger type of the synchronization.
    TriggerType string `json:"triggerType"`


    // Status - The status of the synchronization.
    Status string `json:"status"`


    // Statistics - Statistics of the synchronization.
    Statistics V3synchronizationstatistics `json:"statistics"`


    // VarError - The error that occurred during the synchronization.
    VarError Errorbody `json:"error"`


    // IngestionStatus - The status of the ingestion.
    IngestionStatus string `json:"ingestionStatus"`


    

}

// String returns a JSON representation of the model
func (o *V3synchronization) String() string {
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronization) MarshalJSON() ([]byte, error) {
    type Alias V3synchronization

    if V3synchronizationMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        Source V3sourceref `json:"source"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        DateSourceIntervalStart time.Time `json:"dateSourceIntervalStart"`
        
        DateSourceIntervalEnd time.Time `json:"dateSourceIntervalEnd"`
        
        TriggerType string `json:"triggerType"`
        
        Status string `json:"status"`
        
        Statistics V3synchronizationstatistics `json:"statistics"`
        
        VarError Errorbody `json:"error"`
        
        IngestionStatus string `json:"ingestionStatus"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

