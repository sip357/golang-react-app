import React, { useState, useEffect } from "react";
import TaskCard from "../components/taskUI/taskDisplay"; // Import the TaskCard component
import { fetchTasks, deleteTask, updateTask } from "../api"; // Import necessary API functions

const TaskList = () => {
  const [tasks, setTasks] = useState([]);

  // Fetch tasks on component mount
  useEffect(() => {
    const getTasks = async () => {
      const data = await fetchTasks(); // Fetch tasks from API
      setTasks(data);
    };
    getTasks();
  }, []);

  // Handle task deletion
  const handleDelete = async (id) => {
    if(window.confirm("Are you sure you want to delete this task?")) {
      await deleteTask(id); // Call API to delete task
      setTasks(tasks.filter((task) => task.id !== id)); // Update state
    }
  };

  const handleEditTask = async (updatedTask) => {
    await updateTask(updatedTask); // Call API to update task
    console.log("Edited Task:", updatedTask);
  };

  return (
    <div className="p-6 bg-inherit text-black h-auto w-full">
      <h2 className="text-xl font-semibold text-white text-center mb-4">Task List</h2>
      <ul className="space-y-4">
        {tasks.map((task) => (
          <li key={task.id}>
            {/* Render TaskCard */}
            <TaskCard
              id={task.id}
              title={task.title}
              content={task.content}
              onEdit={handleEditTask}
              onDelete={() => handleDelete(task.id)}
            />
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TaskList;
