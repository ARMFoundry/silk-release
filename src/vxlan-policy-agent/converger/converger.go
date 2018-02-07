package converger

import (
	"fmt"
	"time"
	"vxlan-policy-agent/enforcer"

	"code.cloudfoundry.org/lager"
)

//go:generate counterfeiter -o fakes/planner.go --fake-name Planner . Planner
type Planner interface {
	GetRulesAndChain() (enforcer.RulesWithChain, error)
}

//go:generate counterfeiter -o fakes/rule_enforcer.go --fake-name RuleEnforcer . ruleEnforcer
type ruleEnforcer interface {
	EnforceRulesAndChain(enforcer.RulesWithChain) error
}

//go:generate counterfeiter -o fakes/metrics_sender.go --fake-name MetricsSender . metricsSender
type metricsSender interface {
	SendDuration(string, time.Duration)
}

type SinglePollCycle struct {
	Planners      []Planner
	Enforcer      ruleEnforcer
	MetricsSender metricsSender
	Logger        lager.Logger
	ruleSets      map[enforcer.Chain]enforcer.RulesWithChain
}

const metricEnforceDuration = "iptablesEnforceTime"
const metricPollDuration = "totalPollTime"

func (m *SinglePollCycle) DoCycle() error {
	if m.ruleSets == nil {
		m.ruleSets = make(map[enforcer.Chain]enforcer.RulesWithChain)
	}

	pollStartTime := time.Now()
	var enforceDuration time.Duration
	for _, p := range m.Planners {
		ruleSet, err := p.GetRulesAndChain()
		if err != nil {
			return fmt.Errorf("get-rules: %s", err)
		}
		enforceStartTime := time.Now()

		oldRuleSet := m.ruleSets[ruleSet.Chain]
		if !ruleSet.Equals(oldRuleSet) {
			m.Logger.Debug("poll-cycle", lager.Data{
				"message":       "updating iptables rules",
				"num old rules": len(oldRuleSet.Rules),
				"num new rules": len(ruleSet.Rules),
				"old rules":     oldRuleSet,
				"new rules":     ruleSet,
			})
			err = m.Enforcer.EnforceRulesAndChain(ruleSet)
			if err != nil {
				return fmt.Errorf("enforce: %s", err)
			}
			m.ruleSets[ruleSet.Chain] = ruleSet
		}

		enforceDuration += time.Now().Sub(enforceStartTime)
	}

	pollDuration := time.Now().Sub(pollStartTime)
	m.MetricsSender.SendDuration(metricEnforceDuration, enforceDuration)
	m.MetricsSender.SendDuration(metricPollDuration, pollDuration)

	return nil
}
