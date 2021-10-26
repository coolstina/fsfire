package fsfire

type FSPath string

func (p FSPath) String() string {
	return string(p)
}

const GlobalDefaultDir FSPath = "static"

// Assets files save path definition.
const (
	AssetsSavePathForDatabase FSPath = "static/assets/database"
	AssetsSavePathForExcel    FSPath = "static/assets/excel"
	AssetsSavePathForCSV      FSPath = "static/assets/csv"
)
