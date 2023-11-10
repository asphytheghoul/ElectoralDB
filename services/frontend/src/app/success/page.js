'use client';
import Navbar from "../../components/Navbar"
import React from 'react';

export default function Page() {
    return (
        <div className="bg-white text-black">
        <Navbar />
        <div className="flex justify-center align-center">
            <img src="/check.svg" alt="Logo" width={400} height={400} />
            
        </div>
        <h1 className="text-4xl flex justify-center pb-32">Successfully Submitted!</h1>
        </div>
    );
    }