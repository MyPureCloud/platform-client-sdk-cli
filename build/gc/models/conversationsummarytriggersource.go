package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsummarytriggersourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsummarytriggersourceDud struct { 
    


    


    

}

// Conversationsummarytriggersource
type Conversationsummarytriggersource struct { 
    // SourceType - The configuration entity for which summarization is triggered.
    SourceType string `json:"sourceType"`


    // SourceId - The id value for the source type.
    SourceId string `json:"sourceId"`


    // SourceOutcome - The reason a trigger source finished processing. Only applies to Flow trigger source types.
    SourceOutcome string `json:"sourceOutcome"`

}

// String returns a JSON representation of the model
func (o *Conversationsummarytriggersource) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsummarytriggersource) MarshalJSON() ([]byte, error) {
    type Alias Conversationsummarytriggersource

    if ConversationsummarytriggersourceMarshalled {
        return []byte("{}"), nil
    }
    ConversationsummarytriggersourceMarshalled = true

    return json.Marshal(&struct {
        
        SourceType string `json:"sourceType"`
        
        SourceId string `json:"sourceId"`
        
        SourceOutcome string `json:"sourceOutcome"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

