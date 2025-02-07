package model

type Officer struct {
	Name        string `json:"name"`         // get from WaterChannelOfficer.Name where WaterChannelOfficer.WaterChannelDoorID = x
	Photo       string `json:"photo"`        // get from WaterChannelOfficer.Photo where WaterChannelOfficer.WaterChannelDoorID = x
	PhoneNumber string `json:"phone_number"` // get from WaterChannelOfficer.PhoneNumber where WaterChannelOfficer.WaterChannelDoorID = x
	Task        string `json:"task"`         // get from WaterChannelOfficer.Task where WaterChannelOfficer.WaterChannelDoorID = x
}
