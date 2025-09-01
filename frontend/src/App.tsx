import { useState } from 'react'
import { BrowserRouter as Router } from "react-router-dom";
import Cart from "./pages/cart";

import './App.css'
import React from 'react';

function App() {

  return (
    <>
      <React.StrictMode>
      <Router>

        <Cart />
      </Router>
    </React.StrictMode>
    </>
  )
}

export default App
