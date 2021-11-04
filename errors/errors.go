package errors

// type ErrorInfo struct {
// 	// HttpStatus int       `json:"httpStatus"`
// 	ErrorBody ErrorBody `json:"error"`
// }

/**
 * This struct creates a error info from the two parts: Code and Msg,
 * , which is uint16 and String, respectively.
 *
 * @param Id               the unique Id of short url as a string
 * @param Alias            the unique Alias of short url as a String
 * @param LongUrl          the original Url as a String
 * @param ShortUrl         the short url as a String
 */
type ErrorInfo struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
}

func (e ErrorInfo) IsNil() bool {
	return ErrorInfo{} == e
}

var NoError = ErrorInfo{}

var InvalidInputError = ErrorInfo{
	Code: 40001,
	Msg:  "Invalid input",
}

// Invalid long Url error means the long url is empty
var InvalidTableNumberError = ErrorInfo{
	Code: 40002,
	Msg:  "Invalid table number",
}

// Invalid alias error means the alias is not letter,number or length more than 30 lengths
var InsufficientSpaceError = ErrorInfo{
	Code: 40003,
	Msg:  "Insufficient space",
}

// Alias forbiden error means the alias is used
var AliasForbidenError = ErrorInfo{
	Code: 40301,
	Msg:  "Alias is used",
}

//Url not found error means the long url is not found in db
var UrlNotFoundError = ErrorInfo{
	Code: 404,
	Msg:  "URL is not found.",
}

//Inter server error means the db has unexpected error
var InternalServerError = ErrorInfo{
	Code: 500,
	Msg:  "Internal Server Error",
}
