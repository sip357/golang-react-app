import React from "react";
import { Routes, Route} from "react-router-dom";
import TaskList from "./components/TaskList";
import CreateTask from "./components/CreateTask";
import CreateAccount from "./components/signup";
import NavBar from "./components/nav/navbar";

const App = () => {
  return (
    <div className="flex">
      <NavBar/>
      <main className="ml-[25%] p-4 w-2/3">
        <Routes>
          <Route path="/" element={<TaskList/>} />
          <Route path="/users/signup" element={<CreateAccount/>} />
          <Route
            path="/create"
            element={<CreateTask/>}
          />
          <Route path="*" element={<h2>Page Not Found</h2>} />
        </Routes>  
      </main>
      
    </div>
  );
};

export default App;
