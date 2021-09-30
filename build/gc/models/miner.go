package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinerDud struct { 
    Id string `json:"id"`


    


    


    DateCreated time.Time `json:"dateCreated"`


    Status string `json:"status"`


    ConversationsDateRangeStart time.Time `json:"conversationsDateRangeStart"`


    ConversationsDateRangeEnd time.Time `json:"conversationsDateRangeEnd"`


    DateCompleted time.Time `json:"dateCompleted"`


    Message string `json:"message"`


    ConversationDataUploaded bool `json:"conversationDataUploaded"`


    MediaType string `json:"mediaType"`


    QueueIds []string `json:"queueIds"`


    DateTriggered time.Time `json:"dateTriggered"`


    DateModified time.Time `json:"dateModified"`


    LatestDraftVersion *Draft `json:"latestDraftVersion"`


    SelfUri string `json:"selfUri"`

}

// Miner
type Miner struct { 
    


    // Name - Chat Corpus Name.
    Name string `json:"name"`


    // Language - Language Localization code.
    Language string `json:"language"`


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Miner) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Miner) MarshalJSON() ([]byte, error) {
    type Alias Miner

    if MinerMarshalled {
        return []byte("{}"), nil
    }
    MinerMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Language string `json:"language"`
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

