'use client';
import Head from 'next/head';
import Image from 'next/image';
import Navbar from '../components/Navbar';
const Home = () => {
  return (
    <div className="bg-white text-black">
      <Head>
        <title>ELECTORAL DB</title>
      </Head>
     <Navbar />

      <div className='flex justify-center align-center'>
        <Image src="/india.svg" alt="Logo" width={500} height={500} />
      </div>
    </div>
    );
  };

export default Home;
