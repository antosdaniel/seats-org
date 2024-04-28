package organize

type Preference int

const (
	FrontSeatPreference Preference = iota
	RearSeatPreference
	WindowSeatPreference
	AisleSeatPreference
)

type Preferences []Preference

func (ps Preferences) Has(preference Preference) bool {
	for _, i := range ps {
		if i == preference {
			return true
		}
	}

	return false
}
