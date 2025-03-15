package tester

import (
	"context"

	"github.com/jijimama/newspaper-app/ent"
	"github.com/jijimama/newspaper-app/infrastructure/database"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type DBPostgresSuite struct {
	suite.Suite
	postgresContainer testcontainers.Container
	ctx               context.Context
	DB                *ent.Client
}

func (suite *DBPostgresSuite) SetupTestContainers() error {
	suite.ctx = context.Background()

	config := database.NewConfigPostgres()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       config.Name,
			"POSTGRES_USER":     config.User,
			"POSTGRES_PASSWORD": config.Password,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(suite.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return err
	}

	suite.postgresContainer = container
	return nil
}

func (suite *DBPostgresSuite) SetupSuite() {
	err := suite.SetupTestContainers()
	suite.Require().NoError(err)

	db, err := database.NewDatabaseSQLFactory(database.InstancePostgres)
	suite.Require().NoError(err)

	suite.DB = db
}

func (suite *DBPostgresSuite) TearDownSuite() {
	if suite.postgresContainer != nil {
		err := suite.postgresContainer.Terminate(suite.ctx)
		suite.Require().NoError(err)
	}
}
