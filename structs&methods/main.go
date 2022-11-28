package main

import "fmt"

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
	modifiedProfile := GetProfile("Sharma", "NikhilSharmaWe", "Product Engineer", "xxxcc", "vvv/dfe/rr.jpg", "999999999")
	profile.UpdateProfile(modifiedProfile)
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

func (profile *Profile) UpdateProfile(modifiedProfile Profile) {
	*profile = modifiedProfile
}

func (picture *ProfilePicture) UpdateProfilePicture(modifiedProfilePicture ProfilePicture) {
	*picture = modifiedProfilePicture
}

func (profile Profile) CheckDuplicateProfile(profileToBeChecked Profile) bool {
	return (profile == profileToBeChecked)
}
