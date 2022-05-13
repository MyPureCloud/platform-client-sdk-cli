package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationpropertiesDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Conversationproperties
type Conversationproperties struct { 
    // IsWaiting - Indicates filtering for waiting
    IsWaiting bool `json:"isWaiting"`


    // IsActive - Indicates filtering for active
    IsActive bool `json:"isActive"`


    // IsAcd - Indicates filtering for Acd
    IsAcd bool `json:"isAcd"`


    // IsPreferred - Indicates filtering for Preferred Agent Routing
    IsPreferred bool `json:"isPreferred"`


    // IsScreenshare - Indicates filtering for screenshare
    IsScreenshare bool `json:"isScreenshare"`


    // IsCobrowse - Indicates filtering for Cobrowse
    IsCobrowse bool `json:"isCobrowse"`


    // IsVoicemail - Indicates filtering for Voice mail
    IsVoicemail bool `json:"isVoicemail"`


    // IsFlagged - Indicates filtering for flagged
    IsFlagged bool `json:"isFlagged"`


    // IsMonitored - Indicates filtering for monitored
    IsMonitored bool `json:"isMonitored"`


    // FilterWrapUpNotes - Indicates filtering for WrapUpNotes
    FilterWrapUpNotes bool `json:"filterWrapUpNotes"`


    // MatchAll - Indicates comparison operation, TRUE indicates filters will use AND logic, FALSE indicates OR logic
    MatchAll bool `json:"matchAll"`

}

// String returns a JSON representation of the model
func (o *Conversationproperties) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationproperties) MarshalJSON() ([]byte, error) {
    type Alias Conversationproperties

    if ConversationpropertiesMarshalled {
        return []byte("{}"), nil
    }
    ConversationpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        IsWaiting bool `json:"isWaiting"`
        
        IsActive bool `json:"isActive"`
        
        IsAcd bool `json:"isAcd"`
        
        IsPreferred bool `json:"isPreferred"`
        
        IsScreenshare bool `json:"isScreenshare"`
        
        IsCobrowse bool `json:"isCobrowse"`
        
        IsVoicemail bool `json:"isVoicemail"`
        
        IsFlagged bool `json:"isFlagged"`
        
        IsMonitored bool `json:"isMonitored"`
        
        FilterWrapUpNotes bool `json:"filterWrapUpNotes"`
        
        MatchAll bool `json:"matchAll"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

