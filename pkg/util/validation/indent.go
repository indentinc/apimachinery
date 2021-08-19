package validation

const (
	labelNameFmt           = labelValueFmt
	labelNameErrMsg string = "must consist of alphanumeric characters and '-', '_', '.', '/', or '@'"
)

var labelNameRegexp = labelValueRegexp

// IsIndentLabelName tests whether the name passed is an Indent-acceptable
// label name (key). If the value is not valid, a list of error strings is
// returned. Otherwise an empty list (or nil) is returned.
//
// See IsQualifiedName; this is an Indent labels-specific drop-in replacement
// for that function. IsQualifiedName is used for apimachinery purposes besides
// labels, so it should not be modified.
func IsIndentLabelName(name string) []string {
	var errs []string
	if len(name) == 0 {
		errs = append(errs, "name "+EmptyError())
	} else if len(name) > qualifiedNameMaxLength {
		errs = append(errs, "name "+MaxLenError(qualifiedNameMaxLength))
	}
	if !labelNameRegexp.MatchString(name) {
		errs = append(errs, "name "+RegexError(labelNameErrMsg, labelNameFmt, "MyName", "my.name", "123-abc"))
	}
	return errs
}
