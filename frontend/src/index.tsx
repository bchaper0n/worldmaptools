import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import axios from "axios";
import Country from './Country.tsx';

const client = axios.create({
   baseURL: "http://localhost:8080" 
});

const Countries = () => {
  
  const [countries, showCountries] = useState([]);

  useEffect(() => {
    client.get('/countries').then((response) => {
        showCountries(response.data);
    });
  }, []);

  return (
      <div>
        {countries.map((country: any) => (
          Country(country)
        ))}
      </div>
  )
}

//get root
const root = ReactDOM.createRoot(document.getElementById('root') as Element);
root.render(
  <Countries></Countries>
  //TODO: incorporate: https://www.react-simple-maps.io/docs/map-files/
);
