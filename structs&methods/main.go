package main

import (
	"fmt"
	"reflect"
)

type ProfilePicture struct {
	Name string
	Path string
}

type Profile struct {
	Name          string
	Username      string
	Designation   string
	ContactNumber string
	ProfilePicture
}

func main() {
	profile := GetProfile("Nikhil", "NikhilSharmaWe", "Product Engineer", "xxxcc", "vvv/dfe/rr.jpg", "999999999")
	fmt.Println(profile)
	mp := Profile{
		Name: "Rewak",
	}
	profile.UpdateProfile(mp)
	fmt.Println(profile)
}

func GetProfile(name, username, designation, imageName, imagePath, contactNumber string) Profile {
	profile := Profile{
		Name:          name,
		Username:      username,
		Designation:   designation,
		ContactNumber: contactNumber,
		ProfilePicture: ProfilePicture{
			Name: imageName,
			Path: imagePath,
		},
	}
	return profile
}

// Fields which are not equal to nil in the modified profile struct will be updated
func (profile *Profile) UpdateProfile(modifiedProfile Profile) {
	if modifiedProfile.Name != "" {
		profile.Name = modifiedProfile.Name
	}
	if modifiedProfile.Username != "" {
		profile.Username = modifiedProfile.Username
	}
	if modifiedProfile.Designation != "" {
		profile.Designation = modifiedProfile.Designation
	}
	if modifiedProfile.ContactNumber != "" {
		profile.ContactNumber = modifiedProfile.ContactNumber
	}
	if (modifiedProfile.ProfilePicture != ProfilePicture{}) {
		profile.ProfilePicture.UpdateProfilePicture(modifiedProfile.ProfilePicture)
	}
}

func (picture *ProfilePicture) UpdateProfilePicture(modifiedProfilePicture ProfilePicture) {
	if modifiedProfilePicture.Name != "" {
		picture.Name = modifiedProfilePicture.Name
	}
	if modifiedProfilePicture.Path != "" {
		picture.Path = modifiedProfilePicture.Path
	}
}

func (profile Profile) CheckDuplicateProfile(profileToBeChecked Profile) bool {
	return reflect.DeepEqual(profile, profileToBeChecked)
}
