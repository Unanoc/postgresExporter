package config

// Config is struct for parsing of configuration file.
type Config struct {
	Connection string   `json:"connection"`
	OutputDir  string   `json:"output_dir"`
	Tables     []*Table `json:"tables"`
}

// Table is a struct for exporting.
type Table struct {
	Name     string `json:"name"`
	Query    string `json:"query"`
	MaxLines int    `json:"max_lines"`
}
