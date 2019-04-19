package model

type HeartbeatRequest struct {
	Hostname string       `json:"hostname"`
	Agents   []*RealAgent `json:"agents"`
}

type HeartbeatResponse struct {
	Code          string          `json:"code"`
	Message       string          `json:"message"`
	DesiredAgents []*DesiredAgent `json:"desired_agents"`
}
