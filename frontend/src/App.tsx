import {useEffect, useState } from 'react'
import './App.css'
import Country from './Country'
import axios from 'axios'

const client = axios.create({
  baseURL: "http://localhost:8080" 
});

function App() {
    
  const [countries, showCountries] = useState([]);

    useEffect(() => {
      client.get('/countries').then((response) => {
          showCountries(response.data);
      });
    }, []);

    return (
      countries.map((country: any, index) => (
          <div key={index}>
            {Country(country)}
          </div>    
          ))
    );  
     
}    

export default App
