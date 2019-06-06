package main

type Settings struct {
	ParticiPantsFilePath string `json:"particiPantsFilePath"`
	PortNumber           int    `json:"portNumber"`
	IpAddrss             string `json:"portNumber"`
}

func CreateDefaultSettings() Settings {
	return Settings{
		ParticiPantsFilePath: "participants.json",
		PortNumber:           5000,
		IpAddres:             "",
	}
}
