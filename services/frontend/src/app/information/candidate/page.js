'use client';
import { useEffect, useState } from 'react';
import Head from 'next/head';
import Navbar from "../../../components/Navbar"
import Link from 'next/link';

export default function Registration() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch('http://localhost:8000/getConstDeets');
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
              <th>Constituency Name</th>
              <th>Male Count</th>
              <th>Female Count</th>
              <th>Poll Booth Count</th>
            </tr>
          </thead>
          <tbody>
            {data.map((item, index) => (
              <tr key={index}>
                <td>{item.constituencyName}</td>
                <td>{item.maleCount}</td>
                <td>{item.femaleCount}</td>
                <td>{item.pollBoothCount}</td>
              </tr>
            ))}
          </tbody>
        </table>
        <Link legacyBehavior href="/update">
            <button className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Update
            </button>
        </Link>
      </div>
    </div>
  );
}