package valueobject

type Status string

var (
	Pending   Status = "pending"   // waiting for accespt order request
	Accepted  Status = "accepted"  // accepted request by operator
	Cancelled Status = "cancelled" // request was cancelld
	Operating Status = "operating" // operating during sending or operating
	Successed Status = "successed" // package was sent successfully
	Failed    Status = "failed"    // package sent failed
)
