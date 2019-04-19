package server

import (
	"org/dstealer/agent/model"
	"sync"
)

type AgentMap struct {
	sync.RWMutex
	M map[string]*model.RealAgent
}

func NewAgentMap() *AgentMap {
	return &AgentMap{M: make(map[string]*model.RealAgent)}
}

func (m *AgentMap) Get(agentName string) (*model.RealAgent, bool) {
	m.RLock()
	defer m.RUnlock()
	val, exits := m.M[agentName]
	return val, exits
}
func (m *AgentMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.M)
}
func (m *AgentMap) Put(agentName string, realAgent *model.RealAgent) {
	m.Lock()
	defer m.Unlock()
	m.M[agentName] = realAgent
}

type HostAgentMap struct {
	sync.RWMutex
	M map[string]*AgentMap
}

func newHostAgentMap() *HostAgentMap {
	return &HostAgentMap{M: make(map[string]*AgentMap)}
}

func (m *HostAgentMap) Get(hostname string) (*AgentMap, bool) {
	m.RLock()
	defer m.RUnlock()
	val, exits := m.M[hostname]
	return val, exits
}

func (m *HostAgentMap) Put(hostname string, am *AgentMap) {
	m.Lock()
	defer m.Unlock()
	m.M[hostname] = am
}

func (m *HostAgentMap) Status(agentName string) map[string]*model.RealAgent {
	ret := make(map[string]*model.RealAgent)
	m.RLock()
	defer m.RUnlock()
	for hostname, agentMap := range m.M {
		if agent, exits := agentMap.Get(agentName); exits {
			ret[hostname] = agent
		}
	}
	return ret
}

var HostAgents = newHostAgentMap()
