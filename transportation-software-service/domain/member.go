package domain

import "github.com/google/uuid"

var (
	Owner    MemberRole = "owner"    // do everything kick and invite
	Admin    MemberRole = "admin"    // read & write (every channel)
	Operator MemberRole = "operator" // read & write (some channel)
	Guess    MemberRole = "guess"    // read only (published channel)
)

type MemberRole string

type Member struct {
	MemberID  uuid.UUID  `bson:"member_id" json:"member_id"`
	Role      MemberRole `bson:"role" json:"role"`
	PartnerID uuid.UUID  `bson:"partner_id" json:"partner_id"`
}
