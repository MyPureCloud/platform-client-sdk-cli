package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchconversationeventrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchconversationeventrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Batchconversationeventrequest - A maximum of 100 events are allowed per request
type Batchconversationeventrequest struct { 
    // EndTransferEvents - Voice - EndTransfer events for this batch
    EndTransferEvents []Endtransferevent `json:"endTransferEvents"`


    // PhoneTransferEvents - Voice - PhoneTransfer events for this batch
    PhoneTransferEvents []Phonetransferevent `json:"phoneTransferEvents"`


    // ProgressTransferEvents - Voice - ProgressTransfer events for this batch
    ProgressTransferEvents []Progresstransferevent `json:"progressTransferEvents"`


    // RoutingTransferEvents - Voice - RoutingTransfer events for this batch
    RoutingTransferEvents []Routingtransferevent `json:"routingTransferEvents"`


    // UserTransferEvents - Voice - UserTransfer events for this batch
    UserTransferEvents []Usertransferevent `json:"userTransferEvents"`


    // CommunicationAnsweredEvents - Voice - CommunicationAnswered events for this batch
    CommunicationAnsweredEvents []Communicationansweredevent `json:"communicationAnsweredEvents"`


    // CommunicationDispositionAppliedEvents - Voice - CommunicationDispositionApplied events for this batch
    CommunicationDispositionAppliedEvents []Communicationdispositionappliedevent `json:"communicationDispositionAppliedEvents"`


    // HoldUpdatedEvents - Voice - HoldUpdated events for this batch
    HoldUpdatedEvents []Holdupdatedevent `json:"holdUpdatedEvents"`


    // ExternalEstablishedEvents - Voice - ExternalEstablished events for this batch
    ExternalEstablishedEvents []Externalestablishedevent `json:"externalEstablishedEvents"`


    // IvrEstablishedEvents - Voice - IvrEstablished events for this batch
    IvrEstablishedEvents []Ivrestablishedevent `json:"ivrEstablishedEvents"`


    // PhoneEstablishedEvents - Voice - PhoneEstablished events for this batch
    PhoneEstablishedEvents []Phoneestablishedevent `json:"phoneEstablishedEvents"`


    // RoutingEstablishedEvents - Voice - RoutingEstablished events for this batch
    RoutingEstablishedEvents []Routingestablishedevent `json:"routingEstablishedEvents"`


    // UserEstablishedEvents - Voice - UserEstablished events for this batch
    UserEstablishedEvents []Userestablishedevent `json:"userEstablishedEvents"`


    // AudioUpdatedEvents - Voice - AudioUpdated events for this batch
    AudioUpdatedEvents []Audioupdatedevent `json:"audioUpdatedEvents"`


    // CommunicationEndedEvents - Voice - CommunicationEnded events for this batch
    CommunicationEndedEvents []Communicationendedevent `json:"communicationEndedEvents"`


    // ConsultTransferEvents - Voice - ConsultTransfer events for this batch
    ConsultTransferEvents []Consulttransferevent `json:"consultTransferEvents"`


    // ProgressConsultTransferEvents - Voice - ProgressConsultTransfer events for this batch
    ProgressConsultTransferEvents []Progressconsulttransferevent `json:"progressConsultTransferEvents"`


    // EndConsultTransferEvents - Voice - EndConsultTransfer events for this batch
    EndConsultTransferEvents []Endconsulttransferevent `json:"endConsultTransferEvents"`


    // EmailBeginTransmittingEvents - Email - EmailBeginTransmittingEvent events for this batch
    EmailBeginTransmittingEvents []Emailbegintransmittingevent `json:"emailBeginTransmittingEvents"`


    // EmailCommunicationEndedEvents - Email - EmailCommunicationEndedEvent events for this batch
    EmailCommunicationEndedEvents []Emailcommunicationendedevent `json:"emailCommunicationEndedEvents"`


    // EmailExternalEstablishedEvents - Email - EmailExternalEstablishedEvent events for this batch
    EmailExternalEstablishedEvents []Emailexternalestablishedevent `json:"emailExternalEstablishedEvents"`


    // EmailFlowEstablishedEvents - Email - EmailFlowEstablishedEvent events for this batch
    EmailFlowEstablishedEvents []Emailflowestablishedevent `json:"emailFlowEstablishedEvents"`


    // EmailRoutingEstablishedEvents - Email - EmailRoutingEstablishedEvent events for this batch
    EmailRoutingEstablishedEvents []Emailroutingestablishedevent `json:"emailRoutingEstablishedEvents"`


    // EmailUserEstablishedEvents - Email - EmailUserEstablishedEvent events for this batch
    EmailUserEstablishedEvents []Emailuserestablishedevent `json:"emailUserEstablishedEvents"`


    // EmailCommunicationAnsweredEvents - Email - EmailCommunicationAnsweredEvent events for this batch
    EmailCommunicationAnsweredEvents []Emailcommunicationansweredevent `json:"emailCommunicationAnsweredEvents"`


    // EmailCommunicationDispositionAppliedEvents - Email - EmailCommunicationDispositionAppliedEvent events for this batch
    EmailCommunicationDispositionAppliedEvents []Emailcommunicationdispositionappliedevent `json:"emailCommunicationDispositionAppliedEvents"`


    // EmailCommunicationSentMessageEvents - Email - EmailCommunicationSentMessageEvent events for this batch
    EmailCommunicationSentMessageEvents []Emailcommunicationsentmessageevent `json:"emailCommunicationSentMessageEvents"`


    // EmailHoldUpdatedEvents - Email - EmailHoldUpdatedEvent events for this batch
    EmailHoldUpdatedEvents []Emailholdupdatedevent `json:"emailHoldUpdatedEvents"`


    // EmailEndTransferEvents - Email - EmailEndTransferEvent events for this batch
    EmailEndTransferEvents []Emailendtransferevent `json:"emailEndTransferEvents"`


    // EmailProgressTransferEvents - Email - EmailProgressTransferEvent events for this batch
    EmailProgressTransferEvents []Emailprogresstransferevent `json:"emailProgressTransferEvents"`


    // EmailRoutingTransferEvents - Email - EmailRoutingTransferEvent events for this batch
    EmailRoutingTransferEvents []Emailroutingtransferevent `json:"emailRoutingTransferEvents"`


    // EmailUserTransferEvents - Email - EmailUserTransferEvent events for this batch
    EmailUserTransferEvents []Emailusertransferevent `json:"emailUserTransferEvents"`

}

