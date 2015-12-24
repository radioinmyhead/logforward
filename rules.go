package forward

type Rules map[string]struct {
	Sms  []string
	Rule []string
}

var wbRules Rules

func SetRule(r Rules) {
	wbRules = r
}
