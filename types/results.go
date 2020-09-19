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
	WithSenseSpecificInflectionPluralLabel // XXX: might not belong here? docs unclear
	WithSenseSpecificGrammaticalLabel      // XXX: might not belong here? docs unclear

	WithInflections
	WithCognateCrossReferences

	WithDefinitions

	// misc
	WithDefinedRunOns
	WithUndefinedRunOns
	WithDirectionalCrossReferences
	WithUsages
	WithSynonyms
	WithQuotes
	WithTable

	WithEtymology
	WithFirstKnownDate
	WithArt
	WithVariants

	WithShortDefinitions
}
