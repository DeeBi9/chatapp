import React from 'react';
import './Auth.css';
import {useNavigate } from "react-router-dom";
var jsonResponse = {}

var responseData = {}
function recordResponse(jsonResponse) {
    responseData = {
        'Id' : jsonResponse.data.Id,
        'Username' : jsonResponse.data.username,
    }
    console.log(responseData)
}

async function SubmitForm(event) {
    event.preventDefault(); // Prevent the default form submission behavior
    
    const form = event.target; // Get the form element from the event
    const formData = new FormData(form);
    const plainObject = Object.fromEntries(formData.entries());
    const jsonData = JSON.stringify(plainObject);
    console.log(jsonData);

    try {
        const response = await fetch("http://localhost:8080/authorization", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json' // Ensure the request has the correct content type
            },
            body: jsonData,
            mode: 'cors',
        });

        // Handle the response based on status code
        if (response.ok) {
            jsonResponse = await response.json();
            console.log(jsonResponse);
            // Handle successful registration
            alert("User registered successfully.");
            alert("This is your id make sure to copy this:",jsonResponse.data.Id)
        } else if (response.status === 409) {
            jsonResponse = await response.json();
            console.error(jsonResponse.message); // Log or display the conflict message
            alert("Username is already taken. Please choose a different username.");
        } else {
            jsonResponse = await response.json();
            console.error(jsonResponse.message); // Log or display other error messages
            alert("An error occurred: " + jsonResponse.message);
        }
        
    } catch (e) {    
        console.error('Problem occured while fetching the information:', e);
        alert('There was a problem with the fetch operation.');
    }
    recordResponse(jsonResponse)
}



function Auth() {
    const navigate = useNavigate();
    function redirectTo () {
        navigate("/signin")
    }
    
    return (
      <div id="form-container">
          <form id="form" onSubmit={SubmitForm}>
              <div id="components">
                  <div>
                      <label htmlFor="username" className="username">Username</label>
                      <input 
                          type="text" 
                          name="username" 
                          required 
                          placeholder="Username" 
                      />
                  </div>
                  <br /><br />
                  <div>
                      <label htmlFor="password" className="password">Password</label>
                      <input 
                          type="password" 
                          name="password" 
                          required 
                          placeholder="********" 
                          minLength="8" 
                      />
                  </div>
                  <br /><br />
              </div>
              <div id="footer">
                  <div id="button-container">
                      <button type="submit" id="button">Create Account</button>
                  </div>
                  <br></br>
                  <div id='signin'>
                    <button id='button' onClick={redirectTo}>Sign IN</button>
                  </div>
              </div>
          </form>

      </div>
    );
}

export default Auth;
    