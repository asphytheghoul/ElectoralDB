'use client';
import Image from 'next/image';
import Link from 'next/link';
import { useEffect, useState } from 'react';

const Home = () => {
  const [user, setUser] = useState(null);
  const [showDropdown, setShowDropdown] = useState(false);
  const [selectedOption, setSelectedOption] = useState(null);
  const handleLogout = () => {
    localStorage.removeItem('user');
    setUser(null);
  };
  
  useEffect(() => {
    const storedUser = JSON.parse(localStorage.getItem('user'));
    if (storedUser) {
      setUser(storedUser);
    }
  }, []);


  const options = {
    electors: ["Voter Registration",  "Voter Information", "Query Procedure"],
    candidates: ["Candidate Registration","Candidate Information","Query Procedure"],
    parties: ["Party Registration", "Party Information","Query Procedure"],
    officials: ["Official Registration", "Official Information", "Query Procedure","Edit Constituency","Edit Poll booth","Edit Elections"]
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
    <div className="flex justify-between p-2">
    <Link legacyBehavior href="/">
    <div className="flex items-center cursor-pointer">
      <Image src="/logo.png" alt="Logo" width={125} height={125} />
      <h1 className="text-2xl ml-2 pl-4 ">ELECTORAL DB</h1>
    </div>
    </Link>
    <nav className="ml-64 flex items-center">
    {user && user.role==="voter" && (
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
            {options.electors.map((item, index) => {
              let link;
              switch(index) {
                case 0:
                  link = "/register/voter";
                  break;
                case 1:
                  link = "/information/voter";
                  break;
                case 2:
                  link = "/query";
                  break;
                default:
                  link = "/register/voter";
              }
              return (
                <Link legacyBehavior href={link} key={index}>
                  {item}
                </Link>
              )
            })}
          </div>
        )}
      </div>
    )}
    {/* {user && user.role === 'candidate' && (

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
        <Link legacyBehavior href={item === 'Candidate Information' ? '/information/candidate' : `/register/candidate/`} key={index}>
          {item}
        </Link>
      ))}
    </div>
  )}
</div>
)} */}
{user && user.role === 'candidate' && (
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
        {options.candidates.map((item, index) => {
          let link;
          switch(index) {
            case 0:
              link = "/register/candidate";
              break;
            case 1:
              link = "/information/candidate";
              break;
            case 2:
              link = "/query/candidate";
              break;
            default:
              link = "/register/candidate";
          }
          return (
            <Link legacyBehavior href={link} key={index}>
              {item}
            </Link>
          )
        })}
      </div>
    )}
  </div>
)}
{user && user.role === 'party' && (
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
        {options.parties.map((item, index) => {
          let link;
          switch(index) {
            case 0:
              link = "/register/party";
              break;
            case 1:
              link = "/information/party";
              break;
            case 2:
              link = "/query";
              break;
            default:
              link = "/register/party";
          }
          return (
            <Link legacyBehavior href={link} key={index}>
              {item}
            </Link>
          )
        })}
      </div>
    )}
  </div>
)}
{user && user.role === 'official' && (
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
        {options.officials.map((item, index) => {
          let link;
          switch(index) {
            case 0:
              link = "/register/eci";
              break;
            case 1:
              link = "/information/officials";
              break;
            case 2:
              link = "/query";
              break;
            default:
              link = "/register/eci";
          }
          return (
            <Link legacyBehavior href={link} key={index}>
              {item}
            </Link>
          )
        })}
      </div>
    )}
  </div>
)}
    </nav>
    <div className="ml-auto pt-8">
    <Link legacyBehavior href="/loginuser">
      <button className="bg-black text-white p-4 mr-4 rounded-full">{user ? user.aadhar_id : 'Login/Register'}</button>
      </Link>
      {user && <button onClick={handleLogout}>Logout</button>}
    </div>
  </div>
        );
    };

    export default Home;