'use client';
import { useEffect, useState } from 'react';
import Head from 'next/head';
import Navbar from "../../../components/Navbar"
import Link from 'next/link';

export default function Registration() {
  const [data, setData] = useState([]);
  const user = JSON.parse(localStorage.getItem('user'));
  const aadhar_id = user.aadhar_id;

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(`http://localhost:8000/getofficialinformation?aadharId=${aadhar_id}`);
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
              <tr>
                <td>{data.aadharId}</td>
                <td>{data.firstName}</td>
                <td>{data.lastName}</td>
                <td>{data.middleName}</td>
                <td>{data.gender}</td>
                <td>{data.dob}</td>
                <td>{data.age}</td>
                <td>{data.phoneNumber}</td>
                <td>{data.constituencyAssigned}</td>
                <td>{data.pollBoothAssigned}</td>
                <td>{data.officialId}</td>
                <td>{data.officialRank}</td>
                <td>{data.higherRankId}</td>
              </tr>
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