// String returns a JSON representation of the model
func (o *Batchconversationeventrequest) String() string {
     o.EndTransferEvents = []Endtransferevent{{}} 
     o.PhoneTransferEvents = []Phonetransferevent{{}} 
     o.ProgressTransferEvents = []Progresstransferevent{{}} 
     o.RoutingTransferEvents = []Routingtransferevent{{}} 
     o.UserTransferEvents = []Usertransferevent{{}} 
     o.CommunicationAnsweredEvents = []Communicationansweredevent{{}} 
     o.CommunicationDispositionAppliedEvents = []Communicationdispositionappliedevent{{}} 
     o.HoldUpdatedEvents = []Holdupdatedevent{{}} 
     o.ExternalEstablishedEvents = []Externalestablishedevent{{}} 
     o.IvrEstablishedEvents = []Ivrestablishedevent{{}} 
     o.PhoneEstablishedEvents = []Phoneestablishedevent{{}} 
     o.RoutingEstablishedEvents = []Routingestablishedevent{{}} 
     o.UserEstablishedEvents = []Userestablishedevent{{}} 
     o.AudioUpdatedEvents = []Audioupdatedevent{{}} 
     o.CommunicationEndedEvents = []Communicationendedevent{{}} 
     o.ConsultTransferEvents = []Consulttransferevent{{}} 
     o.ProgressConsultTransferEvents = []Progressconsulttransferevent{{}} 
     o.EndConsultTransferEvents = []Endconsulttransferevent{{}} 
     o.EmailBeginTransmittingEvents = []Emailbegintransmittingevent{{}} 
     o.EmailCommunicationEndedEvents = []Emailcommunicationendedevent{{}} 
     o.EmailExternalEstablishedEvents = []Emailexternalestablishedevent{{}} 
     o.EmailFlowEstablishedEvents = []Emailflowestablishedevent{{}} 
     o.EmailRoutingEstablishedEvents = []Emailroutingestablishedevent{{}} 
     o.EmailUserEstablishedEvents = []Emailuserestablishedevent{{}} 
     o.EmailCommunicationAnsweredEvents = []Emailcommunicationansweredevent{{}} 
     o.EmailCommunicationDispositionAppliedEvents = []Emailcommunicationdispositionappliedevent{{}} 
     o.EmailCommunicationSentMessageEvents = []Emailcommunicationsentmessageevent{{}} 
     o.EmailHoldUpdatedEvents = []Emailholdupdatedevent{{}} 
     o.EmailEndTransferEvents = []Emailendtransferevent{{}} 
     o.EmailProgressTransferEvents = []Emailprogresstransferevent{{}} 
     o.EmailRoutingTransferEvents = []Emailroutingtransferevent{{}} 
     o.EmailUserTransferEvents = []Emailusertransferevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchconversationeventrequest) MarshalJSON() ([]byte, error) {
    type Alias Batchconversationeventrequest

    if BatchconversationeventrequestMarshalled {
        return []byte("{}"), nil
    }
    BatchconversationeventrequestMarshalled = true

    return json.Marshal(&struct {
        
        EndTransferEvents []Endtransferevent `json:"endTransferEvents"`
        
        PhoneTransferEvents []Phonetransferevent `json:"phoneTransferEvents"`
        
        ProgressTransferEvents []Progresstransferevent `json:"progressTransferEvents"`
        
        RoutingTransferEvents []Routingtransferevent `json:"routingTransferEvents"`
        
        UserTransferEvents []Usertransferevent `json:"userTransferEvents"`
        
        CommunicationAnsweredEvents []Communicationansweredevent `json:"communicationAnsweredEvents"`
        
        CommunicationDispositionAppliedEvents []Communicationdispositionappliedevent `json:"communicationDispositionAppliedEvents"`
        
        HoldUpdatedEvents []Holdupdatedevent `json:"holdUpdatedEvents"`
        
        ExternalEstablishedEvents []Externalestablishedevent `json:"externalEstablishedEvents"`
        
        IvrEstablishedEvents []Ivrestablishedevent `json:"ivrEstablishedEvents"`
        
        PhoneEstablishedEvents []Phoneestablishedevent `json:"phoneEstablishedEvents"`
        
        RoutingEstablishedEvents []Routingestablishedevent `json:"routingEstablishedEvents"`
        
        UserEstablishedEvents []Userestablishedevent `json:"userEstablishedEvents"`
        
        AudioUpdatedEvents []Audioupdatedevent `json:"audioUpdatedEvents"`
        
        CommunicationEndedEvents []Communicationendedevent `json:"communicationEndedEvents"`
        
        ConsultTransferEvents []Consulttransferevent `json:"consultTransferEvents"`
        
        ProgressConsultTransferEvents []Progressconsulttransferevent `json:"progressConsultTransferEvents"`
        
        EndConsultTransferEvents []Endconsulttransferevent `json:"endConsultTransferEvents"`
        
        EmailBeginTransmittingEvents []Emailbegintransmittingevent `json:"emailBeginTransmittingEvents"`
        
        EmailCommunicationEndedEvents []Emailcommunicationendedevent `json:"emailCommunicationEndedEvents"`
        
        EmailExternalEstablishedEvents []Emailexternalestablishedevent `json:"emailExternalEstablishedEvents"`
        
        EmailFlowEstablishedEvents []Emailflowestablishedevent `json:"emailFlowEstablishedEvents"`
        
        EmailRoutingEstablishedEvents []Emailroutingestablishedevent `json:"emailRoutingEstablishedEvents"`
        
        EmailUserEstablishedEvents []Emailuserestablishedevent `json:"emailUserEstablishedEvents"`
        
        EmailCommunicationAnsweredEvents []Emailcommunicationansweredevent `json:"emailCommunicationAnsweredEvents"`
        
        EmailCommunicationDispositionAppliedEvents []Emailcommunicationdispositionappliedevent `json:"emailCommunicationDispositionAppliedEvents"`
        
        EmailCommunicationSentMessageEvents []Emailcommunicationsentmessageevent `json:"emailCommunicationSentMessageEvents"`
        
        EmailHoldUpdatedEvents []Emailholdupdatedevent `json:"emailHoldUpdatedEvents"`
        
        EmailEndTransferEvents []Emailendtransferevent `json:"emailEndTransferEvents"`
        
        EmailProgressTransferEvents []Emailprogresstransferevent `json:"emailProgressTransferEvents"`
        
        EmailRoutingTransferEvents []Emailroutingtransferevent `json:"emailRoutingTransferEvents"`
        
        EmailUserTransferEvents []Emailusertransferevent `json:"emailUserTransferEvents"`
        *Alias
    }{

        
        EndTransferEvents: []Endtransferevent{{}},
        


        
        PhoneTransferEvents: []Phonetransferevent{{}},
        


        
        ProgressTransferEvents: []Progresstransferevent{{}},
        


        
        RoutingTransferEvents: []Routingtransferevent{{}},
        


        
        UserTransferEvents: []Usertransferevent{{}},
        


        
        CommunicationAnsweredEvents: []Communicationansweredevent{{}},
        


        
        CommunicationDispositionAppliedEvents: []Communicationdispositionappliedevent{{}},
        


        
        HoldUpdatedEvents: []Holdupdatedevent{{}},
        


        
        ExternalEstablishedEvents: []Externalestablishedevent{{}},
        


        
        IvrEstablishedEvents: []Ivrestablishedevent{{}},
        


        
        PhoneEstablishedEvents: []Phoneestablishedevent{{}},
        


        
        RoutingEstablishedEvents: []Routingestablishedevent{{}},
        


        
        UserEstablishedEvents: []Userestablishedevent{{}},
        


        
        AudioUpdatedEvents: []Audioupdatedevent{{}},
        


        
        CommunicationEndedEvents: []Communicationendedevent{{}},
        


        
        ConsultTransferEvents: []Consulttransferevent{{}},
        


        
        ProgressConsultTransferEvents: []Progressconsulttransferevent{{}},
        


        
        EndConsultTransferEvents: []Endconsulttransferevent{{}},
        


        
        EmailBeginTransmittingEvents: []Emailbegintransmittingevent{{}},
        


        
        EmailCommunicationEndedEvents: []Emailcommunicationendedevent{{}},
        


        
        EmailExternalEstablishedEvents: []Emailexternalestablishedevent{{}},
        


        
        EmailFlowEstablishedEvents: []Emailflowestablishedevent{{}},
        


        
        EmailRoutingEstablishedEvents: []Emailroutingestablishedevent{{}},
        


        
        EmailUserEstablishedEvents: []Emailuserestablishedevent{{}},
        


        
        EmailCommunicationAnsweredEvents: []Emailcommunicationansweredevent{{}},
        


        
        EmailCommunicationDispositionAppliedEvents: []Emailcommunicationdispositionappliedevent{{}},
        


        
        EmailCommunicationSentMessageEvents: []Emailcommunicationsentmessageevent{{}},
        


        
        EmailHoldUpdatedEvents: []Emailholdupdatedevent{{}},
        


        
        EmailEndTransferEvents: []Emailendtransferevent{{}},
        


        
        EmailProgressTransferEvents: []Emailprogresstransferevent{{}},
        


        
        EmailRoutingTransferEvents: []Emailroutingtransferevent{{}},
        


        
        EmailUserTransferEvents: []Emailusertransferevent{{}},
        

        Alias: (*Alias)(u),
    })
}

