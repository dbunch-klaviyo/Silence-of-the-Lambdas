package silenceofthelambdas

import (
	lambsdb "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambs_db"
	log "github.com/sirupsen/logrus"
)

type Killer struct {
	FullName    string   `json:"full_name"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	MovieActors []string `json:"movie_actors"`
	Movies      []string `json:"movies"`
	Nickname    string   `json:"nickname"`
	Profession  string   `json:"profession"`
}

func (c *SilenceOfTheLambdasClient) ReturnKillerByFullName(fullName string, tableName string, db *lambsdb.SilenceOfTheLambsDB) (*Killer, error) {
	fields := log.Fields{"full_name": fullName, "table_name": tableName}

	dynamoKiller, err := c.db.ReturnKillerByFullName(fullName, tableName, db)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR FETCHING KILLER FROM DYNAMODB: %+v", err)
		return nil, err
	}

	killer := &Killer{
		FullName:    dynamoKiller.FullName,
		FirstName:   dynamoKiller.FirstName,
		LastName:    dynamoKiller.FirstName,
		MovieActors: dynamoKiller.MovieActors,
		Movies:      dynamoKiller.Movies,
		Nickname:    dynamoKiller.Nickname,
		Profession:  dynamoKiller.Profession,
	}

	return killer, nil
}
