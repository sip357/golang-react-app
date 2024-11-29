import React, { useState } from "react";
import { createTask } from "../api";
import { useNavigate} from "react-router-dom";

const CreateTask = () => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const task = { title, content };
    await createTask(task);
    setTitle("");
    setContent("");
    navigate("/");
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2 className="segment">Create Task</h2>
      <input
        type="text"
        placeholder="Title"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        required
      />
      <textarea
        placeholder="Content"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        required
      ></textarea>
      <button type="submit">Add Task</button>
    </form>
  );
};

export default CreateTask;
