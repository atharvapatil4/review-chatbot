'use client'
import React from 'react'
import { useState, useEffect } from 'react';
import ChatBot from '@components/ChatBot';

const Products = () => {
  const [userUUID, setUserUUID] = useState('');
  const [showChat, setShowChat] = useState(false);

    useEffect(() => {
        // Extracting userUUID from the window's location
        const params = new URLSearchParams(window.location.search);
        setUserUUID(params.get('userUUID'));
    }, []);
  return (
    <div>Products for User ID

<button onClick={() => setShowChat(!showChat)}>Toggle Chat</button>
      {showChat && <ChatBot user_id={userUUID} reaction="thumbs_up" product_id="944e747c-5525-4267-b01b-4a629d2af8fd"/>}
    </div>
    
  )
}

export default Products