'use client';
import Head from 'next/head';
import Navbar from "../../components/Navbar"
import React, { useState } from 'react';
import Link from 'next/link';

export default function LoginUser() {
    const [formData, setFormData] = useState({
      aadharId: '',
      userName: '',
      password: '',
    });
    const [user, setUser] = useState(null);
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({
          ...formData,
          [name]: value,
        });
      };
      
    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
          const response = await fetch('http://localhost:8000/login', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData),
          });
    
          if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
          }
          const data = await response.json();
          setUser(data);
    
          // Handle successful registration, if needed
          console.log(data);
        } catch (error) {
          alert('no registered user, please register');
        }

    };


return (
    <div className="bg-white text-black">
      <Head>
        <title>ELECTORAL DB</title>
      </Head>
      <Navbar user={user}/>
    <div className="w-full max-w-lg mx-auto">
      <form onSubmit={handleSubmit} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <h2 className="text-black text-xl mb-4">User Login Form</h2>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Aadhar ID
          </label>
          <input
            type="number"
            id="aadharID"
            name="aadharID"
            value={formData.aadharID}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Password
          </label>
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        
        <div className="mb-4">
          <button
            type="submit"
            className="bg-black text-white py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          >
            Submit
          </button>
        </div>
      </form>
      <h2 className="text-black text-xl mb-4 pb-52">
  Don't have an account yet? Click 
  <Link legacyBehavior href='/registeruser'>
    <span style={{ color: 'red', fontWeight: 'bold' ,'cursor':'pointer'}}> here </span>
  </Link> 
  to sign up 
</h2>
    </div>
    </div>
  );
}