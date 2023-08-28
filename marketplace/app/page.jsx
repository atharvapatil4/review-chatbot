'use client'
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';


const Home = () => {
  const targetText = "Buy almost anything, from anywhere, anytime.";
  const router = useRouter();
  const [typedText, setTypedText] = useState("");
  const [charIndex, setCharIndex] = useState(0);
  const [showLogin, setShowLogin] = useState(false);
  const [showSignup, setShowSignup] = useState(false);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState("");

  useEffect(() => {
    // Check if the text has been fully typed out
    if (charIndex < targetText.length) {
      // Use a timer to simulate the delay between key presses
      const timerId = setTimeout(() => {
        setTypedText((prevTypedText) => prevTypedText + targetText[charIndex]);
        setCharIndex((prevCharIndex) => prevCharIndex + 1);
      }, 40); // 150ms delay for each character

      // Clean up the timeout when unmounting or re-rendering
      return () => clearTimeout(timerId);
    }
  }, [charIndex, targetText]);

  const loginUser = async () => {
    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
        });

        const data = await response.json();
        console.log(data)

        if (response.ok) {
            const userUUID = data.user_id;
            console.log(data, data.user_id)
            window.location.href = `/products?userUUID=${userUUID}`; // Navigating using the window object
            
        } else {
            setMessage(data.error || "Failed to log in");  // Assuming "error" is the key for error message in response.
        }
    } catch (error) {
        setMessage("An unexpected error occurred. Please try again later.");
    }
    };

    const signupUser = async () => {
    // Implement sign up logic here.
    // On successful sign-up, you can log the user in or show a success message.
    // On failure, display the error message:
    // setMessage("Failed to sign up");
    };

  return (
    <section className='w-full min-h-screen flex-center flex-col'>
      <h1 className='head_text text-center'>
        Welcome to the
        <br className='max-md:hidden' />
        <span className='orange_gradient text-center'> Marketplace </span>
      </h1>
      <p className='desc text-center font-semibold'>
        {typedText}
      </p>
      <button onClick={() => setShowLogin(!showLogin)} className="black_btn">
                    Login
                </button>
                {showLogin && (
                    <form className="space-y-4" onSubmit={loginUser}>
                        <input className="w-full p-2 border rounded" type="text" placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} />
                        <input className="w-full p-2 border rounded" type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
                        <button type="submit" className="black_btn ">
                            Submit
                        </button>
                    </form>
                )}
                <button onClick={() => setShowSignup(!showSignup)} className="black_btn">
                    Sign Up
                </button>
                {showSignup && (
                    <form className="space-y-4 mt-2" onSubmit={signupUser}>
                        <input className="w-full p-2 border rounded" type="text" placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} />
                        <input className="w-full p-2 border rounded" type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
                        <button type="submit" className="black_btn">
                            Register
                        </button>
                    </form>
                )}
                {message && <p className="text-red-500 mt-4">{message}</p>}
    </section>
  );
};

export default Home;
