package handlers

import (
	"MusicBands/src/structure"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var db *sql.DB

func Initialize(database *sql.DB) {
	db = database
}

func CreateBand(w http.ResponseWriter, r *http.Request) {
	var band structure.Band
	err := json.NewDecoder(r.Body).Decode(&band)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	stmt, err := db.Prepare("INSERT INTO bands(name, country, debut_year) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(band.Name, band.Country, band.Debut_Year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// func GetBands(w http.ResponseWriter, r *http.Request) {
//     rows, err := db.Query("SELECT id, name, country, debut_year FROM bands")
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     defer rows.Close()

//	    var bands []structure.Band
//	    for rows.Next() {
//	        var band structure.Band
//	        if err := rows.Scan(&band.ID, &band.Name, &band.Country, &band.Debut_Year); err != nil {
//	            http.Error(w, err.Error(), http.StatusInternalServerError)
//	            return
//	        }
//	        bands = append(bands, band)
//	    }
//	    json.NewEncoder(w).Encode(bands)
//	}
func GetBands(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	country := r.URL.Query().Get("country")
	debutYear := r.URL.Query().Get("debut_year")
	sort := r.URL.Query().Get("sort") 

	query := "SELECT id, name, country, debut_year FROM bands WHERE 1=1"
	args := []interface{}{}

	if name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+name+"%")
	}
	if country != "" {
		query += " AND country LIKE ?"
		args = append(args, "%"+country+"%")
	}
	if debutYear != "" {
		query += " AND debut_year = ?"
		args = append(args, debutYear)
	}
	if sort != "" {
		order := "ASC"
		field := sort
		if strings.HasPrefix(sort, "-") {
			order = "DESC"
			field = sort[1:]
		}
		switch field {
		case "name", "country", "debut_year":
			query += " ORDER BY " + field + " " + order
		}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bands []structure.Band
	for rows.Next() {
		var band structure.Band
		if err := rows.Scan(&band.ID, &band.Name, &band.Country, &band.Debut_Year); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bands = append(bands, band)
	}
	json.NewEncoder(w).Encode(bands)
}

func UpdateBand(w http.ResponseWriter, r *http.Request) {
	var band structure.Band
	err := json.NewDecoder(r.Body).Decode(&band)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]

	stmt, err := db.Prepare("UPDATE bands SET name = ?, country = ?, debut_year = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(band.Name, band.Country, band.Debut_Year, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteBand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	stmt, err := db.Prepare("DELETE FROM bands WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
