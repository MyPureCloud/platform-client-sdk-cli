package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    


    


    ExternalContact Addressableentityref `json:"externalContact"`


    


    


    Conversation Addressableentityref `json:"conversation"`

}

// Session
type Session struct { 
    


    // CustomerId - Primary identifier of the customer in the source where the events for the session originate from.
    CustomerId string `json:"customerId"`


    // CustomerIdType - Type of source customer identifier (e.g. cookie, email, phone).
    CustomerIdType string `json:"customerIdType"`


    // VarType - Session types indicate the type or category of sessions (e.g. web, ticket, delivery, atm).
    VarType string `json:"type"`


    // ExternalId - Unique identifier in the external system where the events for the session originate from.
    ExternalId string `json:"externalId"`


    // ExternalUrl - A URL that identifies an external system-of-record resource that may have more detailed information on the session.
    ExternalUrl string `json:"externalUrl"`


    // ShortId - Shortened numeric identifier of 4-6 digits.
    ShortId string `json:"shortId"`


    // OutcomeAchievements - List of the outcome achievements by the customer in this session.
    OutcomeAchievements []Outcomeachievement `json:"outcomeAchievements"`


    // SegmentAssignments - List of the segment assignments to the customer in this session.
    SegmentAssignments []Sessionsegmentassignment `json:"segmentAssignments"`


    // Attributes - Attributes projected from the session's event stream.
    Attributes map[string]Customeventattribute `json:"attributes"`


    // AttributeLists - List-type attributes projected from the session's event stream.
    AttributeLists map[string]Customeventattributelist `json:"attributeLists"`


    // Browser - Customer's browser.
    Browser Browser `json:"browser"`


    // Device - Customer's device.
    Device Device `json:"device"`


    // Geolocation - Customer's geolocation.
    Geolocation Journeygeolocation `json:"geolocation"`


    // IpAddress - Customer's IP address.
    IpAddress string `json:"ipAddress"`


    // IpOrganization - Customer's IP-based organization or ISP name.
    IpOrganization string `json:"ipOrganization"`


    // LastPage - The webpage where the customer's last web interaction occurred.
    LastPage Journeypage `json:"lastPage"`


    // MktCampaign - Marketing / traffic source information.
    MktCampaign Journeycampaign `json:"mktCampaign"`


    // Referrer - Identifies the page URL that originally generated the request for the current page being viewed.
    Referrer Referrer `json:"referrer"`


    // SearchTerms - Search terms associated with the session.
    SearchTerms []string `json:"searchTerms"`


    // UserAgentString - String identifying the user agent.
    UserAgentString string `json:"userAgentString"`


    // DurationInSeconds - Indicates how long the session has been active (valid for an individual device).
    DurationInSeconds int `json:"durationInSeconds"`


    // EventCount - The count of all events performed during the session.
    EventCount int `json:"eventCount"`


    // PageviewCount - The count of all pageviews performed during the session.
    PageviewCount int `json:"pageviewCount"`


    // ScreenviewCount - The count of all screenviews performed during the session.
    ScreenviewCount int `json:"screenviewCount"`


    // LastEvent - Information about the most recent event in this session.
    LastEvent Sessionlastevent `json:"lastEvent"`


    // LastConnectedQueue - The last queue connected to this session.
    LastConnectedQueue Connectedqueue `json:"lastConnectedQueue"`


    // LastConnectedUser - The last user connected to this session.
    LastConnectedUser Connecteduser `json:"lastConnectedUser"`


    // LastUserDisposition - The last user disposition connected to this session.
    LastUserDisposition Conversationuserdisposition `json:"lastUserDisposition"`


    // ConversationChannels - Represents the channels used for this conversation.
    ConversationChannels []Conversationchannel `json:"conversationChannels"`


    // OriginatingDirection - The original direction of the conversation.
    OriginatingDirection string `json:"originatingDirection"`


    // ConversationSubject - The subject for the conversation, for example an email subject.
    ConversationSubject string `json:"conversationSubject"`


    // Authenticated - Indicates whether or not the session is authenticated.
    Authenticated bool `json:"authenticated"`


    


    // CreatedDate - Timestamp indicating when the session was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // EndedDate - Timestamp indicating when the session was ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndedDate time.Time `json:"endedDate"`


    


    // AwayDate - Timestamp indicating when the visitor should be considered as away. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    AwayDate time.Time `json:"awayDate"`


    // IdleDate - Timestamp indicating when the visitor should be considered as idle. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    IdleDate time.Time `json:"idleDate"`


    

}

