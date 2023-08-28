
import { useState, useEffect } from 'react';
import axios from 'axios';

function ChatBot({user_id, reaction, product_id, onCloseChat}) {
    const [chatLog, setChatLog] = useState([]);
    const [inputValue, setInputValue] = useState(''); // State to control prompt
    const [showChat, setShowChat] = useState(true);  // State to control chat visibility

    function createBackendURL(route) {
      const backendURL = process.env.BACKEND_URL || "http://localhost:8080";
      return `${backendURL}/${route}`;
    }
    const selectPromptRoute = "selectprompt";
    const selectPromptUrl = createBackendURL(selectPromptRoute) ;

    useEffect(() => {
        fetch(selectPromptUrl, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ reaction: reaction })
        })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          const prompt = data.prompt;
          setChatLog([{ sender: 'bot', message: prompt }]);
        })
        .catch(error => {
          console.error("Error fetching the prompt:", error);
          setChatLog([{ sender: 'bot', message: "Hello! There was an issue fetching the prompt. How can I assist you?" }]);
        });
    }, [reaction]);
    
  const handleInputChange = (e) => {
    setInputValue(e.target.value);
  };


  const writeReviewRoute = "writereview";
    const writeReviewUrl = createBackendURL(writeReviewRoute);

  const handleSend = () => {
    const userMessage = { sender: 'user', message: inputValue };
    setChatLog(prevChatLog => [...prevChatLog, userMessage]);

    const timestamp = new Date().toISOString();

    // Make an API call
    fetch(writeReviewUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            user_id: user_id,
            product_id: product_id,
            review_description: inputValue,
            timestamp: timestamp
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        // On successful response, display a thank you message
        setChatLog(prevChatLog => [...prevChatLog, { sender: 'bot', message: "Thank you!" }]);
        
        // Close the chat window after 3500 seconds
        setTimeout(() => {
            //setShowChat(false);
            onCloseChat();
        }, 3500);
    })
    .catch(error => {
        console.error("Error sending review:", error);
        setChatLog(prevChatLog => [...prevChatLog, { sender: 'bot', message: "Sorry, there was an issue. Please try again." }]);
    });
    
    setInputValue('');
};

if (!showChat) return null;  // Don't render the chat if showChat is false

  return (
    <div className="chatbot">
      <div className="chat-log">
        {chatLog.map((entry, index) => (
          <div key={index} className={`chat-entry ${entry.sender}`}>
            <p>{entry.message}</p>
          </div>
        ))}
      </div>
      <div className="chat-input">
        <input value={inputValue} onChange={handleInputChange} placeholder="Type your message..." />
        <button onClick={handleSend}>Send</button>
      </div>

     <style jsx>{`
        .chatbot {
        width: 300px;
        height: 400px;
        border: 1px solid #ddd;
        position: fixed;
        right: 20px;
        bottom: 20px;
        background-color: #f9f9f9;
        display: flex;
        flex-direction: column;
        }

        .chat-log {
        flex: 1;
        overflow-y: auto;
        padding: 10px;
        }

        .chat-input {
        padding: 10px;
        border-top: 1px solid #ddd;
        display: flex;
        align-items: center;
        }

        .chat-input input {
        flex: 1;
        padding: 5px;
        }

        .chat-entry.bot {
        align-self: flex-start;
        max-width: 80%; 
        margin-right: 20%;  // This will keep the bot's message to the left
        background-color: #e6e6e6;
        padding: 5px;
        margin-bottom: 5px;
        border-radius: 5px;
        }

        .chat-entry.user {
        align-self: flex-end;
        max-width: 80%;
        margin-left: 20%;  // This will push the user's message to the right
        background-color: #4caf50;
        color: white;
        padding: 5px;
        margin-bottom: 5px;
        border-radius: 5px;
        }
    `}</style>
 
    </div>
  );
}

export default ChatBot;
