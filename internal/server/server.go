package server

import (
	"encoding/json"
	"net/http"

	"yanceGG/internal/agent"
)

// Server exposes HTTP handlers for the agent framework.
type Server struct {
	Agent agent.Agent
}

// New constructs a Server with the provided agent.
func New(agent agent.Agent) *Server {
	return &Server{Agent: agent}
}

// Routes registers handlers on a ServeMux.
func (s *Server) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/skills", s.handleSkills)
	mux.HandleFunc("/plan", s.handlePlan)
	mux.HandleFunc("/execute", s.handleExecute)
	return mux
}

type skillResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *Server) handleSkills(w http.ResponseWriter, _ *http.Request) {
	responses := make([]skillResponse, 0)
	for _, skill := range s.Agent.Skills.All() {
		responses = append(responses, skillResponse{
			Name:        skill.Name(),
			Description: skill.Description(),
		})
	}

	writeJSON(w, http.StatusOK, responses)
}

type planRequest struct {
	Goal string `json:"goal"`
}

func (s *Server) handlePlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req planRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	plan, err := s.Agent.BuildPlan(r.Context(), req.Goal)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, plan)
}

type executeRequest struct {
	Plan agent.Plan `json:"plan"`
}

func (s *Server) handleExecute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req executeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	results, err := s.Agent.ExecutePlan(r.Context(), req.Plan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, results)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