// String returns a JSON representation of the model
func (o *Session) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.OutcomeAchievements = []Outcomeachievement{{}} 
    
    
    
     o.SegmentAssignments = []Sessionsegmentassignment{{}} 
    
    
    
     o.Attributes = map[string]Customeventattribute{"": {}} 
    
    
    
     o.AttributeLists = map[string]Customeventattributelist{"": {}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.SearchTerms = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ConversationChannels = []Conversationchannel{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Session) MarshalJSON() ([]byte, error) {
    type Alias Session

    if SessionMarshalled {
        return []byte("{}"), nil
    }
    SessionMarshalled = true

    return json.Marshal(&struct { 
        
        
        CustomerId string `json:"customerId"`
        
        CustomerIdType string `json:"customerIdType"`
        
        VarType string `json:"type"`
        
        ExternalId string `json:"externalId"`
        
        ExternalUrl string `json:"externalUrl"`
        
        ShortId string `json:"shortId"`
        
        OutcomeAchievements []Outcomeachievement `json:"outcomeAchievements"`
        
        SegmentAssignments []Sessionsegmentassignment `json:"segmentAssignments"`
        
        Attributes map[string]Customeventattribute `json:"attributes"`
        
        AttributeLists map[string]Customeventattributelist `json:"attributeLists"`
        
        Browser Browser `json:"browser"`
        
        Device Device `json:"device"`
        
        Geolocation Journeygeolocation `json:"geolocation"`
        
        IpAddress string `json:"ipAddress"`
        
        IpOrganization string `json:"ipOrganization"`
        
        LastPage Journeypage `json:"lastPage"`
        
        MktCampaign Journeycampaign `json:"mktCampaign"`
        
        Referrer Referrer `json:"referrer"`
        
        SearchTerms []string `json:"searchTerms"`
        
        UserAgentString string `json:"userAgentString"`
        
        DurationInSeconds int `json:"durationInSeconds"`
        
        EventCount int `json:"eventCount"`
        
        PageviewCount int `json:"pageviewCount"`
        
        ScreenviewCount int `json:"screenviewCount"`
        
        LastEvent Sessionlastevent `json:"lastEvent"`
        
        LastConnectedQueue Connectedqueue `json:"lastConnectedQueue"`
        
        LastConnectedUser Connecteduser `json:"lastConnectedUser"`
        
        LastUserDisposition Conversationuserdisposition `json:"lastUserDisposition"`
        
        ConversationChannels []Conversationchannel `json:"conversationChannels"`
        
        OriginatingDirection string `json:"originatingDirection"`
        
        ConversationSubject string `json:"conversationSubject"`
        
        Authenticated bool `json:"authenticated"`
        
        
        
        CreatedDate time.Time `json:"createdDate"`
        
        EndedDate time.Time `json:"endedDate"`
        
        
        
        AwayDate time.Time `json:"awayDate"`
        
        IdleDate time.Time `json:"idleDate"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        OutcomeAchievements: []Outcomeachievement{{}},
        

        

        
        SegmentAssignments: []Sessionsegmentassignment{{}},
        

        

        
        Attributes: map[string]Customeventattribute{"": {}},
        

        

        
        AttributeLists: map[string]Customeventattributelist{"": {}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        SearchTerms: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        ConversationChannels: []Conversationchannel{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

