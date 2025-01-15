import React, {useState, useEffect} from "react";
import { loginHandler, protectedRoute } from "../../api";
import { useNavigate} from "react-router-dom";

export default function LoginPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const navigate = useNavigate();

    //Function to check the Logged In status
    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const user = { username, password };
            await loginHandler(user); // Use API to create user
            alert("Login Successful");
            setUsername(""); // Reset all states
            setPassword("");
            // navigate("/");
        } catch (error) {
            console.error("An error occurred:", error);
            alert("Wrong username or password");
        }
    };
    

    useEffect(() => {
        const checkAuth = async () => {
            try {
                await protectedRoute();
                setIsLoggedIn(true)
            } catch (e) {
                console.error("Authentication error:", e);
                setIsLoggedIn(false);
            }
        };
    
        checkAuth();  // Call the async function inside the effect
    }, []);  // Empty dependency array to run it once when the component mounts
    

    return(
        <div>
        {isLoggedIn ? (
            <div>
            <h2>Welcome back!</h2>
            <p>You are logged in.</p>
            </div>
        ) : (
            <form onSubmit={handleSubmit}>
            <h2>Login</h2>
            <input
                type="text"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
            />
            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
            />
            <button type="submit">Login</button>
            </form>
        )}
        </div>
    )
}