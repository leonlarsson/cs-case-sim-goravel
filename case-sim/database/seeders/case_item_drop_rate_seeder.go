package seeders

type CaseItemDropRateSeeder struct {
}

type GradeOdds map[string]float64

var caseGradeOdds = GradeOdds{
	"Mil-Spec Grade":    0.79923,
	"Restricted":        0.15985,
	"Classified":        0.03197,
	"Covert":            0.00639,
	"Rare Special Item": 0.00256,
}

var souvenirGradeOdds = GradeOdds{
	"Consumer Grade":   0.79872,
	"Industrial Grade": 0.15974,
	"Mil-Spec Grade":   0.03328,
	"Restricted":       0.00666,
	"Classified":       0.00133,
	"Covert":           0.00027,
}

// Signature The name and signature of the seeder.
func (s *CaseItemDropRateSeeder) Signature() string {
	return "CaseItemDropRateSeeder"
}

// Run executes the seeder logic.
func (s *CaseItemDropRateSeeder) Run() error {
	return nil
}
