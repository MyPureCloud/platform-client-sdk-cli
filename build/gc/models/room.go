package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoomMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoomDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Room
type Room struct { 
    


    // Name
    Name string `json:"name"`


    // DateCreated - Room's created time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // RoomType - The type of room
    RoomType string `json:"roomType"`


    // Description - Room's description
    Description string `json:"description"`


    // Subject - Room's subject
    Subject string `json:"subject"`


    // ParticipantLimit - Room's size limit
    ParticipantLimit int `json:"participantLimit"`


    // Owners - Room's owners
    Owners []Userreference `json:"owners"`


    // PinnedMessages - Room's pinned messages
    PinnedMessages []Addressableentityref `json:"pinnedMessages"`


    // Jid - The jid of the room
    Jid string `json:"jid"`


    

}

// String returns a JSON representation of the model
func (o *Room) String() string {
    
    
    
    
    
    
     o.Owners = []Userreference{{}} 
     o.PinnedMessages = []Addressableentityref{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Room) MarshalJSON() ([]byte, error) {
    type Alias Room

    if RoomMarshalled {
        return []byte("{}"), nil
    }
    RoomMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        RoomType string `json:"roomType"`
        
        Description string `json:"description"`
        
        Subject string `json:"subject"`
        
        ParticipantLimit int `json:"participantLimit"`
        
        Owners []Userreference `json:"owners"`
        
        PinnedMessages []Addressableentityref `json:"pinnedMessages"`
        
        Jid string `json:"jid"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Owners: []Userreference{{}},
        


        
        PinnedMessages: []Addressableentityref{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

