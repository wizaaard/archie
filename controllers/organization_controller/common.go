package organization_controller

import "archie/models"

func CreateNewOrganization(name string, description string, username string) error {
	organization := models.Organization{
		OrganizeName: name,
		Description:  description,
	}

	return organization.New(username)
}

func InsertUserToOrganization(organizeName string, username string, isOwner bool) error {
	organization := models.Organization{OrganizeName: organizeName}
	organization.FindOneByOrganizeName()

	user := models.FindOneByUsername(username)
	userOrganization := models.UserOrganization{
		UserID:         user.ID,
		OrganizationID: organization.ID,
	}

	return userOrganization.New(isOwner)
}
