'use client'
import React from 'react'
import { useState, useEffect } from 'react';

const Products = () => {
  const [userUUID, setUserUUID] = useState('');

    useEffect(() => {
        // Extracting userUUID from the window's location
        const params = new URLSearchParams(window.location.search);
        setUserUUID(params.get('userUUID'));
    }, []);
  return (
    <div>Products for User ID</div>
  )
}

export default Products