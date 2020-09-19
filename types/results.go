package types

// CollegiateResult is a result from a search through the M-W Collegiate dictionary
type CollegiateResult struct {
	WithMetadata
	WithHomograph
	WithHeadwordInfo

	WithFunctionalLabel
	WithGeneralLabels
	WithSubjectStatusLabels
	WithParenthesizedSubjectStatusLabel
	WithSenseSpecificInflectionPluralLabel // might not belong here?
	WithSenseSpecificGrammaticalLabel      // might not belong here?

	WithInflections
	WithCognateCrossReferences

	WithDefinitions

	// misc
	WithDefinedRunOns
	withUndefinedRunOns
	WithDirectionalCrossReferences
	WithUsages
	WithSynonyms
	WithQuotes
	WithTable

	WithEtymology
	WithFirstKnownDate

	WithShortDefinitions
}
