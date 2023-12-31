'use client';
import { useEffect, useState } from 'react';
import Head from 'next/head';
import Navbar from "../../../components/Navbar"
import Link from 'next/link';

export default function Registration() {
  const [data, setData] = useState([]);
  const user = JSON.parse(localStorage.getItem('user'));

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("http://localhost:8000/getConstDeets");
      const data = await response.json();
      console.log(data)
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
      </div>
    </div>
  );
}