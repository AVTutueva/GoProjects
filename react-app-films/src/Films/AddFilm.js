import React, { useState } from "react";
import PropTypes from "prop-types";

function AddFilm({ onCreate }) {
  const [state, setValue] = useState({
    title: "",
    description: ""
  });

  function handleChange(event) {
    const temp = event.target.value;
    setValue({
      ...state,
      [event.target.name]: temp,
    });
  }

  function submitHandler(event) {
    event.preventDefault();
    // check that both fields are filled
    if (state.title.trim() && state.description.trim()) {
      onCreate(state);
      setValue({title: "", description: ""})
      console.log("film created");
    }
    else{
        console.log("error")
    }
    
  }

  return (
    <form onSubmit={submitHandler}>
      <label>
        Title:
        <input name="title" value={state.title} onChange={handleChange} />
        <br />
      </label>
      <label>
        Description:
        <input name="description" value={state.description} onChange={handleChange} />
        <br />
      </label>
      <button type="submit">Add</button>
    </form>
  );
}

AddFilm.propTypes = {
    onCreate: PropTypes.func.isRequired
}

export default AddFilm;

//   function submitHandler(event) {
//     event.preventDefault() // cancel updating of page
//     if (value.title.trim() && value.description.trim()){
//         onCreate(value)
//     }

//   }

// function handleChange (event) {
//     setValue({...value, [event.target.name]: event.target.value});
// };
