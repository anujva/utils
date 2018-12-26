package utils

// StrPtr returns a pointer to the string
func StrPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to an int
func IntPtr(i int) *int {
	return &i
}

// Int64Ptr returns a pointer to int64
func Int64Ptr(i64 int64) *int64 {
	return &i64
}
