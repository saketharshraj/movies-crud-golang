package main

// Movie represents a film with an ID, ISBN, title, and director.
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director represents the director of a movie.
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
