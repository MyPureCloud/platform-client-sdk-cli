package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HomerrecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HomerrecordDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Homerrecord
type Homerrecord struct { 
    


    // Name
    Name string `json:"name"`


    // Date - metadata associated to the SIP calls. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Date time.Time `json:"date"`


    // MilliTs - metadata associated to the SIP calls
    MilliTs string `json:"milliTs"`


    // MicroTs - metadata associated to the SIP calls
    MicroTs string `json:"microTs"`


    // Method - metadata associated to the SIP calls
    Method string `json:"method"`


    // ReplyReason - metadata associated to the SIP calls
    ReplyReason string `json:"replyReason"`


    // Ruri - metadata associated to the SIP calls
    Ruri string `json:"ruri"`


    // RuriUser - metadata associated to the SIP calls
    RuriUser string `json:"ruriUser"`


    // RuriDomain - metadata associated to the SIP calls
    RuriDomain string `json:"ruriDomain"`


    // FromUser - metadata associated to the SIP calls
    FromUser string `json:"fromUser"`


    // FromDomain - metadata associated to the SIP calls
    FromDomain string `json:"fromDomain"`


    // FromTag - metadata associated to the SIP calls
    FromTag string `json:"fromTag"`


    // ToUser - metadata associated to the SIP calls
    ToUser string `json:"toUser"`


    // ToDomain - metadata associated to the SIP calls
    ToDomain string `json:"toDomain"`


    // ToTag - metadata associated to the SIP calls
    ToTag string `json:"toTag"`


    // PidUser - metadata associated to the SIP calls
    PidUser string `json:"pidUser"`


    // ContactUser - metadata associated to the SIP calls
    ContactUser string `json:"contactUser"`


    // AuthUser - metadata associated to the SIP calls
    AuthUser string `json:"authUser"`


    // Callid - metadata associated to the SIP calls
    Callid string `json:"callid"`


    // CallidAleg - metadata associated to the SIP calls
    CallidAleg string `json:"callidAleg"`


    // Via1 - metadata associated to the SIP calls
    Via1 string `json:"via1"`


    // Via1Branch - metadata associated to the SIP calls
    Via1Branch string `json:"via1Branch"`


    // Cseq - metadata associated to the SIP calls
    Cseq string `json:"cseq"`


    // Diversion - metadata associated to the SIP calls
    Diversion string `json:"diversion"`


    // Reason - metadata associated to the SIP calls
    Reason string `json:"reason"`


    // ContentType - metadata associated to the SIP calls
    ContentType string `json:"contentType"`


    // Auth - metadata associated to the SIP calls
    Auth string `json:"auth"`


    // UserAgent - metadata associated to the SIP calls
    UserAgent string `json:"userAgent"`


    // SourceIp - metadata associated to the SIP calls
    SourceIp string `json:"sourceIp"`


    // SourcePort - metadata associated to the SIP calls
    SourcePort string `json:"sourcePort"`


    // DestinationIp - metadata associated to the SIP calls
    DestinationIp string `json:"destinationIp"`


    // DestinationPort - metadata associated to the SIP calls
    DestinationPort string `json:"destinationPort"`


    // ContactIp - metadata associated to the SIP calls
    ContactIp string `json:"contactIp"`


    // ContactPort - metadata associated to the SIP calls
    ContactPort string `json:"contactPort"`


    // OriginatorIp - metadata associated to the SIP calls
    OriginatorIp string `json:"originatorIp"`


    // OriginatorPort - metadata associated to the SIP calls
    OriginatorPort string `json:"originatorPort"`


    // CorrelationId - metadata associated to the SIP calls
    CorrelationId string `json:"correlationId"`


    // Proto - metadata associated to the SIP calls
    Proto string `json:"proto"`


    // Family - metadata associated to the SIP calls
    Family string `json:"family"`


    // RtpStat - metadata associated to the SIP calls
    RtpStat string `json:"rtpStat"`


    // VarType - metadata associated to the SIP calls
    VarType string `json:"type"`


    // Node - metadata associated to the SIP calls
    Node string `json:"node"`


    // Trans - metadata associated to the SIP calls
    Trans string `json:"trans"`


    // Dbnode - metadata associated to the SIP calls
    Dbnode string `json:"dbnode"`


    // Msg - metadata associated to the SIP calls
    Msg string `json:"msg"`


    // SourceAlias - metadata associated to the SIP calls
    SourceAlias string `json:"sourceAlias"`


    // DestinationAlias - metadata associated to the SIP calls
    DestinationAlias string `json:"destinationAlias"`


    // ConversationId - metadata associated to the SIP calls
    ConversationId string `json:"conversationId"`


    // ParticipantId - metadata associated to the SIP calls
    ParticipantId string `json:"participantId"`


    

}

// String returns a JSON representation of the model
func (o *Homerrecord) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Homerrecord) MarshalJSON() ([]byte, error) {
    type Alias Homerrecord

    if HomerrecordMarshalled {
        return []byte("{}"), nil
    }
    HomerrecordMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Date time.Time `json:"date"`
        
        MilliTs string `json:"milliTs"`
        
        MicroTs string `json:"microTs"`
        
        Method string `json:"method"`
        
        ReplyReason string `json:"replyReason"`
        
        Ruri string `json:"ruri"`
        
        RuriUser string `json:"ruriUser"`
        
        RuriDomain string `json:"ruriDomain"`
        
        FromUser string `json:"fromUser"`
        
        FromDomain string `json:"fromDomain"`
        
        FromTag string `json:"fromTag"`
        
        ToUser string `json:"toUser"`
        
        ToDomain string `json:"toDomain"`
        
        ToTag string `json:"toTag"`
        
        PidUser string `json:"pidUser"`
        
        ContactUser string `json:"contactUser"`
        
        AuthUser string `json:"authUser"`
        
        Callid string `json:"callid"`
        
        CallidAleg string `json:"callidAleg"`
        
        Via1 string `json:"via1"`
        
        Via1Branch string `json:"via1Branch"`
        
        Cseq string `json:"cseq"`
        
        Diversion string `json:"diversion"`
        
        Reason string `json:"reason"`
        
        ContentType string `json:"contentType"`
        
        Auth string `json:"auth"`
        
        UserAgent string `json:"userAgent"`
        
        SourceIp string `json:"sourceIp"`
        
        SourcePort string `json:"sourcePort"`
        
        DestinationIp string `json:"destinationIp"`
        
        DestinationPort string `json:"destinationPort"`
        
        ContactIp string `json:"contactIp"`
        
        ContactPort string `json:"contactPort"`
        
        OriginatorIp string `json:"originatorIp"`
        
        OriginatorPort string `json:"originatorPort"`
        
        CorrelationId string `json:"correlationId"`
        
        Proto string `json:"proto"`
        
        Family string `json:"family"`
        
        RtpStat string `json:"rtpStat"`
        
        VarType string `json:"type"`
        
        Node string `json:"node"`
        
        Trans string `json:"trans"`
        
        Dbnode string `json:"dbnode"`
        
        Msg string `json:"msg"`
        
        SourceAlias string `json:"sourceAlias"`
        
        DestinationAlias string `json:"destinationAlias"`
        
        ConversationId string `json:"conversationId"`
        
        ParticipantId string `json:"participantId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

