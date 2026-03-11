package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesourcessearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesourcessearchrequestDud struct { 
    


    


    


    


    


    


    


    

}

// Knowledgesourcessearchrequest
type Knowledgesourcessearchrequest struct { 
    // Query - Input query to search content on the knowledge setting.
    Query string `json:"query"`


    // KnowledgeSettingId - Knowledge Setting Id to use for search request.
    KnowledgeSettingId string `json:"knowledgeSettingId"`


    // Application - The client application details from which search requested.
    Application V3knowledgesearchclientapplication `json:"application"`


    // ConversationContext - Conversation context information if the search is initiated in the context of a conversation.
    ConversationContext Knowledgev3conversationcontext `json:"conversationContext"`


    // SessionId - The session id for search request.
    SessionId string `json:"sessionId"`


    // QueryType - The type of the query that initiates the search.
    QueryType string `json:"queryType"`


    // GenerationLanguage - The language to use for answer generation.
    GenerationLanguage string `json:"generationLanguage"`


    // ConversationTurns - List of conversation turns to use for stateful search.
    ConversationTurns []Knowledgeconversationturn `json:"conversationTurns"`

}

// String returns a JSON representation of the model
func (o *Knowledgesourcessearchrequest) String() string {
    
    
    
    
    
    
    
     o.ConversationTurns = []Knowledgeconversationturn{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesourcessearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesourcessearchrequest

    if KnowledgesourcessearchrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesourcessearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        KnowledgeSettingId string `json:"knowledgeSettingId"`
        
        Application V3knowledgesearchclientapplication `json:"application"`
        
        ConversationContext Knowledgev3conversationcontext `json:"conversationContext"`
        
        SessionId string `json:"sessionId"`
        
        QueryType string `json:"queryType"`
        
        GenerationLanguage string `json:"generationLanguage"`
        
        ConversationTurns []Knowledgeconversationturn `json:"conversationTurns"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        ConversationTurns: []Knowledgeconversationturn{{}},
        

        Alias: (*Alias)(u),
    })
}

