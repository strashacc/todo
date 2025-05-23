package repo

var (
	UnknownInsertError unknownInsertError
	BadUserError badUserError
	DBWriteError dbWriteError
	NoDocumentError noDocumentError
)

type unknownInsertError struct {
	error
}

func (err *unknownInsertError) Error() string {
	return "Team creation failed: Unknown error"
}

type badUserError struct {
	error
}

func (err *badUserError) Error() string {
	return "User does not exist"
}

type dbWriteError struct {
	error
}

func (err *dbWriteError) Error() string {
	return "Couldn't write to DB"
}

type noDocumentError struct {
	error
}

func (err *noDocumentError) Error() string {
	return "Document not found"
}
