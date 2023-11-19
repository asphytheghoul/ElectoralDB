'use client';
import { useEffect, useState } from 'react';
import Head from 'next/head';
import Navbar from "../../../components/Navbar"
import Link from 'next/link';

export default function Registration() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch('http://localhost:8000/getofficialinformation');
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
              <th>Phone Number</th>
              <th>Constituency Assigned</th>
              <th>Poll Booth Assigned</th>
              <th>Official ID</th>
              <th>Official Rank</th>
              <th>Higher Rank ID</th>
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
                <td>{item.phoneNumber}</td>
                <td>{item.constituencyAssigned}</td>
                <td>{item.pollBoothAssigned}</td>
                <td>{item.officialID}</td>
                <td>{item.officialRank}</td>
                <td>{item.higherRankID}</td>
              </tr>
            ))}
          </tbody>
        </table>
        <Link legacyBehavior href="/update/officials">
            <button className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Update
            </button>
        </Link>
      </div>
    </div>
  );
}