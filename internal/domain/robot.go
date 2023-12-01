package domain

import "github.com/google/uuid"

type Robot struct {
	Id     uuid.UUID
	Name   string
	ApiKey string `db:"api_key"`
}

type RobotRepository interface {
	GetRobotById(id uuid.UUID) (*Robot, error)
	AddRobot(robot *Robot) error
}
