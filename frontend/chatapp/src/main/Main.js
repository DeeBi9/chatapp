import React from 'react';
import './Main.css'

function MainChat() {
    var path = window.location.href;
    console.log("The url for this page is :", path);
    
    var userId = path.split('/');
    userId = userId[userId.length - 1].trim();
    if (userId !== '') {
      console.log(userId);
    } else {
      console.log('No user ID found in the URL.');
    }

    return (
        <div className="container">
            {/* Sidebar */}
            <div className="sidebar">
                <div className="profile">
                    <h2>User Name</h2>
                </div>
                <div className="search">
                    <input type="text" placeholder="Search or start new chat" />
                </div>
                <ul className="contacts">
                    <li className="contact">
                        <div className="contact-info">
                            <h3>User 1</h3>
                            <p>Last message...</p>
                        </div>
                    </li>
                    <li className="contact">
                        <div className="contact-info">
                            <h3>User 2</h3>
                            <p>Last message...</p>
                        </div>
                    </li>
                    {/* Add more contacts as needed */}
                </ul>
            </div>

            {/* Chat Area */}
            <div className="chat-area">
                <div className="chat-header">
                    <h3>User 1</h3>
                </div>
                <div className="chat-messages">
                    <div className="message received">
                        <p className='message-text'>Hello! How are you?</p>
                        <span className="time">10:00 AM</span>
                    </div>
                    <div className="message sent">
                        <p className='message-text'>I'm good, thank you! How about you?</p>
                        <span className="time">10:02 AM</span>
                    </div>
                    {/* Add more messages as needed */}
                </div>
                <div className="chat-input">
                    <input type="text" placeholder="Type a message" />
                </div>
            </div>
        </div>
    );
}

export default MainChat;
