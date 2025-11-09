package router

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	agent "github.com/ramaureirac/devops-ragbot/src/pkg/agent"
)

type Sessions struct {
	Agents map[string]*agent.Agent
	Mutex  sync.RWMutex
}

func (c *Sessions) DropSessionsOlderThan(minutes float64) {
	go func() {
		tckr := time.NewTicker(time.Duration(minutes) * time.Minute)
		defer tckr.Stop()
		for {
			dt := <-tckr.C
			c.Mutex.Lock()
			for key, item := range c.Agents {
				if minutes < dt.Sub(item.LastRequest).Minutes() {
					delete(c.Agents, key)
					log.Println("dropped session: " + key)
				}
			}
			c.Mutex.Unlock()
		}
	}()
}

func (c *Sessions) RegisterSession() (string, error) {
	id := uuid.New().String()
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	agt, err := agent.NewAgent()
	if err != nil {
		return "", err
	}
	c.Agents[id] = agt
	return id, nil
}

func (c *Sessions) AskQuestion(id string, question string) (string, error) {
	c.Mutex.Lock()
	agt, ok := c.Agents[id]
	c.Mutex.Unlock()
	if !ok || agt == nil {
		return "", errors.New("agent not found for id: " + id)
	}
	return agt.AskQuestion(question)
}
