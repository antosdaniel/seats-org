package organize

type Passengers []Passenger

type Passenger struct {
	id       PassengerId
	fullName string

	travelsWith []PassengerId
	preferences Preferences
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

func NewPassenger(id PassengerId, fullName string, preferences ...Preference) Passenger {
	return Passenger{
		id:          id,
		fullName:    fullName,
		travelsWith: nil,
		preferences: preferences,
	}
}

type PassengerId string

func (p Passenger) Id() PassengerId {
	return p.id
}

func (p Passenger) FullName() string {
	return p.fullName
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
