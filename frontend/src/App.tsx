import {useEffect, useState } from 'react'
import './App.css'
import CountryCard from './components/CountryCard'
import axios from 'axios'

const client = axios.create({
  baseURL: "http://localhost:8080" 
});

function App() {
    
  const [countries, showCountryGrid] = useState([]);

    useEffect(() => {
      client.get('/countries').then((response) => {
          showCountryGrid(response.data);
      });
    }, []);

    return (
      <div className="container text-center">
          <div className="row row-cols-4">
            {countries.map((country: any, index) => (
            <div key={index}>
              {CountryCard(country)}
            </div>    
            ))}
          </div>
        </div>
    );  
     
}    

export default App;
