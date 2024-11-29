import React, { useState } from "react";
import { signUp } from "../api";

const CreateAccount = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    const user = { username, email, password };
    const response = await signUp(user); //Use API to create user

    try {
      if(response.status === 201) {
        alert("User created successfully");
        setUsername(""); // Reset all states
        setEmail("");
        setPassword("");
      } else {
        alert(`Error: ${response.statusText}`);
      }
    } catch (error) {
      console.error("An error occured:" `${error}`)
      alert("Failed to create user. Please try again.")
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2 className="segment">Sign Up</h2>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        required
      />
      <input
        type="text"
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        required
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
      />
      <button type="submit">Create Account</button>
    </form>
  );
};

export default CreateAccount;