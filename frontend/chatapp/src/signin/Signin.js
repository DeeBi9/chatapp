import React, { useState } from 'react';
import { useNavigate } from "react-router-dom";
import './Signin.css';

var jsonResponse = {}

// Sends the request to backend server to check for correct Id and Password
// If correct it returns with a JWT token if incorrect it returns with a error
// Incorrect Password or Incorrect Id.
async function CheckForIDPass(jsonData) {

        try {
            const response = await fetch("http://localhost:8080/authentication", {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json' 
                },
                body: jsonData,
                mode: 'cors',
            });
            if (response.ok) {
                jsonResponse = await response.json();
                console.log(jsonResponse);
                // Handle successful ID authentication
                alert("Welcome. Click Ok to proceed")
            } else {
                jsonResponse = await response.json();
                console.error(jsonResponse.message); // Log or display other error messages
                alert("An error occurred: " + jsonResponse.message);
            }
            
        } catch (e) {    
            console.error('Problem occured while fetching the information:', e);
            alert('There was a problem with the fetch operation.');
        }
}

function Signin() {
    const navigate = useNavigate();
    const [userId, setUserId] = useState(null);

    const SubmitForm = (event) => {
        event.preventDefault(); // Prevent the default form submission behavior

        const form = event.target; // Get the form element from the event
        const formData = new FormData(form);
        const plainObject = Object.fromEntries(formData.entries());
        const jsonData = JSON.stringify(plainObject);
        console.log(jsonData);
        console.log(plainObject);

        // Parse and set userId
        const userId = parseInt(plainObject.Id);
        setUserId(userId);

        CheckForIDPass(jsonData);
    };

    return (
        <div id="form-container">
            <form id="form" onSubmit={SubmitForm}>
                <div id="components">
                    <div>
                        <label htmlFor="Id" className="Id">ID</label>
                        <input 
                            type="text" 
                            name="Id" 
                            required 
                            placeholder="Id" 
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
                        <button type="submit" id="button">Sign In</button>
                    </div>
                </div>
            </form>
        </div>
    );
}

export default Signin;
