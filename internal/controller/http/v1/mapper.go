package v1

import "lg/internal/entity"

func projectToDTO(project entity.Project) projectDTO {
	return projectDTO{
		Name:             project.Name,
		Description:      project.Description,
		ProjectLink:      project.ProjectLink,
		PresentationLink: project.PresentationLink,
		CreatorID:        project.CreatorID,
	}
}

func projectToEntity(dto projectDTO) entity.Project {
	return entity.Project{
		Name:             dto.Name,
		Description:      dto.Description,
		ProjectLink:      dto.ProjectLink,
		PresentationLink: dto.PresentationLink,
		CreatorID:        dto.CreatorID,
	}
}
