package model

type BmkgLocation struct {
	Adm4        string `json:"adm4"`
	Province    string `json:"provinsi"`
	City        string `json:"kotkab"`
	District    string `json:"kecamatan"`
	Subdistrict string `json:"desa"`
}
