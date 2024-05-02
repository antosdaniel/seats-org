package organize

type Passengers []Passenger

type Passenger struct {
	id          PassengerId
	signupOrder int

	travelsWith []PassengerId
	preferences Preferences

	Details PassengerDetails
}

func (ps Passengers) WithPreference(preference Preference) Passengers {
	result := Passengers{}
	for _, p := range ps {
		if p.Preferences().Has(preference) {
			result = append(result, p)
		}
	}

	return result
}

func NewPassenger(id PassengerId, signupOrder int, preferences ...Preference) Passenger {
	return Passenger{
		id:          id,
		signupOrder: signupOrder,
		travelsWith: nil,
		preferences: preferences,
	}
}

type PassengerId string

func (p Passenger) Id() PassengerId {
	return p.id
}

func (p Passenger) SignupOrder() int {
	return p.signupOrder
}

func (p Passenger) TravelsAlone() bool {
	return len(p.travelsWith) == 0
}

func (p Passenger) TravelsWith() []PassengerId {
	return p.travelsWith
}

func (p Passenger) Preferences() Preferences {
	return p.preferences
}

type PassengerDetails struct {
	FullName string
}
