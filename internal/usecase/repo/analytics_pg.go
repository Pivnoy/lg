package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type AnalyticsRepo struct {
	*postgres.Postgres
}

var _ usecase.AnalyticsRp = (*AnalyticsRepo)(nil)

func NewAnalyticsRepo(pg *postgres.Postgres) *AnalyticsRepo {
	return &AnalyticsRepo{pg}
}

//Вакансии на проект
//select distinct * from (select lineup.role_uuid from project join
//(select category_uuid
//from project
//where project.uuid = '{project_uuid}') as ctg
//on
//project.category_uuid = ctg.category_uuid
//join lineup on project.uuid = lineup.project_uuid
//where project.uuid <> '{project_uuid}'
//) as foo
/////
//select uuid from public.role;

func (a *AnalyticsRepo) GetUserUUIDListOfVacsProject(ctx context.Context, project uuid.UUID) ([]uuid.UUID, error) {
	query := `select distinct * from (select lineup.role_uuid from project join (select category_uuid from project
		where project.uuid = $1) as ctg on project.category_uuid = ctg.category_uuid
		join lineup on project.uuid = lineup.project_uuid where project.uuid <> $1 ) as foo`

	rows, err := a.Pool.Query(ctx, query, project)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	var usersUUID []uuid.UUID
	for rows.Next() {
		var ud uuid.UUID
		err = rows.Scan(&ud)
		if err != nil {
			return nil, fmt.Errorf("cannot parse struct: %v", err)
		}
		usersUUID = append(usersUUID, ud)
	}
	return usersUUID, nil
}

//select project_uuid from profile join public.role on profile.specialization_uuid = public.role.specialization_uuid join
//(SELECT project.uuid as project_uuid, role_uuid
//from project join lineup on lineup.project_uuid = project.uuid
//where profile_uuid IS NULL) as lineup_needs on public.role.uuid = lineup_needs.role_uuid where user_uuid = '{user_uuid}';
/////
//SELECT project.uuid as project_uuid
//from project join lineup on lineup.project_uuid = project.uuid
//where profile_uuid IS NULL;

//Люди на свободные вакансии
//select user_uuid from
//(select user_uuid, public.role.uuid as role_uuid from profile join public.role ON profile.specialization_uuid = public.role.specialization_uuid) as tb1
//join
//(select project_uuid, role_uuid from project join lineup on project.uuid = lineup.project_uuid where public.lineup.profile_uuid is NULL) as tb2
//on
//tb1.role_uuid = tb2.role_uuid
//where
//project_uuid = '{project_uuid}';
/////
//select user_uuid from profile;

func (a *AnalyticsRepo)