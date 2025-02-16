package follows

type followsError struct{ msg string }

func (e followsError) Error() string { return e.msg }

var ErrFollowInvalidArgs = followsError{msg: "error: follow args invalid or missing"}
var ErrURLFeedRetrieve = followsError{msg: "error: feed at the provided address does not exist"}
var ErrFollowCreation = followsError{msg: "error: failed to create follow"}
var ErrFollowingInvalidArgs = followsError{msg: "error: following args invalid or missing"}
var ErrRetrieveFeedFollows = followsError{msg: "error: failed to retrieve user's follows"}
var ErrUnfollowInvalidArgs = followsError{msg: "error: unfollow args invalid or missing"}
var ErrDeleteFeedFollow = followsError{msg: "error: failed to unfollow user from a feed"}
