// import React, {useEffect, useState} from "react";
import React from "react";
import FilmTable from "./Films/FilmTable";
import Context from "./Films/context";
import AddFilm from "./Films/AddFilm";
// import axios from "axios";

function App() {

  const [films, setFilms] = React.useState([
    { id: 1, title: "Titanic", description: "simple desc1" },
    { id: 2, title: "Avatar", description: "simple desc2" },
    { id: 3, title: "The Lord of the Rings", description: "simple desc3" },
  ]);


  // const fetchData = () => {
  //   return axios.get("http://localhost:8080/films")
  //         .then((response) => console.log(response.data));
  // }

  // useEffect(() => {
  //   fetchData();
  // },[])


  function removeFilm(id) {
    setFilms(
      // to future: remove film from database, get new list of films and return
      films.filter((film) => film.id !== id)
    );
  }

  function addFilm(state){
    setFilms(films.concat([
      {
        id: Date.now(),
        title: state.title,
        description: state.description
      }
    ])
    )
  }


  return (
    <Context.Provider value={{ removeFilm }}>
      <div className="wrapper">
        <h1>List of films</h1>
        <AddFilm onCreate={addFilm}/>
        {films.length ? <FilmTable films={films} /> : <p>No films to show</p>}
      </div>
    </Context.Provider>
  );
}

export default App;
