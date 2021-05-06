package watchman

// aggregatorValue stores entropy of project
var AggregatorValue int

// Aggregate is called by other routines to update the aggregatorValue
func Aggregate(entropy int) {
	AggregatorValue += entropy
}

// ResetAggregator sets the project entropy to zero, to be run after each cycle
func ResetAggregator(){
	AggregatorValue = 0
}
