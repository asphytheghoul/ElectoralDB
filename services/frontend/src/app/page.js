'use client';
import Head from 'next/head';
import Image from 'next/image';
import Navbar from '../components/Navbar';
import React ,{useState} from 'react';
import AuthContext from '../components/AuthContext';
const Home = () => {
  const [user, setUser] = useState(null);
  return (
    <AuthContext.Provider value={{user, setUser}}>
    <div className="bg-white text-black">
      <Head>
        <title>ELECTORAL DB</title>
      </Head>
     <Navbar />

      <div className='flex justify-center align-center'>
        <Image src="/india.svg" alt="Logo" width={500} height={500} />
      </div>
    </div>
    </AuthContext.Provider>
    );
  };

export default Home;
