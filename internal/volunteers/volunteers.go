package volunteers

type Volunteers struct {
	Items []struct {
		AdditionalComments         string `json:"additional_comments"`
		CollectionID               string `json:"collectionId"`
		CollectionName             string `json:"collectionName"`
		Created                    string `json:"created"`
		Email                      string `json:"email"`
		FirstName                  string `json:"first_name"`
		ID                         string `json:"id"`
		LastName                   string `json:"last_name"`
		Pronouns                   string `json:"pronouns"`
		Telephone                  string `json:"telephone"`
		Updated                    string `json:"updated"`
		VolunteerFriday            bool   `json:"volunteer_friday"`
		VolunteerMonday            bool   `json:"volunteer_monday"`
		VolunteerMondayAfternoon   bool   `json:"volunteer_monday_afternoon"`
		VolunteerMondayMorning     bool   `json:"volunteer_monday_morning"`
		VolunteerSaturday          bool   `json:"volunteer_saturday"`
		VolunteerSaturdayAfternoon bool   `json:"volunteer_saturday_afternoon"`
		VolunteerSaturdayEvening   bool   `json:"volunteer_saturday_evening"`
		VolunteerSaturdayMorning   bool   `json:"volunteer_saturday_morning"`
		VolunteerSunday            bool   `json:"volunteer_sunday"`
		VolunteerSundayAfternoon   bool   `json:"volunteer_sunday_afternoon"`
		VolunteerSundayEvening     bool   `json:"volunteer_sunday_evening"`
		VolunteerSundayMorning     bool   `json:"volunteer_sunday_morning"`
	} `json:"items"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}
