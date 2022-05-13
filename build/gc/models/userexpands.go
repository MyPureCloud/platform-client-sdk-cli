package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserexpandsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserexpandsDud struct { 
    RoutingStatus Routingstatus `json:"routingStatus"`


    Presence Userpresence `json:"presence"`


    IntegrationPresence Userpresence `json:"integrationPresence"`


    ConversationSummary Userconversationsummary `json:"conversationSummary"`


    OutOfOffice Outofoffice `json:"outOfOffice"`


    Geolocation Geolocation `json:"geolocation"`


    Station Userstations `json:"station"`


    Authorization Userauthorization `json:"authorization"`

}

// Userexpands
type Userexpands struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Userexpands) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userexpands) MarshalJSON() ([]byte, error) {
    type Alias Userexpands

    if UserexpandsMarshalled {
        return []byte("{}"), nil
    }
    UserexpandsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

