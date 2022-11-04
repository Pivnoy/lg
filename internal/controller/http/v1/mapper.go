package v1

import "lg/internal/entity"

func projectToDTO(project entity.Project) projectDTO {
	return projectDTO{
		UUID:             project.UUID,
		Name:             project.Name,
		Description:      project.Description,
		CategoryUUID:     project.CategoryUUID,
		ProjectLink:      project.ProjectLink,
		PresentationLink: project.PresentationLink,
		CreatorUUID:      project.CreatorUUID,
		IsVisible:        project.IsVisible,
	}
}

func projectToEntity(dto projectDTO) entity.Project {
	return entity.Project{
		UUID:             dto.UUID,
		Name:             dto.Name,
		Description:      dto.Description,
		CategoryUUID:     dto.CategoryUUID,
		ProjectLink:      dto.ProjectLink,
		PresentationLink: dto.PresentationLink,
		CreatorUUID:      dto.CreatorUUID,
		IsVisible:        dto.IsVisible,
	}
}

func countryToDTO(country entity.Country) countryDTO {
	return countryDTO{
		UUID: country.UUID.String(),
		Name: country.Name,
		Code: country.Code,
	}
}

func citizenshipToDTO(citizenship entity.Citizenship) citizenshipDTO {
	return citizenshipDTO{
		UUID: citizenship.UUID.String(),
		Name: citizenship.Name,
	}
}

func universityToDTO(university entity.University) universityDTO {
	return universityDTO{
		UUID: university.UUID.String(),
		Name: university.Name,
	}
}

func eduspecialityToDTO(eduspeciality entity.Eduspeciality) eduspecialityDTO {
	return eduspecialityDTO{
		UUID: eduspeciality.UUID.String(),
		Name: eduspeciality.Name,
		Code: eduspeciality.Code,
	}
}

func employmentToDTO(employment entity.Employment) employmentDTO {
	return employmentDTO{
		UUID:  employment.UUID.String(),
		Name:  employment.Name,
		Value: employment.Value,
	}
}

func specializationToDTO(specialization entity.Specialization) specializationDTO {
	return specializationDTO{
		UUID:  specialization.UUID.String(),
		Name:  specialization.Name,
		Value: specialization.Value,
	}
}

func cityToDTO(city entity.City) cityDTO {
	return cityDTO{
		UUID: city.UUID.String(),
		Name: city.Name,
	}
}

func categoryToDTO(category entity.Category) categoryDTO {
	return categoryDTO{
		UUID: category.UUID.String(),
		Name: category.Name,
	}
}
