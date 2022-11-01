package cli

import (
	"strings"

	"mun/x/claim/types"
)

// NormalizeConditionType normalizes specified condition type.
func NormalizeConditionType(ConditionType string) types.ConditionType {
	switch strings.ToLower(ConditionType) {
	case "i", "init":
		return types.ConditionTypeInit
	case "s", "swap":
		return types.ConditionTypeSwap
	case "ls", "stake":
		return types.ConditionTypeStake
	case "v", "vote":
		return types.ConditionTypeVote
	default:
		return types.ConditionTypeUnspecified
	}
}
