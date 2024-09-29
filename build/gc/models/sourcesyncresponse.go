package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SourcesyncresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SourcesyncresponseDud struct { 
    


    


    


    


    


    

}

// Sourcesyncresponse
type Sourcesyncresponse struct { 
    // State - Sync state.
    State string `json:"state"`


    // VarError - Sync error.
    VarError Errorbody `json:"error"`


    // DateCreated - Synchronization creation date-time for this source. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Synchronization last modification date-time for this source. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // KnowledgeBase - Knowledge base to which this synchronization belongs.
    KnowledgeBase Addressableentityref `json:"knowledgeBase"`


    // Source - Source to which this synchronization belongs.
    Source Addressableentityref `json:"source"`

}

// String returns a JSON representation of the model
func (o *Sourcesyncresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sourcesyncresponse) MarshalJSON() ([]byte, error) {
    type Alias Sourcesyncresponse

    if SourcesyncresponseMarshalled {
        return []byte("{}"), nil
    }
    SourcesyncresponseMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        VarError Errorbody `json:"error"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        KnowledgeBase Addressableentityref `json:"knowledgeBase"`
        
        Source Addressableentityref `json:"source"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

