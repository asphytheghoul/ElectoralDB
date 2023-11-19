'use client';
import { useEffect, useState } from 'react';
import Head from 'next/head';
import Navbar from "../../../components/Navbar"
import Link from 'next/link';

export default function Registration() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch('http://localhost:8000/getvoterinformation');
      const data = await response.json();
      setData(data);
    };

    fetchData();
  }, []);

  return (
    <div className="bg-white text-black">
      <Head>
        <title>ELECTORAL DB</title>
      </Head>
      <Navbar />
      <div className="w-full max-w-lg mx-auto pb-96">
        <table>
          <thead>
            <tr>
              <th>Aadhar ID</th>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Middle Name</th>
              <th>Gender</th>
              <th>DOB</th>
              <th>Age</th>
              <th>State</th>
              <th>Phone Number</th>
              <th>Constituency Name</th>
              <th>Polling Booth ID</th>
              <th>Voter ID</th>
            </tr>
          </thead>
          <tbody>
            {data.map((item, index) => (
              <tr key={index}>
                <td>{item.aadharID}</td>
                <td>{item.firstName}</td>
                <td>{item.lastName}</td>
                <td>{item.middleName}</td>
                <td>{item.gender}</td>
                <td>{item.dob}</td>
                <td>{item.age}</td>
                <td>{item.state}</td>
                <td>{item.phoneNumber}</td>
                <td>{item.constituencyName}</td>
                <td>{item.pollingBoothID}</td>
                <td>{item.voterID}</td>
              </tr>
            ))}
          </tbody>
        </table>
        <Link legacyBehavior href="/update/voter">
            <button className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Update
            </button>
        </Link>
      </div>
    </div>
  );
}