package event

type Services struct {
	CreateEvent EventModel
	EditEvent   EventModel
	DeleteEvent EventModel
	Landing     ViewLanding
	SearchEvent QueryEvent
	SingleEvent ViewFullEventModel

	CreateTimeSlot TimeSlot
	EditTimeSlot   TimeSlot
	DeleteTimeSlot TimeSlot
	FindTimeSlot   QueryTimeSlot

	CreateEventType EventsType
	EditEventType   EventsType
	DeleteEventType EventsType
	ListEventTypes  QueryEventsType

	CreateEventStatus EventsStatus
	EditEventStatus   EventsStatus
	DeleteEventStatus EventsStatus
	ListEventStatus   QueryEventsStatus

	AddEventScore       Score
	CalculateEventScore QueryScore
	ListHighScores      QueryScore

	AddEventCrew     Crew
	EditEventCrew    Crew
	DeleteEventCrew  Crew
	FindEventCrews   QueryCrew
	FindEventsByCrew QueryCrew
}
