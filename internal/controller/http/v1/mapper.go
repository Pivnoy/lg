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

func messageToDTO(message entity.Message) messageDTO {
	return messageDTO{
		Content: message.Content,
		Sender:  message.AuthorUUID.String(),
		Date:    message.CreationDate.String(),
	}
}

func chatItemToDTO(chatItem entity.ChatItem) chatItemDTO {
	return chatItemDTO{
		ChatName:    chatItem.ChatName,
		ChatUUID:    chatItem.ChatUUID.String(),
		LastMessage: messageToDTO(chatItem.LastMessage),
		ImageURL:    "",
	}
}
