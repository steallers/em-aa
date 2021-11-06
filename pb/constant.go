package pb

const (
	RequestSuccess      = "success"
	RequestFailed       = "failed"
	RequestNeedApproval = "need approval"

	NewRequestMark = "new request to '%s' from user-id (%d) with params: \n %+v"
	InvalidInput   = "invalid input '%s' with format (%s), required format is (%s)"
)
