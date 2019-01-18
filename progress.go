package gedcom

// Progress contains information about the progress of an operation.
//
// Progress will consist of a value for Add or Done. It will also optionally
// contain a value for Total.
type Progress struct {
	// Done is how many operations have been completed so far. If Done is zero
	// you should add the value of Add instead.
	Done int64

	// Add represents how many operations were performed since the last
	// operation. It is possible for both the Done and Add to be zero. This
	// means the progress did not change. Add may also be a negative value.
	Add int64

	// Total is the expected number of total operations. Total may change
	// throughout the processing to be larger to smaller that any previous
	// value.
	//
	// If total is zero, you should not change the existing total value.
	Total int64
}
