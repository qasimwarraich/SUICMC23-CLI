package participants

type Participants struct {
	Items []struct {
		AdditionalComments string `json:"additional_comments"`
		Age                string `json:"age"`
		Bike               string `json:"bike"`
		CargoRace          bool   `json:"cargo_race"`
		Category           string `json:"category"`
		City               string `json:"city"`
		CollectionID       string `json:"collectionId"`
		CollectionName     string `json:"collectionName"`
		Created            string `json:"created"`
		Email              string `json:"email"`
		FirstName          string `json:"first_name"`
		Gears              string `json:"gears"`
		Heat               string `json:"heat"`
		Housing            string `json:"housing"`
		HousingFriday      bool   `json:"housing_friday"`
		HousingSaturday    bool   `json:"housing_saturday"`
		HousingSunday      bool   `json:"housing_sunday"`
		ID                 string `json:"id"`
		IntendedPayment    int    `json:"intended_payment"`
		InternalComments   string `json:"internal_comments"`
		LastName           string `json:"last_name"`
		Nabio              bool   `json:"nabio"`
		NickName           string `json:"nick_name"`
		Paid               bool   `json:"paid"`
		PaymentMethod      string `json:"payment_method"`
		PreEvent           bool   `json:"pre_event"`
		Pronouns           string `json:"pronouns"`
		RaceNumber         int    `json:"race_number"`
		RankSelection      string `json:"rank_selection"`
		Team               string `json:"team"`
		TshirtSize         string `json:"tshirt_size"`
		Updated            string `json:"updated"`
		Volunteering       bool   `json:"volunteering"`
	} `json:"items"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}
