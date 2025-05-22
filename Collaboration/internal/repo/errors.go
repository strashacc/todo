package repo

type UnknownInsertError struct {
	error
}
func (err *UnknownInsertError) Error() string {
	return "Team creation failed: Unknown error"
}

type BadUserError struct {
	error
}
func (err *BadUserError) Error() string {
	return "User does not exist"
}

type DBWriteError struct {
	error
}
func (err DBWriteError) Error() string {
	return "Couldn't write to DB"
}

type NoDocumentError struct {
	error
}
func (err *NoDocumentError) Error() string {
	return "Document not found"
}