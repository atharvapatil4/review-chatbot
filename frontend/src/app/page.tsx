'use client'
import Image from 'next/image'
import { useState } from 'react';

export default function Home() {
    const [showLogin, setShowLogin] = useState(false);
    const [showSignup, setShowSignup] = useState(false);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [message, setMessage] = useState("");

    const loginUser = async () => {
        // Implement login logic here, e.g., hitting the Go backend.
        // On successful login, redirect user:
        // router.push('/dashboard');
        // On failure, display the error message:
        // setMessage("Failed to log in");
    };

    const signupUser = async () => {
        // Implement sign up logic here.
        // On successful sign-up, you can log the user in or show a success message.
        // On failure, display the error message:
        // setMessage("Failed to sign up");
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <div className="bg-white p-8 rounded-xl shadow-md w-96">
                <h1 className="text-2xl font-bold mb-4 text-center">Welcome to the Marketplace!</h1>
                <button onClick={() => setShowLogin(!showLogin)} className="block w-full mb-2 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:bg-blue-700">
                    Login
                </button>
                {showLogin && (
                    <form className="space-y-4" onSubmit={loginUser}>
                        <input className="w-full p-2 border rounded" type="text" placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} />
                        <input className="w-full p-2 border rounded" type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
                        <button type="submit" className="block w-full bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:bg-green-700">
                            Submit
                        </button>
                    </form>
                )}
                <button onClick={() => setShowSignup(!showSignup)} className="block w-full mt-2 bg-indigo-500 hover:bg-indigo-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:bg-indigo-700">
                    Sign Up
                </button>
                {showSignup && (
                    <form className="space-y-4 mt-2" onSubmit={signupUser}>
                        <input className="w-full p-2 border rounded" type="text" placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} />
                        <input className="w-full p-2 border rounded" type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
                        <button type="submit" className="block w-full bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:bg-green-700">
                            Register
                        </button>
                    </form>
                )}
                {message && <p className="text-red-500 mt-4">{message}</p>}
            </div>
        </div>
    );
}
