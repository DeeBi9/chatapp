import React from 'react';
import './Auth.css';

async function SubmitForm(event) {
    event.preventDefault(); // Prevent the default form submission behavior
    
    const form = event.target; // Get the form element from the event
    const formData = new FormData(form);

    try {
        const response = await fetch("http://localhost:8080", {
            method: "POST",
            body: formData,
            mode: 'cors', 
        });

        // Ensure the response is okay and parse JSON
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const jsonResponse = await response.json();
        console.log(jsonResponse);
    } catch (e) {
        console.error('There was a problem with the fetch operation:', e);
    }
}


function Auth() {
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
              </div>
          </form>
      </div>
    );
}

export default Auth;
    