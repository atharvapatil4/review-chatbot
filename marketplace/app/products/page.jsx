'use client'
import React from 'react'
import { useState, useEffect } from 'react';
import ChatBot from '@components/ChatBot';

function ReactionModal({ show, onReactionSelected, productId }) {
  if (!show) return null;

  return (
    <div className="modal">
      <div className="content">
        <div className="question">Did you like this product?</div>
        <div className="buttons">
          <button className="like-button" onClick={() => onReactionSelected('thumbs_up', productId)}>👍 Like</button>
          <button className="dislike-button" onClick={() => onReactionSelected('thumbs_down', productId)}>👎 Dislike</button>
        </div>
      </div>

      <style jsx>{`
        .modal {
          position: fixed;
          top: 0;
          left: 0;
          right: 0;
          bottom: 0;
          background: rgba(0, 0, 0, 0.5);
          display: flex;
          justify-content: center;
          align-items: center;
        }
        .content {
          background: white;
          padding: 20px;
          border-radius: 8px;
          width: 80%;
          max-width: 400px;
          text-align: center;
          box-shadow: 0px 0px 20px rgba(0,0,0,0.2);
        }
        .question {
          margin-bottom: 20px;
          font-size: 18px;
        }
        .buttons {
          display: flex;
          justify-content: space-around;
        }
        .like-button, .dislike-button {
          padding: 10px 20px;
          font-size: 16px;
          border-radius: 5px;
          border: none;
          cursor: pointer;
          transition: background-color 0.3s;
        }
        .like-button {
          background-color: #4caf50;
          color: white;
        }
        .like-button:hover {
          background-color: #45a049;
        }
        .dislike-button {
          background-color: #f44336;
          color: white;
        }
        .dislike-button:hover {
          background-color: #da190b;
        }
      `}</style>
    </div>
  );
}


const Products = () => {
  const [userUUID, setUserUUID] = useState('');
  const [showChat, setShowChat] = useState(false);
  const [products, setProducts] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [selectedProductId, setSelectedProductId] = useState(null);
  const [selectedReaction, setSelectedReaction] = useState(null);

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

  const handleChatClose = () => {
    setShowChat(false);
  };

  return (
    <div>
      <div>Welcome to the Marketplace! Click to purchase any item.</div>
      
      {showChat && <ChatBot user_id={userUUID} reaction={selectedReaction} product_id={selectedProductId} onCloseChat={handleChatClose} />}

      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(200px, 1fr))' }}>
        {products.map(product => (
          <div key={product.product_name} style={{ margin: '10px' }}>
            <img
              src={`/assets/images/${getBasename(product.product_image_url)}`}
              alt={product.product_name}
              style={{ width: '100%' }}
              onClick={() => {
                setSelectedProductId(product.product_id);
                setShowModal(true);
              }}
            />
            <div>{product.product_name}</div>
          </div>
        ))}
      </div>

      <ReactionModal
        show={showModal}
        onClose={() => setShowModal(false)}
        onReactionSelected={(reaction, productId) => {
          setShowChat(true);
          setSelectedReaction(reaction);
          setShowModal(false);
        }}
        productId={selectedProductId}
      />
    </div>
  );
}

export default Products;