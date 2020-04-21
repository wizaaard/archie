package models

import (
	"archie/connection/postgres_conn"
	"archie/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type UserOrganization struct {
	UserID         string `gorm:"type:uuid;primary_key;"`
	OrganizationID string `gorm:"type:uuid;primary_key;"`
	User           *User
	Organization   *Organization
	IsOwner        bool  `gorm:"type:bool"`
	JoinTime       int64 `gorm:"type:bigint"`
}

func (userOrganization *UserOrganization) TableName() string {
	return "user_organizations"
}

// 寻找指定 Organization ID 的 members
func (userOrganization *UserOrganization) FindMembers(organizationID string, members *[]User) error {
	var userOrganizations []UserOrganization

	err := postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		return db.Model(userOrganization).Preload("User").Where("organization_id = ?", organizationID).Find(&userOrganizations).Error
	})

	utils.ArrayMap(userOrganizations, func(item interface{}) interface{} {
		return *item.(UserOrganization).User
	}, members)

	return err
}

func (userOrganization *UserOrganization) New(isOwner bool) error {
	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		userOrganization.JoinTime = utils.Now()
		userOrganization.IsOwner = isOwner

		return db.Create(userOrganization).Error
	})
}

type OrganizationOwnerInfo struct {
	OwnerInfo User `json:"ownerInfo"`
	Organization
	Members  []User
	JoinTime int64 `json:"joinTime"`
}

func findOwnerByID(id string, owners []User) (User, bool) {
	for _, owner := range owners {
		if owner.ID == id {
			return owner, true
		}
	}

	return User{}, false
}

func (userOrganization *UserOrganization) FindUserJoinOrganizations() ([]OrganizationOwnerInfo, error) {
	var infos []OrganizationOwnerInfo

	err := postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		var result *gorm.DB

		result = db.
			Raw(
				"select * from user_organizations inner join organizations on organization_id=organizations.id where user_id=?",
				userOrganization.UserID,
			).
			Scan(&infos)

		if len(infos) == 0 {
			return nil
		}

		var organizationIds []string
		for _, info := range infos {
			organizationIds = append(organizationIds, fmt.Sprintf("'%s'", info.Owner))
		}

		var owners []User

		result = db.
			Raw(fmt.Sprintf("select * from users where id in (%s)", strings.Join(organizationIds, ","))).
			Scan(&owners)

		// dist organizations
		for i, organization := range infos {
			owner, ok := findOwnerByID(organization.Owner, owners)

			if !ok {
				continue
			}

			// attach members
			var members []User
			if err := userOrganization.FindMembers(infos[i].ID, &members); err == nil {
				infos[i].Members = members
			} else {
				infos[i].Members = []User{}
			}

			infos[i].OwnerInfo = owner
		}

		return result.Error
	})

	return infos, err
}
