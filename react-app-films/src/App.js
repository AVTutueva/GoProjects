import React, { useEffect } from "react";
// import React, {useEffect, useState} from "react";
import FilmTable from "./Films/FilmTable";
import Context from "./Films/context";
import AddFilm from "./Films/AddFilm";
import axios from "axios";

function App() {
  const [films, setFilms] = React.useState({});
  // const [films, setFilms] = React.useState([
  //   { id: 1, title: "Titanic", description: "simple desc1" },
  //   { id: 2, title: "Avatar", description: "simple desc2" },
  //   { id: 3, title: "The Lord of the Rings", description: "simple desc3" },
  // ]);

  useEffect(() => {
    axios.get("http://localhost:8080/films").then((response) =>
      setFilms(
        response.data.map(function (element) {
          const new_film = {
            id: element.FilmId,
            title: element.Title,
            description: element.Description,
          };
          return new_film;
        })
      )
    );
  }, []);


  function removeFilm(id) {
    // make error handling
    try {
      const request_to_delete = `http://localhost:8080/films/${id}`;
      fetch(request_to_delete, {
        method: "DELETE",
        headers: { "content-type": "application/json" },
      })
        .then((data) => data.json())
        .then((data) => {
          setFilms(films.filter((film) => film.id !== id));
        });
    } catch (error) {
      console.log(error);
    }
  }

  // adding a new film
  function addFilm(state) {
    const jsonFilm = {
      Title: state.title,
      Description: state.description,
      ReleaseYear: parseInt(state.year),
      LanguageId: 1,
      OriginalLanguageId: 0,
      RentalDuration: 7,
      RentalRate: 2.99,
      Length: 136,
      ReplacementCost: 14.99,
      Rating: "PG",
      SpecialFeatures: "Trailers",
      Categories: [],
    };

    fetch("http://localhost:8080/films/1", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(jsonFilm),
    }).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return Promise.reject(response);
    })
    .then((response) => {
        setFilms(
          films.concat([{
              id: response.FilmId,
              title: state.title,
              description: state.description,
            },])
        );
      })
      .catch((error) =>{
        alert(error.statusText)
        console.log("Error", error.statusText)
      });
  }

  return (
    <Context.Provider value={{ removeFilm }}>
      <div className="wrapper">
        <h1>List of films</h1>
        <AddFilm onCreate={addFilm} />
        {films.length ? <FilmTable films={films} /> : <p>No films to show</p>}
      </div>
    </Context.Provider>
  );
}

export default App;
