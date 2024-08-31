import React from 'react';
import { useNavigate } from "react-router-dom";
import './Signin.css';

// Function to handle the form submission and make the request
async function checkForIDPass(jsonData, navigate) {
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
            const jsonResponse = await response.json();
            const token = jsonResponse.token; 
            console.log("Received JWT:", token);

            // Store the token in localStorage
            localStorage.setItem('jwtToken', token);

            // Navigate to another page or update the app state
            alert("Welcome. Click Ok to proceed");
            navigate('/main'); // Example: Navigate to a protected route
        } else {
            const jsonResponse = await response.json();
            console.error("Error:", jsonResponse.message || "Unknown error");
            console.log(jsonResponse)
            alert("An error occurred: " + (jsonResponse.message || "Unknown error"));
        }

    } catch (e) {
        console.error('Problem occurred while fetching the information:', e);
        alert('There was a problem with the fetch operation.');
    }
}

function Signin() {
    const navigate = useNavigate(); // useNavigate hook for navigation

    const submitForm = (event) => {
        event.preventDefault(); // Prevent the default form submission behavior

        const form = event.target; // Get the form element from the event
        const formData = new FormData(form);
        const plainObject = Object.fromEntries(formData.entries());
        const jsonData = JSON.stringify(plainObject);

        console.log("Submitted JSON Data:", jsonData);

        // Call the function to check ID and password, and handle JWT
        checkForIDPass(jsonData, navigate);
    };

    return (
        <div id="form-container">
            <form id="form" onSubmit={submitForm}>
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
