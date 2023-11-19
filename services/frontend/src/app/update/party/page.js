import Head from 'next/head';
import Navbar from "../../../components/Navbar"
import React, { useState } from 'react';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export default function VoterUpdate() {
  const handleDelete = async () => {
    try {
      const response = await fetch('http://localhost:8000/deleteParty', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Something went wrong');
      }

      toast.success('Successfully deleted Party');
    } catch (error) {
      toast.error(error.message);
    }
  }

  const [formData, setFormData] = useState({
      'partyName': '',
      'partySymbol': '',
      'president': '',
      'partyFunds': '',
      'headquarters': '',
      'partyMemberCount': ''

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
      const response = await fetch('http://localhost:8000/updateParty', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Something went wrong');
      }

      toast.success('Successfully updated Party');
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
        <h2 className="text-black text-xl mb-4">Party Update Form</h2>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="aadharId">
            Party Name
          </label>
          <input
            type="number"
            id="partyName"
            name="partyName"
            value={formData.partyName}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Party Symbol
          </label>
          <input
            type="text"
            id="partySymbol"
            name="partySymbol"
            value={formData.partySymbol}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            President
          </label>
          <input
            type="text"
            id="president"
            name="president"
            value={formData.president}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>

        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Party Funds
          </label>
          <input
            type="text"
            id="partyFunds"
            name="partyFunds"
            value={formData.partyFunds}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Headquarters
          </label>
          <input
            type="text"
            id="headquarters"
            name="headquarters"
            value={formData.headquarters}
            onChange={handleChange}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
            Party Member Count
          </label>
          <input
            type="text"
            id="partyMemberCount"
            name="partyMemberCount"
            value={formData.partyMemberCount}
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