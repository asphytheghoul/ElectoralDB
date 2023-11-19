import Head from 'next/head';
import Navbar from "../../../components/Navbar"
import React, { useState } from 'react';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export default function VoterUpdate() {
  const handleDelete = async () => {
    try {
      const response = await fetch('http://localhost:8000/deleteVoter', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Something went wrong');
      }

      toast.success('Successfully deleted voter');
    } catch (error) {
      toast.error(error.message);
    }
  }

  const [formData, setFormData] = useState({
    aadharId: '',
    firstName: '',
    lastName: '',
    middleName: '',
    gender: '',
    dob: '',
    state: '',
    phone: '',
    voterId: ''
  });

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
      const response = await fetch('http://localhost:8000/updateVoter', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Something went wrong');
      }

      toast.success('Successfully updated voter');
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
        <h2 className="text-black text-xl mb-4">Candidate Update Form</h2>
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
            Gender
          </label>
          <input
            type="text"
            id="gender"
            name="gender"
            value={formData.gender}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            DOB
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
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            State
          </label>
          <input
            type="text"
            id="state"
            name="state"
            value={formData.state}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
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
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Voter ID
          </label>
          <input
            type="text"
            id="voterId"
            name="voterId"
            value={formData.voterId}
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
            <button onClick={handleDelete} className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Delete Record
            </button>
    </div>
    <ToastContainer/>
    </div>
  );
}