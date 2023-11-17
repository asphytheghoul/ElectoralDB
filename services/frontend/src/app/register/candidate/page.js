'use client';
import Head from 'next/head';
import Navbar from "../../../components/Navbar"
import React, { useState } from 'react';
import { ToastContainer,toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Link from 'next/link';

export default function Registration() {
  const [formData, setFormData] = useState({
    aadharId: '',
    firstName: '',
    middleName: '',
    lastName: '',
    gender: '',
    dob: '',
    age: '',
    phone: '',
    constituency: '',
    candidateId: '',
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };
  const handleSubmit = async(e) => {
    e.preventDefault();    
    try {
      const response = await fetch('http://localhost:8000/insertCandidate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Something went wrong');
      }

      toast.success('Successfully registered candidate');
    } catch (error) {
      toast.error(error.message);
    }
  };

  return (
    <div className="bg-white text-black">
      <Head>
        <title>ELECTORAL DB</title>
      </Head>
      <Navbar />
    <div className="w-full max-w-lg mx-auto">
      <form onSubmit={handleSubmit} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <h2 className="text-black text-xl mb-4">Candidate Registration Form</h2>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="aadharId">
            Aadhar ID
          </label>
          <input
            type="number"
            id="aadharId"
            name="aadharId"
            value={formData.aadharId}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            First Name
          </label>
          <input
            type="text"
            id="firstName"
            name="firstName"
            value={formData.firstName}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Middle Name
          </label>
          <input
            type="text"
            id="middleName"
            name="middleName"
            value={formData.middleName}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Last Name
          </label>
          <input
            type="text"
            id="lastName"
            name="lastName"
            value={formData.lastName}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <label className=" text-gray-700 text-sm font-bold mb-2" htmlFor="firstName"> Gender</label>
       
        <div className="flex items-center justify-around mb-4 mt-4">
            <div className='ml-2'>
            <input
                type="radio"
                id="gender"
                name="gender"
                value="Male"
                checked={formData.gender === "Male"}
                onChange={handleChange}
                className="form-radio text-black h-5 w-5"
            />
            
            <div >Male</div>
            </div>
            <div className='ml-2'>
            <input
                type="radio"
                id="gender"
                name="gender"
                value="Female"
                checked={formData.gender === "Female"}
                onChange={handleChange}
                className="form-radio text-black h-5 w-5"
            />
            <div >Female</div>
            </div>
        </div>
        
        <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="dob">
                Date of Birth (DD-MM-YY)
            </label>

            <input
            type="text"
            id="dob"
            name="dob"
            value={formData.dob}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
            </div>
        

        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="aadharId">
            Phone Number
          </label>
          <input
            type="number"
            id="phone"
            name="phone"
            value={formData.phone}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>

        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="aadharId">
            Constituency Name
          </label>
          <input
            type="text"
            id="constituency"
            name="constituency"
            value={formData.constituency}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="aadharId">
            Party Representative
          </label>
          <input
            type="text"
            id="partyRep"
            name="partyRep"
            value={formData.partyRep}
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
    </div>
    </div>
  );
}
