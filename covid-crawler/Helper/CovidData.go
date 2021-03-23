package Helper

type CovidData struct {
	// Day is this data's day of occurrence
	Day int
	// Month is this data's day of occurrence
	Month int
	// Year is this data's day of occurrence
	Year int

	// InvestigatingCases is the number of possible cases that does not yet have the test's result.
	InvestigatingCases int
	// DiscardedCases is the number of cases that tested negative for covid-19.
	DiscardedCases int
	// ConfirmedCases is the number of cases that tested positive for covid-19.
	ConfirmedCases int
	//ConfirmedPatients is the number of patients that are currently with covid-19.
	ConfirmedPatients int
	// NotifiedCases is the number of cases notified by the town hall.
	NotifiedCases int
	// ExcludedCases is the number of cases that do not required test, due to being excluded by other reasons.
	ExcludedCases int

	// InvestigatingDeaths is the number of deaths that are currently waiting for the test's result
	InvestigatingDeaths int
	// ConfirmedDeath is the number of deaths that tested positive for covid-19.
	ConfirmedDeath int

	// FluSymptoms is the number of people that have flu symptoms but no other covid-19 symptoms
	FluSymptoms int
}