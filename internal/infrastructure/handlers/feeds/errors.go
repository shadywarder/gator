package feeds

type feedsError struct{ msg string }

func (e feedsError) Error() string { return e.msg }

var ErrAggIntervalAbsence = feedsError{msg: "error: time interval is not specified"}
var ErrAggInvalidTimeFormat = feedsError{msg: "error: time format is invalid"}
var ErrRSSWrongTimeFormat = feedsError{msg: "error: rss time format is invalid"}
var ErrAddFeedInvalidArgs = feedsError{msg: "error: addfeed args is invalid or missing"}
var ErrFeedCreation = feedsError{msg: "error: failed to create feed"}
var ErrFeedFollow = feedsError{msg: "error: failed to create follow"}
var ErrRetrieveFeeds = feedsError{msg: "error: failed to retrieve feeds"}
var ErrEmptyFeeds = feedsError{msg: "error: no feeds found"}
var ErrUnknownUser = feedsError{msg: "error: failed to found user with specified id"}
var ErrInvalidBrowseArgs = feedsError{msg: "error: browse args is invalid or missing"}
