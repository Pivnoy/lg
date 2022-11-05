package v1

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
)

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

func messageToDTO(message entity.Message, usDto userDTO) messageDTO {
	return messageDTO{
		Content: message.Content,
		Sender:  usDto,
		Date:    message.CreationDate.String(),
	}
}

func chatItemToDTO(chatItem entity.ChatItem, msg entity.Message, us userDTO) chatItemDTO {
	prj := chatItem.ProjectUUID.String()
	if chatItem.ProjectUUID == uuid.Nil {
		prj = ""
	}
	return chatItemDTO{
		ChatName:    chatItem.ChatName,
		ChatUUID:    chatItem.ChatUUID.String(),
		LastMessage: messageToDTO(msg, us),
		ProjectUUID: prj,
		ImageURL:    "",
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

func profileToEntity(dto profileRequestDTO) (entity.Profile, error) {
	var (
		universityUUID    uuid.UUID
		eduspecialityUUID uuid.UUID
		teamUUID          uuid.UUID

		err error
	)

	if dto.UniversityUUID == "" {
		universityUUID = uuid.Nil
	} else {
		universityUUID, err = uuid.Parse(dto.UniversityUUID)
		if err != nil {
			return entity.Profile{}, fmt.Errorf("error parsing university uuid: %w", err)
		}
	}

	if dto.EduspecialityUUID == "" {
		eduspecialityUUID = uuid.Nil
	} else {
		eduspecialityUUID, err = uuid.Parse(dto.EduspecialityUUID)
		if err != nil {
			return entity.Profile{}, fmt.Errorf("error parsing eduspeciality uuid: %w", err)
		}
	}

	if dto.TeamUUID == "" {
		teamUUID = uuid.Nil
	} else {
		teamUUID, err = uuid.Parse(dto.TeamUUID)
		if err != nil {
			return entity.Profile{}, fmt.Errorf("error parsing team uuid: %w", err)
		}
	}

	if dto.TeamUUID == "" {
		teamUUID = uuid.Nil
	} else {
		teamUUID, err = uuid.Parse(dto.TeamUUID)
		if err != nil {
			return entity.Profile{}, fmt.Errorf("error parsing team uuid: %w", err)
		}
	}

	userUUID, err := uuid.Parse(dto.UserUUID)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("error parsing user uuid: %w", err)
	}
	countryUUID, err := uuid.Parse(dto.CountryUUID)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("error parsing country uuid: %w", err)
	}
	cityUUID, err := uuid.Parse(dto.CityUUID)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("error parsing city uuid: %w", err)
	}
	citizenshipUUID, err := uuid.Parse(dto.CitizenshipUUID)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("error parsing citizenship uuid: %w", err)
	}
	employmentUUID, err := uuid.Parse(dto.EmploymentUUID)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("error parsing employment uuid: %w", err)
	}
	achievementUUID, err := uuid.Parse(dto.AchievementUUID)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("error parsing achievement uuid: %w", err)
	}

	specializationUUID, err := uuid.Parse(dto.SpecializationUUID)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("error parsing specialization uuid: %w", err)
	}
	newPatr := sql.NullString{}
	if dto.Patronymic != "" {
		newPatr.String = dto.Patronymic
		newPatr.Valid = true
	} else {
		newPatr.Valid = false
	}
	return entity.Profile{
		UserUUID:           userUUID,
		Firstname:          dto.Firstname,
		Lastname:           dto.Lastname,
		Patronymic:         newPatr,
		CountryUUID:        countryUUID,
		CityUUID:           cityUUID,
		CitizenshipUUID:    citizenshipUUID,
		Gender:             dto.Gender,
		Phone:              dto.Phone,
		Email:              dto.Email,
		UniversityUUID:     universityUUID,
		EduspecialityUUID:  eduspecialityUUID,
		GraduationYear:     dto.GraduationYear,
		EmploymentUUID:     employmentUUID,
		Experience:         dto.Experience,
		AchievementUUID:    achievementUUID,
		TeamUUID:           teamUUID,
		SpecializationUUID: specializationUUID,
	}, nil
}

func profileToDTO(profile entity.Profile) profileResponseDTO {
	return profileResponseDTO{
		UUID:       profile.UserUUID.String(),
		Firstname:  profile.Firstname,
		Lastname:   profile.Lastname,
		Patronymic: profile.Patronymic.String,
	}
}
func userToDTO(profile entity.Profile) userDTO {
	return userDTO{
		UUID:       profile.UserUUID.String(),
		FirstName:  profile.Firstname,
		LastName:   profile.Lastname,
		Patronymic: profile.Patronymic.String,
	}
}
