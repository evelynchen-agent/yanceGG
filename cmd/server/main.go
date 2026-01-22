package main

import (
	"log"
	"net/http"

	"yanceGG/internal/agent"
	"yanceGG/internal/server"
	"yanceGG/internal/skills"
)

func main() {
	skillSet := agent.NewSkillSet()
	if err := skillSet.Register(skills.EchoSkill{}); err != nil {
		log.Fatal(err)
	}
	if err := skillSet.Register(skills.UppercaseSkill{}); err != nil {
		log.Fatal(err)
	}

	app := agent.Agent{
		Planner: agent.BasicPlanRole{},
		Skills:  skillSet,
	}

	srv := server.New(app)

	addr := ":8080"
	log.Printf("agent server listening on %s", addr)
	if err := http.ListenAndServe(addr, srv.Routes()); err != nil {
		log.Fatal(err)
	}
}
