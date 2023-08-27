'use client'
import React from 'react'
import { useState, useEffect } from 'react';
import ChatBot from '@components/ChatBot';

const Products = () => {
  const [userUUID, setUserUUID] = useState('');
  const [showChat, setShowChat] = useState(false);
  const [products, setProducts] = useState([]);

    const getBasename = (filepath) => {
      return filepath.split('/').pop();
    };

    useEffect(() => {
        // Extracting userUUID from the window's location
        const params = new URLSearchParams(window.location.search);
        setUserUUID(params.get('userUUID'));
    }, []);

    useEffect(() => {
      fetch('http://localhost:8080/getproducts')
        .then(response => response.json())
        .then(data => setProducts(data))
        .catch(error => console.error("Error fetching products:", error));
    }, []);

  return (
    <div>
      <div>Welcome to the Marketplace! Click to purchase</div>
      <button onClick={() => setShowChat(!showChat)}>Toggle Chat</button>
      {showChat && <ChatBot user_id={userUUID} reaction="thumbs_up" product_id="944e747c-5525-4267-b01b-4a629d2af8fd"/>}
      

      <div style={{display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(200px, 1fr))'}}>
        {products.map(product => (
          <div key={product.product_name} style={{margin: '10px'}}>
            <img 
              src={`/assets/images/${getBasename(product.product_image_url)}`} 
              alt={product.product_name} 
              style={{width: '100%'}}
              onClick={() => {
                // Do something when image is clicked, e.g. navigate or open a modal
                console.log(`Clicked on product: ${product.product_name}`);
              }}
            />
            <div>{product.product_name}</div>
          </div>
        ))}
      </div>
    </div>

    
  )
}

export default Products