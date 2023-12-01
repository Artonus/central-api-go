package repository

import (
	"errors"
	"fmt"
	"github.com/Artonus/central-api-go/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

type robotRepository struct {
	DB *sqlx.DB
}

func (r *robotRepository) GetRobotById(id uuid.UUID) (*domain.Robot, error) {
	var robot domain.Robot
	err := r.DB.Select(&robot, "select * from robots where id=$1", id)
	if err != nil {
		return nil, err
	}
	if (domain.Robot{}) == robot {
		return nil, nil
	}
	return &robot, nil
}

func (r *robotRepository) AddRobot(robot *domain.Robot) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	robot.Id = id

	_, err = r.DB.NamedExec("insert into robots(id, name, api_key) VALUES (:id, :name, :api_key)", &robot)

	if err != nil {
		if strings.Contains(err.Error(), "robots_name_unique_idx") {
			return errors.New(fmt.Sprintf("there is already a robot with name \"%s\"", robot.Name))
		}
		return err
	}

	return nil
}

func NewRobotRepository(db *sqlx.DB) domain.RobotRepository {
	return &robotRepository{
		db,
	}
}
