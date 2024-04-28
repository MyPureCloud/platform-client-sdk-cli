package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatconversationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Chatconversation
type Chatconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Participants - The list of participants involved in the conversation.
    Participants []Chatmediaparticipant `json:"participants"`


    // OtherMediaUris - The list of other media channels involved in the conversation.
    OtherMediaUris []string `json:"otherMediaUris"`


    // RecentTransfers - The list of the most recent 20 transfer commands applied to this conversation.
    RecentTransfers []Transferresponse `json:"recentTransfers"`


    // UtilizationLabelId - An optional label that categorizes the conversation.  Max-utilization settings can be configured at a per-label level
    UtilizationLabelId string `json:"utilizationLabelId"`


    

}

// String returns a JSON representation of the model
func (o *Chatconversation) String() string {
    
     o.Participants = []Chatmediaparticipant{{}} 
     o.OtherMediaUris = []string{""} 
     o.RecentTransfers = []Transferresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatconversation) MarshalJSON() ([]byte, error) {
    type Alias Chatconversation

    if ChatconversationMarshalled {
        return []byte("{}"), nil
    }
    ChatconversationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Participants []Chatmediaparticipant `json:"participants"`
        
        OtherMediaUris []string `json:"otherMediaUris"`
        
        RecentTransfers []Transferresponse `json:"recentTransfers"`
        
        UtilizationLabelId string `json:"utilizationLabelId"`
        *Alias
    }{

        


        


        
        Participants: []Chatmediaparticipant{{}},
        


        
        OtherMediaUris: []string{""},
        


        
        RecentTransfers: []Transferresponse{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

