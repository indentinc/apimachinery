package validation

// IsIndentLabelName tests whether the value passed is an Indent-acceptable
// label name (key). If the value is not valid, a list of error strings is
// returned. Otherwise an empty list (or nil) is returned.
//
// See IsQualifiedName; this is an Indent labels-specific drop-in replacement
// for that function. IsQualifiedName is used for apimachinery purposes besides
// labels, so it should not be modified.
func IsIndentLabelName(value string) []string {
	var errs []string
	if len(value) == 0 {
		errs = append(errs, "name "+EmptyError())
	} else if len(value) > qualifiedNameMaxLength {
		errs = append(errs, "name "+MaxLenError(qualifiedNameMaxLength))
	}
	return errs
}
