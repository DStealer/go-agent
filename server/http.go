package server

import (
	"encoding/json"
	"log"
	"net/http"
	"org/dstealer/agent/composite"
	"org/dstealer/agent/model"
	"strings"
)

func renderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, _ = w.Write(bs)
}

func healthFunc() {
	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("ok"))
	})
	http.HandleFunc("/version", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(composite.Version))
	})
}

func heartbeatFunc() {
	http.HandleFunc("/heartbeat", func(writer http.ResponseWriter, request *http.Request) {
		if request.ContentLength == 0 {
			http.Error(writer, "body is blank", http.StatusBadRequest)
			return
		}
		var req model.HeartbeatRequest
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(writer, "body format error", http.StatusBadRequest)
			return
		}
		log.Print("receive msg:", req)
		if req.Hostname == "" {
			http.Error(writer, "hostname is blank", http.StatusBadRequest)
			return
		}

		if req.Agents != nil {
			agentMap, exits := HostAgents.Get(req.Hostname)
			if exits {
				for _, agent := range req.Agents {
					agentMap.Put(agent.Name, agent)
				}
			} else {
				agentMap = NewAgentMap()
				for _, agent := range req.Agents {
					agentMap.Put(agent.Name, agent)
				}
				HostAgents.Put(req.Hostname, agentMap)
			}
		}

		resp := model.HeartbeatResponse{Code: "NA", Message: "正常", DesiredAgents: nil}
		renderJson(writer, resp)
	})
}

func configFunc() {
	http.HandleFunc("/config", func(writer http.ResponseWriter, request *http.Request) {
		if !strings.HasPrefix(request.RemoteAddr, "127.0.0.1:") {
			_, _ = writer.Write([]byte("no privilege"))
			return
		}
		renderJson(writer, Config())
	})
}

func status() {
	http.HandleFunc("/status/", func(writer http.ResponseWriter, request *http.Request) {
		agentName := request.URL.Path[len("/status/"):]
		if agentName == "" {
			http.Error(writer, "agent name is blank", http.StatusBadRequest)
			return
		}
		data := HostAgents.Status(agentName)
		renderJson(writer, data)
	})
}
