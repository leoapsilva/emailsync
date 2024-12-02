package model

type Connection struct {
	URL string
}

func (con Connection) FormatURL(path string) string {
	return con.URL + path
}
