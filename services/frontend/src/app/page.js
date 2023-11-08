'use client';
import Head from 'next/head';
import Image from 'next/image';
import { useState } from 'react';

const Home = () => {
  const [showDropdown, setShowDropdown] = useState(false);
  const [selectedOption, setSelectedOption] = useState(null);

  const options = {
    electors: ["Voter Registration", "Candidate Information", "Voter Information"],
    candidates: ["Candidate Registration", "Voter Information", "Candidate Information", "Party Information"],
    parties: ["Party Registration", "Voter Information", "Candidate Information", "Party Information"],
    officials: ["Official Registration", "Official Information", "Candidate Information", "Party Information", "Voter Information"]
  };

  const toggleDropdown = (option) => {
    if (selectedOption === option) {
      setShowDropdown(false);
      setSelectedOption(null);
    } else {
      setShowDropdown(true);
      setSelectedOption(option);
    }
  };

  return (
    <div className="bg-white h-screen text-black">
      <Head>
        <title>ELECTORAL DB</title>
      </Head>
      <div className="flex justify-between p-2">
        <div className="flex items-center">
          <Image src="/logo.png" alt="Logo" width={125} height={125} />
          <h1 className="text-2xl ml-2 pl-4 ">ELECTORAL DB</h1>
        </div>
        <nav className="ml-64 flex items-center">
          <div className="dropdown">
            <a
              href="#"
              className={`text-xl mx-4 ${selectedOption === 'electors' ? 'active' : ''}`}
              onClick={() => toggleDropdown('electors')}
            >
              Electors
            </a>
            {showDropdown && selectedOption === 'electors' && (
              <div className="dropdown-content">
                {options.electors.map((item, index) => (
                  <a key={index} href="#">
                    {item}
                  </a>
                ))}
              </div>
            )}
          </div>

          <div className="dropdown">
            <a
              href="#"
              className={`text-xl mx-4 ${selectedOption === 'candidates' ? 'active' : ''}`}
              onClick={() => toggleDropdown('candidates')}
            >
              Candidates
            </a>
            {showDropdown && selectedOption === 'candidates' && (
              <div className="dropdown-content">
                {options.candidates.map((item, index) => (
                  <a key={index} href="#">
                    {item}
                  </a>
                ))}
              </div>
            )}
          </div>

          <div className="dropdown">
            <a
              href="#"
              className={`text-xl mx-4 ${selectedOption === 'parties' ? 'active' : ''}`}
              onClick={() => toggleDropdown('parties')}
            >
              Parties
            </a>
            {showDropdown && selectedOption === 'parties' && (
              <div className="dropdown-content">
                {options.parties.map((item, index) => (
                  <a key={index} href="#">
                    {item}
                  </a>
                ))}
              </div>
            )}
          </div>

          <div className="dropdown">
            <a
              href="#"
              className={`text-xl mx-4 ${selectedOption === 'officials' ? 'active' : ''}`}
              onClick={() => toggleDropdown('officials')}
            >
              ECI Officials
            </a>
            {showDropdown && selectedOption === 'officials' && (
              <div className="dropdown-content">
                {options.officials.map((item, index) => (
                  <a key={index} href="#">
                    {item}
                  </a>
                ))}
              </div>
            )}
          </div>
        </nav>
        <div className="ml-auto pt-8">
          <button className="bg-black text-white p-4 mr-4 rounded-full">Login / Register</button>
        </div>
      </div>

      <div className='flex justify-center align-center'>
        <Image src="/india.svg" alt="Logo" width={500} height={500} />
      </div>
    </div>
  );
};

export default Home;
