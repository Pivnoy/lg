package v1

import (
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

func userToDTO(profile entity.Profile) userDTO {
	return userDTO{
		UUID:       profile.UserUUID.String(),
		FirstName:  profile.Firstname,
		LastName:   profile.Lastname,
		Patronymic: profile.Patronymic.String,
	}
}
